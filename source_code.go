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