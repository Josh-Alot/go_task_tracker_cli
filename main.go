package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	t "github.com/go_task_tracker/tasks"
)

const fileName = "tasks.json"

func main() {
	// define the subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("usage: todo-cli [command]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NArg() == 0 {
			fmt.Println("Error: task description is required")
			os.Exit(1)
		}

		description := addCmd.Args()[0]
		newTask := t.Task{
			Description: description,
			CreatedAt:   time.Now().Format(time.RFC3339),
			UpdatedAt:   time.Now().Format(time.RFC3339),
		}

		err := t.CreateTask([]t.Task{newTask}, fileName)

		if err != nil {
			fmt.Printf("Error creating new task: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Task added successfully")

	case "list":
		listCmd.Parse(os.Args[2:])
		var tasks []t.Task
		var err error

		if listCmd.NArg() > 0 {
			status := listCmd.Args()[0]

			switch status {
			case "done":
				tasks, err = t.ListCompleteTasks(fileName)
			case "in_progress":
				tasks, err = t.ListInProgressTasks(fileName)
			case "todo":
				tasks, err = t.ListTodoTasks(fileName)
			case "incomplete":
				tasks, err = t.ListIncompleteTasks(fileName)
			default:
				fmt.Println("Error: Unknown status. Use only \"todo\", \"in_progress\", \"done\" or \"incomplete\" status")
				os.Exit(1)
			}
		} else {
			tasks, err = t.ListTasks(fileName)
		}

		if err != nil {
			fmt.Printf("Error loading tasks: %v", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 2, 2, 0, ' ', 1)
		fmt.Println("ID\tDescription\tStatus")

		for _, task := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, task.Description, task.Status)
		}

		w.Flush()
	case "update":
		updateCmd.Parse(os.Args[2:])

		if updateCmd.NArg() < 2 {
			fmt.Println("Error: usage is update [id] [status|description] [value]")
			os.Exit(1)
		}

		idStr := updateCmd.Args()[0]
		updateBy := updateCmd.Arg(1)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Error: invalid ID: %v\n", err)
			os.Exit(1)
		}

		switch updateBy {
		case "todo", "in_progress", "done":
			var status t.Status

			switch updateBy {
			case "todo":
				status = t.Todo
			case "in_progress":
				status = t.InProgress
			case "done":
				status = t.Done
			}

			_, err := t.UpdateTaskStatus(fileName, id, status)
			if err != nil {
				fmt.Printf("Error updating task: %v", err)
			}
		case "description":
			desc := updateCmd.Args()[2]
			_, err := t.UpdateTaskDescription(fileName, id, desc)
			if err != nil {
				fmt.Printf("Error updating task: %v", err)
				os.Exit(1)
			}
		default:
			fmt.Println("Error: update by status or description only")
			os.Exit(1)
		}

		fmt.Println("Task updated successfully")
	case "delete":
		deleteCmd.Parse(os.Args[2:])

		if deleteCmd.NArg() < 1 {
			fmt.Printf("Error: insert an ID to delete\n")
			os.Exit(1)
		}

		idStr := deleteCmd.Args()[0]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Error: invalid ID: %v\n", err)
			os.Exit(1)
		}

		_, err = t.DeleteTask(fileName, id)
		if err != nil {
			fmt.Printf("Error deleting task: %v", err)
			os.Exit(1)
		}
	}
}
