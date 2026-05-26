package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////
// Constants
//////////////////////////////////////////////////////

// Status values
const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

// Priority values
const (
	PriorityHigh   = "High"
	PriorityMedium = "Medium"
	PriorityLow    = "Low"
)

// Priority rank for sorting
var priorityRank = map[string]int{
	PriorityHigh:   3,
	PriorityMedium: 2,
	PriorityLow:    1,
}

// Valid sets for O(1) validation
var validStatuses = map[string]bool{
	StatusTodo: true, StatusInProgress: true, StatusDone: true,
}
var validPriorities = map[string]bool{
	PriorityHigh: true, PriorityMedium: true, PriorityLow: true,
}
var validFilters = map[string]bool{
	StatusTodo: true, StatusInProgress: true, StatusDone: true, "all": true,
}

//////////////////////////////////////////////////////
// Data Models
//////////////////////////////////////////////////////

// Task represents a single task entity.
type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
}

// TaskStore holds all tasks and manages ID generation.
// Tasks are stored as a map for O(1) lookup/update/delete.
type TaskStore struct {
	Tasks  map[int]*Task `json:"tasks"`
	NextID int           `json:"next_id"`
}

// NewTaskStore creates an initialised, empty store.
func NewTaskStore() *TaskStore {
	return &TaskStore{Tasks: make(map[int]*Task), NextID: 1}
}

//////////////////////////////////////////////////////
// Business Logic (pure — no I/O)
//////////////////////////////////////////////////////

// AddTask creates and stores a new task. Returns the created task or an error.
func (s *TaskStore) AddTask(title, priority string) (*Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, errors.New("task title cannot be empty")
	}
	priority = normalisePriority(priority)
	if !validPriorities[priority] {
		return nil, fmt.Errorf("invalid priority %q — use High, Medium, or Low", priority)
	}
	t := &Task{
		ID:       s.NextID,
		Title:    title,
		Status:   StatusTodo,
		Priority: priority,
	}
	s.Tasks[t.ID] = t
	s.NextID++
	return t, nil
}

// UpdateStatus changes the status of a task by ID.
func (s *TaskStore) UpdateStatus(id int, status string) error {
	if !validStatuses[status] {
		return fmt.Errorf("invalid status %q — use todo, in-progress, or done", status)
	}
	t, ok := s.Tasks[id]
	if !ok {
		return fmt.Errorf("task %d not found", id)
	}
	t.Status = status
	return nil
}

// DeleteTask removes a task by ID.
func (s *TaskStore) DeleteTask(id int) (*Task, error) {
	t, ok := s.Tasks[id]
	if !ok {
		return nil, fmt.Errorf("task %d not found", id)
	}
	delete(s.Tasks, id)
	return t, nil
}

// Filter returns tasks matching the given status filter (or all tasks).
func (s *TaskStore) Filter(filter string) []*Task {
	var result []*Task
	for _, t := range s.Tasks {
		if filter == "all" || t.Status == filter {
			result = append(result, t)
		}
	}
	return result
}

// Search returns tasks whose titles contain the keyword (case-insensitive).
func (s *TaskStore) Search(keyword string) []*Task {
	kw := strings.ToLower(keyword)
	var result []*Task
	for _, t := range s.Tasks {
		if strings.Contains(strings.ToLower(t.Title), kw) {
			result = append(result, t)
		}
	}
	return result
}

// SortedTasks returns a stable-sorted copy of the task slice.
func (s *TaskStore) SortedTasks(by string, asc bool) []*Task {
	all := s.Filter("all")
	sort.SliceStable(all, func(i, j int) bool {
		var less bool
		switch by {
		case "id":
			less = all[i].ID < all[j].ID
		case "title":
			less = strings.ToLower(all[i].Title) < strings.ToLower(all[j].Title)
		case "status":
			less = all[i].Status < all[j].Status
		case "priority":
			less = priorityRank[all[i].Priority] < priorityRank[all[j].Priority]
		default:
			less = all[i].ID < all[j].ID
		}
		if asc {
			return less
		}
		return !less
	})
	return all
}

//////////////////////////////////////////////////////
// Persistence
//////////////////////////////////////////////////////

// Save atomically writes the store to disk (write-then-rename).
func (s *TaskStore) Save(path string) error {
	tmp := path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return fmt.Errorf("save: create tmp: %w", err)
	}
	w := bufio.NewWriter(f)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		f.Close()
		os.Remove(tmp)
		return fmt.Errorf("save: encode: %w", err)
	}
	if err := w.Flush(); err != nil {
		f.Close()
		os.Remove(tmp)
		return fmt.Errorf("save: flush: %w", err)
	}
	f.Close()
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return fmt.Errorf("save: rename: %w", err)
	}
	return nil
}

// Load reads the store from disk. Returns a fresh store if the file doesn't exist.
func Load(path string) (*TaskStore, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return NewTaskStore(), nil
	}
	if err != nil {
		return nil, fmt.Errorf("load: open: %w", err)
	}
	defer f.Close()

	store := &TaskStore{}
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(store); err != nil {
		return nil, fmt.Errorf("load: decode: %w", err)
	}
	if store.Tasks == nil {
		store.Tasks = make(map[int]*Task)
	}
	return store, nil
}

//////////////////////////////////////////////////////
// CLI Application
//////////////////////////////////////////////////////

const filePath = "tasks.json"

