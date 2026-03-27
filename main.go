package main

import (
	"fmt"

	t "github.com/go_task_tracker/tasks"
)

func main() {
	fmt.Printf("Go Task Tracker CLI")

	tasks, err := t.ListTasks("tasks.json")
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(tasks)

	tasks, err = t.DeleteAllTasks("tasks.json")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(tasks)
}
