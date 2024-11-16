// File: source_code.go
package main

import (
	"bufio"       // Provides buffered I/O for reading user input
	"encoding/json" // Enables JSON encoding/decoding for file operations
	"fmt"          // Provides functions for formatted I/O
	"os"           // Provides functions to interact with the file system
	"strconv"      // Provides functions to convert strings to integers and vice versa
	"strings"      // Provides string manipulation functions
)

// Task represents a single task in the tracker.
type Task struct {
	ID int `json:"id"` // Unique identifier for the task
	Title string `json:"title"` // Description/title of the task
	Status string `json:"status"` // Status of the task: "todo", "in-progress", or "done"
}

// TaskList manages a slice of Task and the next ID.
type TaskList struct {
	Tasks  []Task `json:"tasks"`   // List of tasks
	NextID int    `json:"next_id"` // Next ID to ensure task IDs are unique
}

// Global task list to store all tasks and manage IDs
var taskList TaskList

func main() {
	// Load tasks from file at the start of the program
	if err := loadTasksFromFile("tasks.json"); err != nil {
		// If there's an error loading the file, print a warning
		fmt.Println("Warning: Could not load tasks:", err)
	}

	// Ensure tasks are saved to the file when the program exits
	defer saveTasksToFile("tasks.json")

	// Scanner to read user input from the terminal
	scanner := bufio.NewScanner(os.Stdin)

	// Main loop to display the menu and process user commands
	for {
		displayMenu() // Show the CLI menu to the user

		// Read the user's choice
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text()) // Clean up whitespace

		// Handle the user's choice
		switch choice {
		case "1":
			addTask(scanner) // Add a new task
		case "2":
			updateTaskStatus(scanner) // Update the status of an existing task
		case "3":
			listTasks(scanner) // List tasks based on status
		case "4":
			fmt.Println("Exiting... Goodbye!")
			return // Exit the program
		default:
			// Handle invalid choices
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

// displayMenu shows the CLI options to the user.
func displayMenu() {
	fmt.Println("\nTask Tracker CLI")
	fmt.Println("1. Add Task")       // Option to add a new task
	fmt.Println("2. Update Task Status") // Option to change the status of a task
	fmt.Println("3. List Tasks")         // Option to view tasks
	fmt.Println("4. Exit")               // Option to exit the program
	fmt.Print("Choose an option: ")      // Prompt the user to choose
}

// addTask allows the user to add a new task to the list.
func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ") // Ask the user for a task title
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text()) // Read and clean up the title

	// Validate that the title is not empty
	if title == "" {
		fmt.Println("Error: Task title cannot be empty.")
		return
	}

	// Create a new task with a unique ID
	newTask := Task{
		ID:     taskList.NextID, // Assign the next available ID
		Title:  title,           // Set the title provided by the user
		Status: "todo",          // Default status for new tasks
	}

	// Add the new task to the list and increment the next ID
	taskList.Tasks = append(taskList.Tasks, newTask)
	taskList.NextID++

	// Confirm the task has been added
	fmt.Printf("Task '%s' added with ID %d.\n", title, newTask.ID)
}

// updateTaskStatus allows the user to change the status of an existing task.
func updateTaskStatus(scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to update: ") // Ask for the task ID
	scanner.Scan()
	idStr := strings.TrimSpace(scanner.Text()) // Read and clean up the ID input
	id, err := strconv.Atoi(idStr)            // Convert the ID to an integer
	if err != nil {
		fmt.Println("Error: Invalid task ID.") // Handle invalid input
		return
	}

	// Search for the task by ID
	for i, task := range taskList.Tasks {
		if task.ID == id {
			// If task is found, ask for the new status
			fmt.Print("Enter new status (todo, in-progress, done): ")
			scanner.Scan()
			status := strings.TrimSpace(scanner.Text()) // Read and clean up the status

			// Validate that the status is valid
			if !isValidStatus(status) {
				fmt.Println("Error: Invalid status. Valid statuses are: todo, in-progress, done.")
				return
			}

			// Update the task's status
			taskList.Tasks[i].Status = status
			fmt.Printf("Task ID %d updated to status '%s'.\n", id, status)
			return
		}
	}

	// If the task ID was not found, notify the user
	fmt.Println("Error: Task not found.")
}

// listTasks displays tasks filtered by their status.
func listTasks(scanner *bufio.Scanner) {
	fmt.Print("Enter status to filter by (todo, in-progress, done, all): ")
	scanner.Scan()
	status := strings.TrimSpace(scanner.Text()) // Read and clean up the status input

	// Validate that the status filter is valid
	if status != "todo" && status != "in-progress" && status != "done" && status != "all" {
		fmt.Println("Error: Invalid status filter.")
		return
	}

	// Display tasks that match the filter
	fmt.Println("\nTasks:")
	count := 0
	for _, task := range taskList.Tasks {
		// Include tasks that match the filter or show all
		if status == "all" || task.Status == status {
			fmt.Printf("ID: %d | Title: %s | Status: %s\n", task.ID, task.Title, task.Status)
			count++
		}
	}

	// Notify if no tasks are found
	if count == 0 {
		fmt.Println("No tasks found.")
	}
}