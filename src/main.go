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

// Task represents a single task in the tracker.
type Task struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Status   string `json:"status"`
    Priority string `json:"priority"` // New priority field
}

// TaskList manages a slice of Task and the next ID.
type TaskList struct {
    Tasks  []Task `json:"tasks"`
    NextID int    `json:"next_id"`
}

var taskList TaskList
var filePath = "tasks.json"

func main() {
    if err := loadTasksFromFile(filePath); err != nil {
        fmt.Println("Warning: Could not load tasks:", err)
    }
    defer saveTasksToFile(filePath)

    scanner := bufio.NewScanner(os.Stdin)
    for {
        displayMenu()
        scanner.Scan()
        choice := strings.TrimSpace(scanner.Text())
        switch choice {
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
            fmt.Println("Exiting... Goodbye!")
            return
        default:
            fmt.Println("Invalid choice. Please select a valid option.")
        }
    }
}

func displayMenu() {
    fmt.Println("\nEnhanced Task Tracker CLI")
    fmt.Println("1. Add Task")
    fmt.Println("2. Update Task Status")
    fmt.Println("3. List Tasks")
    fmt.Println("4. Delete Task")
    fmt.Println("5. Search Tasks")
    fmt.Println("6. Sort Tasks")
    fmt.Println("7. Exit")
    fmt.Print("Choose an option: ")
}

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
    if priority != "High" && priority != "Medium" && priority != "Low" {
        fmt.Println("Error: Invalid priority.")
        return
    }

    newTask := Task{
        ID:       taskList.NextID,
        Title:    title,
        Status:   "todo",
        Priority: priority,
    }
    taskList.Tasks = append(taskList.Tasks, newTask)
    taskList.NextID++
    fmt.Printf("Task '%s' added with ID %d.\n", title, newTask.ID)
}

func updateTaskStatus(scanner *bufio.Scanner) {
    fmt.Print("Enter task ID to update: ")
    scanner.Scan()
    idStr := strings.TrimSpace(scanner.Text())
    id, err := strconv.Atoi(idStr)
    if err != nil || id < 1 {
        fmt.Println("Error: Invalid task ID.")
        return
    }

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

func listTasks(scanner *bufio.Scanner) {
    fmt.Print("Enter status to filter by (todo, in-progress, done, all): ")
    scanner.Scan()
    status := strings.TrimSpace(scanner.Text())
    if status != "todo" && status != "in-progress" && status != "done" && status != "all" {
        fmt.Println("Error: Invalid status filter.")
        return
    }

    fmt.Println("\nTasks:")
    count := 0
    for _, task := range taskList.Tasks {
        if status == "all" || task.Status == status {
            fmt.Printf("ID: %d | Title: %s | Status: %s | Priority: %s\n", task.ID, task.Title, task.Status, task.Priority)
            count++
        }
    }
    if count == 0 {
        fmt.Println("No tasks found.")
    }
}

func deleteTask(scanner *bufio.Scanner) {
    fmt.Print("Enter task ID to delete: ")
    scanner.Scan()
    idStr := strings.TrimSpace(scanner.Text())
    id, err := strconv.Atoi(idStr)
    if err != nil || id < 1 {
        fmt.Println("Error: Invalid task ID.")
        return
    }

    for i, task := range taskList.Tasks {
        if task.ID == id {
            fmt.Printf("Are you sure you want to delete task '%s'? (yes/no): ", task.Title)
            scanner.Scan()
            confirmation := strings.TrimSpace(scanner.Text())
            if confirmation == "yes" {
                taskList.Tasks = append(taskList.Tasks[:i], taskList.Tasks[i+1:]...)
                fmt.Printf("Task ID %d deleted.\n", id)
                return
            } else {
                fmt.Println("Deletion cancelled.")
                return
            }
        }
    }
    fmt.Println("Error: Task not found.")
}

func searchTasks(scanner *bufio.Scanner) {
    fmt.Print("Enter keyword to search for tasks: ")
    scanner.Scan()
    keyword := strings.TrimSpace(scanner.Text())
    if keyword == "" {
        fmt.Println("Error: Keyword cannot be empty.")
        return
    }

    fmt.Println("\nSearch Results:")
    count := 0
    for _, task := range taskList.Tasks {
        if strings.Contains(strings.ToLower(task.Title), strings.ToLower(keyword)) {
            fmt.Printf("ID: %d | Title: %s | Status: %s | Priority: %s\n", task.ID, task.Title, task.Status, task.Priority)
            count++
        }
    }
    if count == 0 {
        fmt.Println("No tasks found matching the keyword.")
    }
}

func sortTasks(by string, ascending bool) {
    switch by {
    case "id":
        sort.Slice(taskList.Tasks, func(i, j int) bool {
            if ascending {
                return taskList.Tasks[i].ID < taskList.Tasks[j].ID
            }
            return taskList.Tasks[i].ID > taskList.Tasks[j].ID
        })
    case "title":
        sort.Slice(taskList.Tasks, func(i, j int) bool {
            if ascending {
                return taskList.Tasks[i].Title < taskList.Tasks[j].Title
            }
            return taskList.Tasks[i].Title > taskList.Tasks[j].Title
        })
    case "status":
        sort.Slice(taskList.Tasks, func(i, j int) bool {
            if ascending {
                return taskList.Tasks[i].Status < taskList.Tasks[j].Status
            }
            return taskList.Tasks[i].Status > taskList.Tasks[j].Status
        })
    case "priority":
        sort.Slice(taskList.Tasks, func(i, j int) bool {
            if ascending {
                return taskList.Tasks[i].Priority < taskList.Tasks[j].Priority
            }
            return taskList.Tasks[i].Priority > taskList.Tasks[j].Priority
        })
    default:
        fmt.Println("Error: Invalid sort option.")
        return
    }
    fmt.Println("Tasks sorted successfully.")
}

func sortTasksMenu(scanner *bufio.Scanner) {
    fmt.Println("Sort tasks by:")
    fmt.Println("1. ID (Ascending)")
    fmt.Println("2. ID (Descending)")
    fmt.Println("3. Title (Ascending)")
    fmt.Println("4. Title (Descending)")
    fmt.Println("5. Status (Ascending)")
    fmt.Println("6. Status (Descending)")
    fmt.Println("7. Priority (Ascending)")
    fmt.Println("8. Priority (Descending)")
    fmt.Print("Choose an option: ")
    scanner.Scan()
    choice := strings.TrimSpace(scanner.Text())

    var by string
    var ascending bool

    switch choice {
    case "1":
        by, ascending = "id", true
    case "2":
        by, ascending = "id", false
    case "3":
        by, ascending = "title", true
    case "4":
        by, ascending = "title", false
    case "5":
        by, ascending = "status", true
    case "6":
        by, ascending = "status", false
    case "7":
        by, ascending = "priority", true
    case "8":
        by, ascending = "priority", false
    default:
        fmt.Println("Invalid choice.")
        return
    }
    sortTasks(by, ascending)
}

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

func isValidStatus(status string) bool {
    return status == "todo" || status == "in-progress" || status == "done"
}