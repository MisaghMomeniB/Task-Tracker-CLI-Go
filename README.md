# ✅ Task Tracker CLI (Go)

A **lightweight and feature-rich command-line task manager** built in Go. Manage your tasks—create, update, delete, search, sort, and track progress—all from the terminal with ease.

---

## 📋 Table of Contents

1. [Overview](#overview)  
2. [Features](#features)  
3. [Tech Stack & Requirements](#tech-stack--requirements)  
4. [Installation & Build](#installation--build)  
5. [Usage Examples](#usage-examples)  
6. [Code Structure](#code-structure)  
7. [Future Improvements](#future-improvements)  
8. [Contributing](#contributing)  
9. [License](#license)

---

## 💡 Overview

This CLI tool allows you to manage tasks efficiently through a text-based menu and JSON-backed persistence. It supports task priorities, status tracking, searching, and sorting—ideal for boosting productivity in your terminal workflow. :contentReference[oaicite:1]{index=1}

---

## ✅ Features

- 📝 **Add Tasks**: Create tasks with Title and Priority (High, Medium, Low)  
- 🔁 **Update Status**: Change between `todo`, `in-progress`, and `done`  
- 📄 **List Tasks**: Filter by status or view all  
- 🔎 **Search Tasks**: Quick lookup by keyword  
- 🔄 **Sort Tasks**: By ID, title, status, or priority  
- 🗑️ **Delete Tasks**: Remove entries with confirmation prompt  
- 💾 **Persistent Storage**: Saves to `tasks.json` for data continuity  
- 🧩 **Intuitive CLI Menu**: Navigate options with clear prompts :contentReference[oaicite:2]{index=2}

---

## 🛠️ Tech Stack & Requirements

- **Go 1.18+** (module support required)  
- Standard Go libraries (`encoding/json`, `os`, etc.)  
- No external dependencies

---

## ⚙️ Installation & Build

Clone and compile the project:

```bash
git clone https://github.com/MisaghMomeniB/Task-Tracker-CLI-Go.git
cd Task-Tracker-CLI-Go/src
go build -o task-tracker
````

Or run directly:

```bash
go run main.go
```

---

## 🚀 Usage Examples

### Start the Task Manager

```
$ ./task-tracker
Task Tracker CLI
1. Add Task
2. Update Task Status
3. List Tasks
4. Delete Task
5. Search Tasks
6. Sort Tasks
7. Exit
Choose an option:
```

### Add a Task

```
Enter task title: Fix README typos
Select priority (High/Medium/Low): High
✅ Task 'Fix README typos' added (ID: 1)
```

### List All Tasks

```
Status? (todo/in-progress/done/all): all
ID:1 | Title: Fix README typos | Status: todo | Priority: High
```

### Update Task

```
Enter task ID: 1
Select new status (todo/in-progress/done): in-progress
🔄 Task ID 1 status updated to in-progress
```

### Search & Sort Tasks

Supports keyword search:

```
Enter keyword: README
```

Supports sorting by ID, Title, Status, Priority.

---

## 📁 Code Structure

```
Task-Tracker-CLI-Go/
├── src/
│   ├── main.go          # CLI menu & command handling
│   └── task.go          # Task struct + JSON (de)serialization
└── README.md            # This file
```

* **main.go**: menu navigation, user prompts, action logic
* **task.go**: defines Task model and JSON load/save functions

---

## 🔧 Future Improvements

* Add **due dates** and **reminders**
* Include **project labels** or tags
* Support **batch import/export**
* Implement **CLI flags** vs interactive mode
* Add **automated tests** and performance benchmarks

---

## 🤝 Contributing

Contributions are welcome! To improve:

1. Fork the repo
2. Create a branch (`feature/...`)
3. Code enhancements with proper comments
4. Submit a detailed Pull Request

---

## 📄 License

This project is licensed under the **MIT License** — see `LICENSE` for details.
