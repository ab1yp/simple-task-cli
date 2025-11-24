# Task Manager CLI (Go)

A simple and clean **Command-Line Task Manager** built in **Go**, featuring full CRUD functionality and JSON-based storage.  
Designed as a learning project to practice Go fundamentals including structs, slices, file handling, and CLI input.

---

## Features

- **Add Task** – Create tasks with a title, content, and auto-generated ID  
- **Show Tasks** – List all saved tasks in a clean formatted layout  
- **Update Task** – Edit task details by ID  
- **Delete Task** – Remove tasks by ID  
- **Search Tasks** – Search tasks by ID, title, or content  
- **Clear Tasks** – Delete all tasks and reset storage

---

## Storage

Tasks are stored in a JSON file:

tasks.json

yaml
Copy code

Using Go’s built-in `encoding/json` package for reading and writing data.

---

## Usage

Clone the repository:

```bash
git clone <your-repo-url>
cd task-manager-cli
Run the application:

go run main.go

Example Output

==========================

        TASK MANAGER

 ==========================



Options:
        1) Add task
        1) Add task
        2) Show all tasks
        3) Update task
        4) Delete task
        5) Search task
        6) Clear all tasks

Choice and enter number below:


```
License
MIT License – free to use, modify, and distribute.
