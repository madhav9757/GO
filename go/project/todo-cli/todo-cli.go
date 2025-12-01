package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Colors
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

// Todo struct
type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Category  string    `json:"category"`
	DueDate   string    `json:"due_date"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

var todos []Todo
var dataFile = "todo_advanced.json"

// ---------- UTIL: Clean Input Reader ----------
var reader = bufio.NewReader(os.Stdin)

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// ---------- Load Todos ----------
func loadTodos() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			todos = []Todo{}
			return
		}
		panic(err)
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&todos)
}

// ---------- Save Todos (Pretty JSON) ----------
func saveTodos() {
	file, err := os.Create(dataFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pretty, _ := json.MarshalIndent(todos, "", "  ")
	file.Write(pretty)
}

// ---------- Add Task ----------
func addTask(task, category, due string) {
	id := 1
	if len(todos) > 0 {
		id = todos[len(todos)-1].ID + 1
	}

	todo := Todo{
		ID:        id,
		Task:      task,
		Category:  category,
		DueDate:   due,
		Done:      false,
		CreatedAt: time.Now(),
	}

	todos = append(todos, todo)
	fmt.Println(Green+"✔ Task added:"+Reset, task)
}

// ---------- List Tasks ----------
func listTasks() {
	if len(todos) == 0 {
		fmt.Println(Yellow + "No tasks found." + Reset)
		return
	}

	fmt.Println("\n" + Blue + "-------- Your Tasks --------" + Reset)
	for _, t := range todos {
		status := Red + "❌" + Reset
		if t.Done {
			status = Green + "✔️" + Reset
		}

		fmt.Printf(
			"%d. [%s] %s  (%s, Due: %s)\n",
			t.ID, status, t.Task, t.Category, t.DueDate,
		)
	}
}

// ---------- Mark Done ----------
func markDone(id int) {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Done = true
			fmt.Println(Green+"✔ Marked as Done:"+Reset, t.Task)
			return
		}
	}
	fmt.Println(Red + "Task not found!" + Reset)
}

// ---------- Delete Task ----------
func deleteTask(id int) {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println(Red+"🗑 Deleted:"+Reset, t.Task)
			return
		}
	}
	fmt.Println(Red + "Task not found!" + Reset)
}

// ---------- MAIN ----------
func main() {
	loadTodos()

	for {
		fmt.Println("\n" + Blue + "------- TODO CLI MENU -------" + Reset)
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		choice := input("Enter choice: ")

		switch choice {
		case "1":
			task := input("Enter task: ")
			category := input("Enter category (Work/Study/Personal): ")
			due := input("Enter due date (YYYY-MM-DD): ")
			addTask(task, category, due)

		case "2":
			listTasks()

		case "3":
			idStr := input("Enter task ID to mark done: ")
			var id int
			fmt.Sscanf(idStr, "%d", &id)
			markDone(id)

		case "4":
			idStr := input("Enter task ID to delete: ")
			var id int
			fmt.Sscanf(idStr, "%d", &id)
			deleteTask(id)

		case "5":
			saveTodos()
			fmt.Println(Green + "Saved. Goodbye!" + Reset)
			return

		default:
			fmt.Println(Red + "Invalid option, try again." + Reset)
		}
	}
}