func main() {
	store, err := Load(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not load tasks: %v\n", err)
		store = NewTaskStore()
	}

	defer func() {
		if err := store.Save(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		printMenu()
		scanner.Scan()
		switch strings.TrimSpace(scanner.Text()) {
		case "1":
			cmdAdd(scanner, store)
		case "2":
			cmdUpdateStatus(scanner, store)
		case "3":
			cmdList(scanner, store)
		case "4":
			cmdDelete(scanner, store)
		case "5":
			cmdSearch(scanner, store)
		case "6":
			cmdSort(scanner, store)
		case "7":
			fmt.Println("Exiting... Goodbye 👋")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

//////////////////////////////////////////////////////
// Command Handlers (I/O layer)
//////////////////////////////////////////////////////

func cmdAdd(sc *bufio.Scanner, s *TaskStore) {
	title := prompt(sc, "Enter task title: ")
	priority := prompt(sc, "Enter priority (High, Medium, Low): ")

	t, err := s.AddTask(title, priority)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("✅ Task added (ID: %d)\n", t.ID)
}

func cmdUpdateStatus(sc *bufio.Scanner, s *TaskStore) {
	id, ok := readID(sc)
	if !ok {
		return
	}
	status := prompt(sc, "Enter new status (todo, in-progress, done): ")
	if err := s.UpdateStatus(id, status); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("✅ Status updated.")
}

func cmdList(sc *bufio.Scanner, s *TaskStore) {
	filter := prompt(sc, "Filter by status (todo, in-progress, done, all): ")
	if !validFilters[filter] {
		fmt.Println("Error: invalid filter.")
		return
	}
	tasks := s.Filter(filter)
	sortByID(tasks)
	printTasks(tasks)
}

func cmdDelete(sc *bufio.Scanner, s *TaskStore) {
	id, ok := readID(sc)
	if !ok {
		return
	}
	t, err := s.DeleteTask(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	confirm := prompt(sc, fmt.Sprintf("Delete task %q? (yes/no): ", t.Title))
	if confirm != "yes" {
		// Put it back — user cancelled.
		s.Tasks[t.ID] = t
		fmt.Println("Deletion cancelled.")
		return
	}
	fmt.Println("🗑️  Task deleted.")
}

func cmdSearch(sc *bufio.Scanner, s *TaskStore) {
	kw := prompt(sc, "Enter keyword: ")
	if kw == "" {
		fmt.Println("Error: keyword cannot be empty.")
		return
	}
	tasks := s.Search(kw)
	sortByID(tasks)
	fmt.Println("\n🔍 Search results:")
	printTasks(tasks)
}

func cmdSort(sc *bufio.Scanner, s *TaskStore) {
	type option struct{ by string; asc bool }
	opts := map[string]option{
		"1": {"id", true}, "2": {"id", false},
		"3": {"title", true}, "4": {"title", false},
		"5": {"status", true}, "6": {"status", false},
		"7": {"priority", true}, "8": {"priority", false},
	}
	fmt.Println("Sort by:")
	fmt.Println("  1 ID ↑   2 ID ↓")
	fmt.Println("  3 Title ↑   4 Title ↓")
	fmt.Println("  5 Status ↑   6 Status ↓")
	fmt.Println("  7 Priority ↑   8 Priority ↓")
	choice := prompt(sc, "Choose: ")
	opt, ok := opts[choice]
	if !ok {
		fmt.Println("Invalid sort option.")
		return
	}
	tasks := s.SortedTasks(opt.by, opt.asc)
	fmt.Printf("\n✅ Sorted by %s (%s):\n", opt.by, map[bool]string{true: "asc", false: "desc"}[opt.asc])
	printTasks(tasks)
}

//////////////////////////////////////////////////////
// Helpers
//////////////////////////////////////////////////////

func printMenu() {
	fmt.Println("\n🚀 Task Tracker")
	fmt.Println("1. Add task")
	fmt.Println("2. Update status")
	fmt.Println("3. List tasks")
	fmt.Println("4. Delete task")
	fmt.Println("5. Search tasks")
	fmt.Println("6. Sort tasks")
	fmt.Println("7. Exit")
	fmt.Print("Choose: ")
}

func prompt(sc *bufio.Scanner, msg string) string {
	fmt.Print(msg)
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

func readID(sc *bufio.Scanner) (int, bool) {
	raw := prompt(sc, "Enter task ID: ")
	id, err := strconv.Atoi(raw)
	if err != nil || id < 1 {
		fmt.Println("Error: invalid task ID.")
		return 0, false
	}
	return id, true
}

func printTask(t *Task) {
	fmt.Printf("  [%d] %-40s  %-12s  %s\n", t.ID, t.Title, t.Status, t.Priority)
}

func printTasks(tasks []*Task) {
	if len(tasks) == 0 {
		fmt.Println("  No tasks found.")
		return
	}
	fmt.Printf("  %-4s %-40s  %-12s  %s\n", "ID", "Title", "Status", "Priority")
	fmt.Println("  " + strings.Repeat("─", 66))
	for _, t := range tasks {
		printTask(t)
	}
}

func sortByID(tasks []*Task) {
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})
}

// normalisePriority converts "high" / "HIGH" → "High" etc.
func normalisePriority(p string) string {
	p = strings.ToLower(strings.TrimSpace(p))
	switch p {
	case "high":
		return PriorityHigh
	case "medium":
		return PriorityMedium
	case "low":
		return PriorityLow
	}
	return p // return as-is so validation catches it cleanly
}