### **Imports**
```go
import (
	"bufio"         // Provides buffered I/O for reading user input
	"encoding/json" // Enables JSON encoding/decoding for file operations
	"fmt"           // Provides functions for formatted I/O
	"os"            // Provides functions to interact with the file system
	"strconv"       // Provides functions to convert strings to integers and vice versa
	"strings"       // Provides string manipulation functions
)
```
- Importing necessary Go packages:
  - `bufio`: Handles user input with a scanner.
  - `encoding/json`: Facilitates JSON reading/writing for saving and loading tasks.
  - `fmt`: Used for printing and formatting text.
  - `os`: For file system operations (e.g., open, create).
  - `strconv`: Converts strings to integers and vice versa.
  - `strings`: Provides string manipulation (e.g., trimming, splitting).

---

### **Task Struct**
```go
type Task struct {
	ID     int    `json:"id"`     // Unique identifier for the task
	Title  string `json:"title"`  // Description/title of the task
	Status string `json:"status"` // Status of the task: "todo", "in-progress", or "done"
}
```
- Defines the structure of a task:
  - `ID`: A unique identifier (integer).
  - `Title`: A string representing the task description.
  - `Status`: Indicates the task's progress ("todo", "in-progress", "done").

---

### **TaskList Struct**
```go
type TaskList struct {
	Tasks  []Task `json:"tasks"`   // List of tasks
	NextID int    `json:"next_id"` // Next ID to ensure task IDs are unique
}
```
- Represents a collection of tasks and keeps track of the next available ID.
- `Tasks`: A slice of `Task` structs.
- `NextID`: A counter ensuring task IDs remain unique.

---

### **Global Variable**
```go
var taskList TaskList
```
- Declares a global instance of `TaskList` to hold tasks in memory.

---

### **Main Function**
```go
func main() {
	if err := loadTasksFromFile("tasks.json"); err != nil {
		fmt.Println("Warning: Could not load tasks:", err)
	}
	defer saveTasksToFile("tasks.json")
```
- **Load Tasks**: Tries to load tasks from a `tasks.json` file when the program starts.
  - If loading fails, a warning is displayed.
- **Defer Save**: Ensures that tasks are saved back to the file when the program exits, even if it exits early.

```go
	scanner := bufio.NewScanner(os.Stdin)
```
- Initializes a `Scanner` for reading user input from the terminal.

```go
	for {
		displayMenu()
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
```
- Enters an infinite loop, displaying the menu and waiting for user input.
- Reads the user's choice and trims any extra spaces.

```go
		switch choice {
		case "1":
			addTask(scanner)
		case "2":
			updateTaskStatus(scanner)
		case "3":
			listTasks(scanner)
		case "4":
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
```
- Depending on the user’s input:
  - `"1"`: Calls `addTask()` to add a new task.
  - `"2"`: Calls `updateTaskStatus()` to update an existing task’s status.
  - `"3"`: Calls `listTasks()` to display tasks.
  - `"4"`: Exits the loop and program.
  - Default: Prints an error message for invalid input.

---

### **Menu Display**
```go
func displayMenu() {
	fmt.Println("\nTask Tracker CLI")
	fmt.Println("1. Add Task")
	fmt.Println("2. Update Task Status")
	fmt.Println("3. List Tasks")
	fmt.Println("4. Exit")
	fmt.Print("Choose an option: ")
}
```
- Displays the CLI menu.

---

### **Add Task**
```go
func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())
```
- Prompts the user for a task title and reads the input.

```go
	if title == "" {
		fmt.Println("Error: Task title cannot be empty.")
		return
	}
```
- Validates that the title is not empty.

```go
	newTask := Task{
		ID:     taskList.NextID,
		Title:  title,
		Status: "todo",
	}
	taskList.Tasks = append(taskList.Tasks, newTask)
	taskList.NextID++
	fmt.Printf("Task '%s' added with ID %d.\n", title, newTask.ID)
}
```
- Creates a new task with the next available ID and default status ("todo").
- Appends the task to the list, increments `NextID`, and confirms the addition.

---

### **Update Task Status**
```go
func updateTaskStatus(scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to update: ")
	scanner.Scan()
	idStr := strings.TrimSpace(scanner.Text())
	id, err := strconv.Atoi(idStr)
```
- Prompts the user for a task ID and converts it from string to integer.

```go
	if err != nil {
		fmt.Println("Error: Invalid task ID.")
		return
	}
```
- Validates the input. If it's not a valid number, an error is shown.

```go
	for i, task := range taskList.Tasks {
		if task.ID == id {
			fmt.Print("Enter new status (todo, in-progress, done): ")
			scanner.Scan()
			status := strings.TrimSpace(scanner.Text())
			if !isValidStatus(status) {
				fmt.Println("Error: Invalid status. Valid statuses are: todo, in-progress, done.")
				return
			}
			taskList.Tasks[i].Status = status
			fmt.Printf("Task ID %d updated to status '%s'.\n", id, status)
			return
		}
	}
	fmt.Println("Error: Task not found.")
}
```
- Searches for the task by ID. If found:
  - Prompts for a new status and validates it.
  - Updates the task's status.
- If the task is not found, it displays an error.

---

### **List Tasks**
```go
func listTasks(scanner *bufio.Scanner) {
	fmt.Print("Enter status to filter by (todo, in-progress, done, all): ")
	scanner.Scan()
	status := strings.TrimSpace(scanner.Text())
```
- Prompts the user for a status filter.

```go
	if status != "todo" && status != "in-progress" && status != "done" && status != "all" {
		fmt.Println("Error: Invalid status filter.")
		return
	}
```
- Validates the filter.

```go
	fmt.Println("\nTasks:")
	count := 0
	for _, task := range taskList.Tasks {
		if status == "all" || task.Status == status {
			fmt.Printf("ID: %d | Title: %s | Status: %s\n", task.ID, task.Title, task.Status)
			count++
		}
	}
	if count == 0 {
		fmt.Println("No tasks found.")
	}
}
```
- Displays tasks that match the filter.
- If no tasks match, informs the user.

---

### **File Operations**
```go
func saveTasksToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(taskList)
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
	}
}
```
- Saves the current tasks to a file (`tasks.json`) using JSON encoding.

```go
func loadTasksFromFile(filename string) error {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		taskList = TaskList{NextID: 1}
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&taskList)
	if err != nil {
		taskList = TaskList{NextID: 1}
		return err
	}
	return nil
}
```
- Loads tasks from the file. If the file doesn't exist, initializes an empty task list.

---

### **Validation Helper**
```go
func isValidStatus(status string) bool {
	return status == "todo" || status == "in-progress" || status == "done"
}
```
- Validates if the given status is valid.
