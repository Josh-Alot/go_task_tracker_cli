package tasks

import (
	"encoding/json"
	"os"
	"testing"
)

func TestCreateTask(t *testing.T) {
	testFile := "test.json"
	defer os.Remove(testFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: 0},
	}

	err := CreateTask(tasks, testFile)
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Error("Expected tasks file to be created")
	}
}

func TestCreateMultipletTasks(t *testing.T) {
	testFile := "test.json"
	defer os.Remove(testFile)

	// Creates a new test file
	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: 0},
		{ID: 2, Description: "Task 2", Status: 1},
	}

	err := CreateTask(tasks, testFile)
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Test the new file
	var savedTasks []Task
	if err := json.Unmarshal(data, &savedTasks); err != nil {
		t.Fatalf("Failed to unmarshal json file: %v", err)
	}

	if len(savedTasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(savedTasks))
	}

	if savedTasks[0].Description != "Task 1" || savedTasks[1].Description != "Task 2" {
		t.Errorf("Task description didn't match")
	}

	// Creating a third task and test the file
	newTasks := []Task{
		{ID: 3, Description: "Task 3", Status: 0},
	}

	err = CreateTask(newTasks, testFile)
	if err != nil {
		t.Fatalf("Failed including new task: %v", err)
	}

	data, err = os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if err := json.Unmarshal(data, &savedTasks); err != nil {
		t.Fatalf("Failed to unmarshal json file: %v", err)
	}

	if len(savedTasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(savedTasks))
	}

	if savedTasks[2].Description != "Task 3" {
		t.Errorf("Task description didn't match")
	}
}

func TestListTasks(t *testing.T) {
	testFile := "test.json"
	defer os.Remove(testFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: 0},
		{ID: 2, Description: "Task 2", Status: 1},
	}

	// first creates the test file
	CreateTask(tasks, testFile)

	// check if the file returns
	listTasks, err := ListTasks(testFile)

	if _, err = os.Stat(testFile); os.IsNotExist(err) {
		t.Errorf("Tasks file not found")
	}

	if len(listTasks) != 2 {
		t.Errorf("Expected 3 tasks, got %d", len(listTasks))
	}

	if listTasks[0].Description != "Task 1" || listTasks[1].Description != "Task 2" {
		t.Errorf("Task description didn't match")
	}
}

func TestListTasksFileNotDoesntExist(t *testing.T) {
	testFile := "test.json"
	defer os.Remove(testFile)

	listTasks, err := ListTasks("test2.json")
	if err == nil {
		t.Errorf("Expected an error for non-existent file, got nil")
	}

	if listTasks != nil {
		t.Errorf("Expected nil tasks, got a slice with tasks")
	}
}

func TestListIncompleteTasks(t *testing.T) {
	testFile := "test.json"
	defer os.Remove(testFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: 0},
		{ID: 2, Description: "Task 2", Status: 1},
		{ID: 3, Description: "Task 3", Status: 2},
	}
	CreateTask(tasks, testFile)

	tasks, err := ListIncompleteTasks(testFile)
	if err != nil {
		t.Errorf("Failed to filter incomplete tasks: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestListCompleteTasks(t *testing.T) {
	testFile := "test.json"
	defer os.Remove(testFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: 0},
		{ID: 2, Description: "Task 2", Status: 1},
		{ID: 3, Description: "Task 3", Status: 2},
	}
	CreateTask(tasks, testFile)

	tasks, err := ListCompleteTasks(testFile)
	if err != nil {
		t.Errorf("Failed to filter complete tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestListTodoTasks(t *testing.T) {
	testFile := "test.json"
	defer os.Remove(testFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: 0},
		{ID: 2, Description: "Task 2", Status: 1},
		{ID: 3, Description: "Task 3", Status: 2},
	}
	CreateTask(tasks, testFile)

	tasks, err := ListTodoTasks(testFile)
	if err != nil {
		t.Errorf("Failed to filter incomplete tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestListInProgressTasks(t *testing.T) {
	testFile := "test.json"
	defer os.Remove(testFile)

	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: 0},
		{ID: 2, Description: "Task 2", Status: 1},
		{ID: 3, Description: "Task 3", Status: 2},
	}
	CreateTask(tasks, testFile)

	tasks, err := ListInProgressTasks(testFile)
	if err != nil {
		t.Errorf("Failed to filter incomplete tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}
