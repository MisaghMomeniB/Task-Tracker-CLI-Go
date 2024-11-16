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