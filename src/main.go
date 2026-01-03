package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////
// Data Models
//////////////////////////////////////////////////////

// Task represents a single task entity in the system.
type Task struct {
	ID       int    `json:"id"`       // Unique identifier
	Title    string `json:"title"`    // Task title/description
	Status   string `json:"status"`   // todo | in-progress | done
	Priority string `json:"priority"` // High | Medium | Low
}

// TaskList holds all tasks and manages ID generation.
type TaskList struct {
	Tasks  []Task `json:"tasks"`   // Collection of tasks
	NextID int    `json:"next_id"` // Auto-incrementing ID
}

//////////////////////////////////////////////////////
// Global Variables
//////////////////////////////////////////////////////

var (
	taskList TaskList
	filePath = "tasks.json"
)

//////////////////////////////////////////////////////
// Application Entry Point
//////////////////////////////////////////////////////

func main() {
	// Load existing tasks from file (if any)
	if err := loadTasksFromFile(filePath); err != nil {
		fmt.Println("Warning: Failed to load tasks:", err)
	}

	// Ensure tasks are saved when program exits
	defer saveTasksToFile(filePath)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		displayMenu()
		scanner.Scan()

		switch strings.TrimSpace(scanner.Text()) {
		case "1":
			addTask(scanner)
		case "2":
			updateTaskStatus(scanner)
		case "3":
			listTasks(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			searchTasks(scanner)
		case "6":
			sortTasksMenu(scanner)
		case "7":
			fmt.Println("Exiting... Goodbye 👋")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

//////////////////////////////////////////////////////
// Menu & UI
//////////////////////////////////////////////////////

// displayMenu prints the main CLI menu.
func displayMenu() {
	fmt.Println("\n🚀 Enhanced Task Tracker CLI")
	fmt.Println("1. Add Task")
	fmt.Println("2. Update Task Status")
	fmt.Println("3. List Tasks")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Search Tasks")
	fmt.Println("6. Sort Tasks")
	fmt.Println("7. Exit")
	fmt.Print("Choose an option: ")
}

//////////////////////////////////////////////////////
// Core Features
//////////////////////////////////////////////////////

// addTask creates a new task and adds it to the task list.
func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())

	if title == "" {
		fmt.Println("Error: Task title cannot be empty.")
		return
	}

	fmt.Print("Enter priority (High, Medium, Low): ")
	scanner.Scan()
	priority := strings.TrimSpace(scanner.Text())

	if !isValidPriority(priority) {
		fmt.Println("Error: Invalid priority.")
		return
	}

	task := Task{
		ID:       taskList.NextID,
		Title:    title,
		Status:   "todo",
		Priority: priority,
	}

	taskList.Tasks = append(taskList.Tasks, task)
	taskList.NextID++

	fmt.Printf("✅ Task added successfully (ID: %d)\n", task.ID)
}

// updateTaskStatus updates the status of an existing task.
func updateTaskStatus(scanner *bufio.Scanner) {
	id := readTaskID(scanner)
	if id == -1 {
		return
	}

	for i := range taskList.Tasks {
		if taskList.Tasks[i].ID == id {
			fmt.Print("Enter new status (todo, in-progress, done): ")
			scanner.Scan()
			status := strings.TrimSpace(scanner.Text())

			if !isValidStatus(status) {
				fmt.Println("Error: Invalid status.")
				return
			}

			taskList.Tasks[i].Status = status
			fmt.Println("✅ Task status updated.")
			return
		}
	}

	fmt.Println("Error: Task not found.")
}

// listTasks prints tasks based on status filter.
func listTasks(scanner *bufio.Scanner) {
	fmt.Print("Filter by status (todo, in-progress, done, all): ")
	scanner.Scan()
	filter := strings.TrimSpace(scanner.Text())

	if !isValidFilter(filter) {
		fmt.Println("Error: Invalid filter.")
		return
	}

	fmt.Println("\n📋 Tasks:")
	found := false

	for _, task := range taskList.Tasks {
		if filter == "all" || task.Status == filter {
			printTask(task)
			found = true
		}
	}

	if !found {
		fmt.Println("No tasks found.")
	}
}

// deleteTask removes a task after confirmation.
func deleteTask(scanner *bufio.Scanner) {
	id := readTaskID(scanner)
	if id == -1 {
		return
	}

	for i, task := range taskList.Tasks {
		if task.ID == id {
			fmt.Printf("Delete task \"%s\"? (yes/no): ", task.Title)
			scanner.Scan()

			if strings.TrimSpace(scanner.Text()) == "yes" {
				taskList.Tasks = append(taskList.Tasks[:i], taskList.Tasks[i+1:]...)
				fmt.Println("🗑️ Task deleted.")
			} else {
				fmt.Println("Deletion cancelled.")
			}
			return
		}
	}

	fmt.Println("Error: Task not found.")
}

// searchTasks finds tasks by keyword in title.
func searchTasks(scanner *bufio.Scanner) {
	fmt.Print("Enter keyword: ")
	scanner.Scan()
	keyword := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if keyword == "" {
		fmt.Println("Error: Keyword cannot be empty.")
		return
	}

	fmt.Println("\n🔍 Search Results:")
	found := false

	for _, task := range taskList.Tasks {
		if strings.Contains(strings.ToLower(task.Title), keyword) {
			printTask(task)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching tasks found.")
	}
}

//////////////////////////////////////////////////////
// Sorting
//////////////////////////////////////////////////////

// sortTasks sorts tasks by given field and order.
func sortTasks(by string, asc bool) {
	priorityRank := map[string]int{"High": 3, "Medium": 2, "Low": 1}

	sort.Slice(taskList.Tasks, func(i, j int) bool {
		var less bool

		switch by {
		case "id":
			less = taskList.Tasks[i].ID < taskList.Tasks[j].ID
		case "title":
			less = taskList.Tasks[i].Title < taskList.Tasks[j].Title
		case "status":
			less = taskList.Tasks[i].Status < taskList.Tasks[j].Status
		case "priority":
			less = priorityRank[taskList.Tasks[i].Priority] <
				priorityRank[taskList.Tasks[j].Priority]
		}

		if asc {
			return less
		}
		return !less
	})

	fmt.Println("✅ Tasks sorted successfully.")
}

// sortTasksMenu handles sorting options.
func sortTasksMenu(scanner *bufio.Scanner) {
	options := map[string]struct {
		by  string
		asc bool
	}{
		"1": {"id", true},
		"2": {"id", false},
		"3": {"title", true},
		"4": {"title", false},
		"5": {"status", true},
		"6": {"status", false},
		"7": {"priority", true},
		"8": {"priority", false},
	}

	fmt.Println("Sort by:")
	fmt.Println("1-8 (ID / Title / Status / Priority)")
	fmt.Print("Choose: ")
	scanner.Scan()

	if opt, ok := options[strings.TrimSpace(scanner.Text())]; ok {
		sortTasks(opt.by, opt.asc)
	} else {
		fmt.Println("Invalid sort option.")
	}
}

//////////////////////////////////////////////////////
// Persistence
//////////////////////////////////////////////////////

// saveTasksToFile persists tasks to disk.
func saveTasksToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	_ = json.NewEncoder(file).Encode(taskList)
}

// loadTasksFromFile loads tasks from disk.
func loadTasksFromFile(filename string) error {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		taskList = TaskList{NextID: 1}
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&taskList)
}

//////////////////////////////////////////////////////
// Helpers & Validation
//////////////////////////////////////////////////////

func readTaskID(scanner *bufio.Scanner) int {
	fmt.Print("Enter task ID: ")
	scanner.Scan()

	id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || id < 1 {
		fmt.Println("Error: Invalid task ID.")
		return -1
	}
	return id
}

func printTask(task Task) {
	fmt.Printf("ID:%d | %s | %s | Priority:%s\n",
		task.ID, task.Title, task.Status, task.Priority)
}

func isValidStatus(s string) bool {
	return s == "todo" || s == "in-progress" || s == "done"
}

func isValidPriority(p string) bool {
	return p == "High" || p == "Medium" || p == "Low"
}

func isValidFilter(f string) bool {
	return f == "todo" || f == "in-progress" || f == "done" || f == "all"
}
