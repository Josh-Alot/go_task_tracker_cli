package tasks

import (
	"encoding/json"
	"log"
	"os"
)

type Status struct {
	Name string
}

type Task struct {
	ID          int
	Description string
	Status
	CreatedAt string
	UpdatedAt string
}

func CreateTask(newTasks []Task) {
	filePath := "tasks.json"
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

		json.Unmarshal(data, &existingTasks)
	}

	existingTasks = append(existingTasks, newTasks...)

	// convert struct to JSON
	data, err := json.MarshalIndent(existingTasks, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(filePath, data, 0644)
}
