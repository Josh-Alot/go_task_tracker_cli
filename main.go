package main

import (
	"fmt"

	t "github.com/go_task_tracker/tasks"
)

func main() {
	//tasks := []t.Task{
	//	{ID: 1, Description: "Go shopping", Status: 0},
	//	{ID: 2, Description: "Sweep house", Status: 1},
	//	{ID: 3, Description: "Feed kitten", Status: 2},
	//}
	//
	//err := t.CreateTask(tasks, "tasks.json")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}

	tasks, _ := t.ListIncompleteTasks("tasks.json")
	fmt.Println(tasks)
}
