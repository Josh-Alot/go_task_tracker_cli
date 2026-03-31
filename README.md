# Task Tracker CLI

A simple command-line tool to track and manage your tasks. Built with Go.

## Overview

Task Tracker CLI allows you to manage your to-do list directly from the terminal. Tasks are stored in a local `tasks.json` file in the current directory.

### How It Works

The application uses Go's native `flag` package for CLI argument parsing and the `os`/`encoding/json` packages for file operations. No external dependencies are required.

- **Data Storage**: Tasks are persisted in `tasks.json` (created automatically if it doesn't exist)
- **Task Properties**: Each task has an ID, description, status (todo/in_progress/done), createdAt, and updatedAt timestamp

### Architecture

```
main.go       - CLI entry point, handles command routing
tasks/tasks.go - Core task operations (CRUD, list by status)
```

---

## Installation & Compilation

### Prerequisites

- Go 1.22+ installed

### Build the CLI

```bash
go build -o todo-cli .
```

This produces an executable named `todo-cli` in the current directory.

### Install Globally (optional)

```bash
# Move to your PATH
mv todo-cli /usr/local/bin/

# Or add to PATH temporarily
export PATH=$PATH:/path/to/directory
```

---

## Usage

### Add a Task

```bash
./todo-cli add "Buy groceries"
# Output: Task added successfully
```

### List All Tasks

```bash
./todo-cli list
```

### List Tasks by Status

```bash
./todo-cli list todo        # Tasks not yet started
./todo-cli list in_progress # Tasks in progress
./todo-cli list done       # Completed tasks
./todo-cli list incomplete # Both todo and in_progress
```

### Update a Task

```bash
# Update status
./todo-cli update 1 todo
./todo-cli update 1 in_progress
./todo-cli update 1 done

# Update description
./todo-cli update 1 description "New description"
```

### Delete a Task

```bash
./todo-cli delete 1
```

---

## Example

```bash
# Add tasks
$ ./todo-cli add "Buy groceries"
Task added successfully

$ ./todo-cli add "Walk the dog"
Task added successfully

# List all
$ ./todo-cli list
ID	Description	Status
1	Buy groceries	todo
2	Walk the dog	todo

# Mark one as in progress
$ ./todo-cli update 1 in_progress
Task updated successfully

# List by status
$ ./todo-cli list in_progress
ID	Description	Status
1	Buy groceries	in_progress

# Mark as done
$ ./todo-cli update 1 done
Task updated successfully

# Delete
$ ./todo-cli delete 2
```

---

## Data File

Tasks are stored in `tasks.json` in the directory where you run the command:

```json
[
	{
		"ID": 1,
		"Description": "Buy groceries",
		"Status": "done",
		"CreatedAt": "2026-03-31T10:00:00Z",
		"UpdatedAt": "2026-03-31T12:00:00Z"
	}
]
```