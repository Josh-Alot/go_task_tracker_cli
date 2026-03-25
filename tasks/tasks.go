package tasks

import (
	"encoding/json"
	"log"
	"os"
)

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (status Status) String() string {
	switch status {
	case Todo:
		return "todo"
	case InProgress:
		return "in_progress"
	case Done:
		return "done"
	default:
		return "unknown"
	}
}

func (status Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.String())
}

func (status *Status) UnmarshalJSON(data []byte) error {
	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch str {
	case "todo":
		*status = Todo
	case "in_progress":
		*status = InProgress
	case "done":
		*status = Done
	default:
		*status = Todo
	}

	return nil
}

type Task struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   string
	UpdatedAt   string
}

func CreateTask(newTasks []Task, filePath string) error {
	var err error
	var existingTasks []Task

	// check if the file already exists
	// if the file doesn't exist, it will create a new file
	// if the file exists, it recover the content into an slice
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
	} else {
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		// converts the JSON content into a slice
		json.Unmarshal(data, &existingTasks)
	}

	existingTasks = append(existingTasks, newTasks...)

	// convert slice content to JSON content
	data, err := json.MarshalIndent(existingTasks, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(filePath, data, 0644)
	return err
}

func ListTasks(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	tasks := []Task{}

	json.Unmarshal(data, &tasks)

	return tasks, err
}

func ListIncompleteTasks(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	json.Unmarshal(data, &tasks)

	var incompleteTasks []Task
	for _, task := range tasks {
		if task.Status != 2 {
			incompleteTasks = append(incompleteTasks, task)
		}
	}

	return incompleteTasks, err
}

func ListCompleteTasks(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	json.Unmarshal(data, &tasks)

	var completeTasks []Task
	for _, task := range tasks {
		if task.Status == 2 {
			completeTasks = append(completeTasks, task)
		}
	}

	return completeTasks, err
}

func ListTodoTasks(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	json.Unmarshal(data, &tasks)

	var todoTasks []Task
	for _, task := range tasks {
		if task.Status == 0 {
			todoTasks = append(todoTasks, task)
		}
	}

	return todoTasks, err
}

func ListInProgressTasks(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	json.Unmarshal(data, &tasks)

	var inProgressTasks []Task
	for _, task := range tasks {
		if task.Status == 1 {
			inProgressTasks = append(inProgressTasks, task)
		}
	}

	return inProgressTasks, err
}
