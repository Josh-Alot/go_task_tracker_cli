package tasks

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
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

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		nextID := getNextID(nil)
		for i := range newTasks {
			newTasks[i].ID = nextID
			nextID++
		}
	} else {
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(data, &existingTasks)

		nextID := getNextID(existingTasks)
		for i := range newTasks {
			newTasks[i].ID = nextID
			nextID++
		}
	}

	existingTasks = append(existingTasks, newTasks...)

	data, err := json.MarshalIndent(existingTasks, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(filePath, data, 0644)
	return err
}

func getNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	maxTask := slices.MaxFunc(tasks, func(a, b Task) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return maxTask.ID + 1
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
		if task.Status != Done {
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
		if task.Status == Done {
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
		if task.Status == Todo {
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
		if task.Status == InProgress {
			inProgressTasks = append(inProgressTasks, task)
		}
	}

	return inProgressTasks, err
}

func UpdateTaskDescription(filePath string, id int, description string) ([]Task, error) {
	found := false
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	json.Unmarshal(data, &tasks)

	for i, _ := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("Task %d not found", id)
	}

	data, err = json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return nil, err
	}

	os.WriteFile(filePath, data, 0644)

	return tasks, err
}

func UpdateTaskStatus(filePath string, id int, status Status) ([]Task, error) {
	found := false
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	json.Unmarshal(data, &tasks)

	for i, _ := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("Task %d not found", id)
	}

	data, err = json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return nil, err
	}

	os.WriteFile(filePath, data, 0644)

	return tasks, err
}

func DeleteTask(filePath string, id int) ([]Task, error) {
	found := false
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	json.Unmarshal(data, &tasks)

	for i, _ := range tasks {
		if tasks[i].ID == id {
			tasks = slices.Delete(tasks, i, i+1)

			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("Task %d not found", id)
	}

	data, err = json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return nil, err
	}

	os.WriteFile(filePath, data, 0644)

	return tasks, err
}

func DeleteAllTasks(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	json.Unmarshal(data, &tasks)

	if len(tasks) < 1 {
		return nil, fmt.Errorf("Tasks not found")
	}

	data, err = json.Marshal([]Task{})
	if err != nil {
		return nil, err
	}

	os.WriteFile(filePath, data, 0644)

	return nil, err
}
