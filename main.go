package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
)

var FileName string = "tasks.json"
var Tasks []Task

type Task struct {
	Title   string
	Content string
	TaskId  string
}

func generateTaskId() string {
	return fmt.Sprintf("%04d", rand.IntN(10000))
}

func loadTasks() {

	data, err := os.ReadFile(FileName)
	if err == nil && len(data) > 0 {
		json.Unmarshal(data, &Tasks)
	}
}
func saveTasks() {
	data, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(FileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func addTask() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input task title below: ")
	scanner.Scan()
	taskTitle := scanner.Text()

	fmt.Println("input task content below: ")
	scanner.Scan()
	taskContent := scanner.Text()

	if strings.TrimSpace(taskTitle) == "" || strings.TrimSpace(taskContent) == "" {
		fmt.Println("Invalid Values!.")
		return
	}

	taskId := generateTaskId()
	task := Task{
		Title:   taskTitle,
		Content: taskContent,
		TaskId:  taskId,
	}

	Tasks = append(Tasks, task)

	saveTasks()

	fmt.Println(">>>>> Task added successfully and saved to file.")

}

func showTasks() {
	loadTasks()
	for _, t := range Tasks {
		fmt.Println(">>>>> TASK INFO <<<<<")
		fmt.Println("ID:", t.TaskId)
		fmt.Println("Title:", t.Title)
		fmt.Println("Content:", t.Content)
		fmt.Println("---------------------")
	}
}

func updateTask() {
	loadTasks()
	var taskId string
	fmt.Println("Please enter task taskId below: ")
	fmt.Scanln(&taskId)
	
	scanner := bufio.NewScanner(os.Stdin)

	for task := range Tasks {
		if Tasks[task].TaskId == taskId {

			fmt.Println("Enter new Task ID below: ")
			scanner.Scan()
			newTaskId := scanner.Text()
			
			fmt.Println("Enter new Task title below: ")
			scanner.Scan()
			newTitle := scanner.Text()

			fmt.Println("Enter new Task content below: ")
			scanner.Scan()
			newContent := scanner.Text()

			if strings.TrimSpace(newTaskId) == "" || strings.TrimSpace(newTitle) == "" || strings.TrimSpace(newContent) == "" {
				fmt.Println("Invalid Values!.")
				return
			}

			Tasks[task].TaskId = newTaskId
			Tasks[task].Title = newTitle
			Tasks[task].Content = newContent

			saveTasks()
			fmt.Println("Task updated successfully!")
			return
		}
	}
}

func deleteTask() bool {

	loadTasks()

	var taskId string
	fmt.Println("Please enter task taskId below: ")
	fmt.Scanln(&taskId)

	for task := range Tasks {
		if Tasks[task].TaskId == taskId {
			Tasks = append(Tasks[:task], Tasks[task+1:]...)
			saveTasks()
			fmt.Println(">>>>> Task deleted successfully.")
			return true
		}
	}

	return false
}

func searchTask() {
scanner := bufio.NewScanner(os.Stdin)
fmt.Println("Enter search keyword (ID, title, or content) below : ")
scanner.Scan()
keyword := scanner.Text()
found := false
fmt.Println(">>>>> Search INFO <<<<<")

for task := range Tasks {
	if strings.Contains(strings.ToLower(Tasks[task].TaskId), strings.ToLower(keyword)) ||
	 strings.Contains(strings.ToLower(Tasks[task].Title), strings.ToLower(keyword)) ||
	 strings.Contains(strings.ToLower(Tasks[task].Content), strings.ToLower(keyword)) {
		fmt.Println(">>>>> TASK INFO <<<<<")
		fmt.Println("ID:", Tasks[task].TaskId)
    fmt.Println("Title:", Tasks[task].Title)
    fmt.Println("Content:", Tasks[task].Content)
    fmt.Println("---------------------")
		found = true
	}
	} 
	if !found {
    fmt.Println(">>>>> No task found.")
	}
}



func CAT() {
	Tasks = []Task{}
	saveTasks()
}
func main() {
	for {
		
		fmt.Println(" ==========================")
		fmt.Println("")
		fmt.Println("        TASK MANAGER")
		fmt.Println("")
		fmt.Println(" ==========================")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("	1) Add task")
		fmt.Println("	1) Add task")
		fmt.Println("	2) Show all tasks")
		fmt.Println("	3) Update task")
		fmt.Println("	4) Delete task")
		fmt.Println("	5) Search task")
		fmt.Println("	6) Clear all tasks")
		fmt.Println("")
		var inputValue int
		fmt.Println("Choice and enter number below: ")
		fmt.Scanln(&inputValue)

		switch inputValue {
		case 1:
			addTask()
		case 2:
			showTasks()
		case 3:
			updateTask()
		case 4:
			deleteTask()
		case 5:
			searchTask()
		case 6:
			CAT()
		default:
			fmt.Println(">>>>> Invalid choice")
		}

	}

}
