<div align="center">

# ✅ Task Tracker CLI — Go Edition

**A blazing-fast, minimalist task manager that lives right in your terminal.**

[![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Stars](https://img.shields.io/github/stars/MisaghMomeniB/Task-Tracker-CLI-Go?style=for-the-badge&color=yellow)](https://github.com/MisaghMomeniB/Task-Tracker-CLI-Go/stargazers)
[![Forks](https://img.shields.io/github/forks/MisaghMomeniB/Task-Tracker-CLI-Go?style=for-the-badge&color=blue)](https://github.com/MisaghMomeniB/Task-Tracker-CLI-Go/network/members)

> *"Stop juggling tasks in your head. Let the terminal handle it."*

</div>

---

## 🧭 Table of Contents

- [About the Project](#-about-the-project)
- [Features](#-features)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [How to Use](#-how-to-use)
- [Data Persistence](#-data-persistence)
- [Roadmap](#-roadmap)
- [Contributing](#-contributing)
- [Author](#-author)

---

## 🗂 About the Project

**Task Tracker CLI** is a lightweight, no-dependency command-line application written entirely in **Go**. It gives you everything you need to manage your daily tasks without ever leaving the terminal — no browser, no app, no bloat.

Built for developers, power users, and productivity nerds who prefer keyboard over clicks. Tasks are stored locally as JSON, so they're yours — no cloud, no account, no nonsense.

```
$ go run main.go

  ╔══════════════════════════╗
  ║    📋 Task Tracker CLI   ║
  ╠══════════════════════════╣
  ║  1. Add Task             ║
  ║  2. Update Task Status   ║
  ║  3. List Tasks           ║
  ║  4. Delete Task          ║
  ║  5. Search Tasks         ║
  ║  6. Sort Tasks           ║
  ║  7. Exit                 ║
  ╚══════════════════════════╝
  Choose an option:
```

---

## ✨ Features

| Feature | Description |
|---|---|
| ➕ **Add Tasks** | Create tasks with a title and priority level (High / Medium / Low) |
| 🔁 **Update Status** | Move tasks across `todo` → `in-progress` → `done` |
| 📋 **List & Filter** | View all tasks or filter by status in one command |
| 🔎 **Search** | Instantly find tasks by keyword |
| 🔀 **Sort** | Order your list by ID, title, status, or priority |
| 🗑️ **Delete** | Remove tasks safely with a confirmation prompt |
| 💾 **Auto-Save** | Everything persists to `tasks.json` automatically |
| ⚡ **Zero Dependencies** | Pure Go standard library — no installs, no fuss |

---

## 📁 Project Structure

```
Task-Tracker-CLI-Go/
│
├── src/
│   ├── main.go        # 🎛️  CLI menu, user prompts & action routing
│   └── task.go        # 🧱  Task model, JSON serialization & file I/O
│
└── README.md          # 📖  You are here
```

**`main.go`** — The brain of the app. Handles the interactive menu loop, reads user input, and dispatches to the right action.

**`task.go`** — Defines the `Task` struct with fields like `ID`, `Title`, `Status`, and `Priority`. Manages loading and saving to `tasks.json`.

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.18+](https://golang.org/dl/) installed on your machine

### Installation

```bash
# 1. Clone the repository
git clone https://github.com/MisaghMomeniB/Task-Tracker-CLI-Go.git

# 2. Navigate into the source directory
cd Task-Tracker-CLI-Go/src

# 3a. Run directly with Go
go run main.go

# 3b. Or build a binary
go build -o task-tracker
./task-tracker
```

---

## 🖥️ How to Use

### ➕ Adding a Task

```
Choose an option: 1
Enter task title: Write unit tests
Select priority (High/Medium/Low): High
✅ Task 'Write unit tests' added (ID: 1)
```

### 📋 Listing Tasks

```
Choose an option: 3
Status? (todo/in-progress/done/all): all

ID:1 | Title: Write unit tests | Status: todo | Priority: High
ID:2 | Title: Update docs      | Status: done | Priority: Low
```

### 🔁 Updating a Task

```
Choose an option: 2
Enter task ID: 1
Select new status (todo/in-progress/done): in-progress
🔄 Task ID 1 status updated to 'in-progress'
```

### 🔎 Searching

```
Choose an option: 5
Enter keyword: unit
🔍 Found: ID:1 | Title: Write unit tests | Status: in-progress | Priority: High
```

### 🔀 Sorting

```
Choose an option: 6
Sort by (id/title/status/priority): priority
```

### 🗑️ Deleting a Task

```
Choose an option: 4
Enter task ID: 2
Are you sure you want to delete task ID 2? (y/n): y
🗑️ Task ID 2 deleted.
```

---

## 💾 Data Persistence

All tasks are automatically saved to a `tasks.json` file in the same directory as the executable. The file is created on first run and updated after every action — no manual saving required.

```json
[
  {
    "id": 1,
    "title": "Write unit tests",
    "status": "in-progress",
    "priority": "High"
  }
]
```

---

## 🛣️ Roadmap

- [x] Core CRUD operations
- [x] Priority levels
- [x] Status filtering & sorting
- [x] Keyword search
- [x] JSON persistence
- [ ] Due dates & deadline warnings ⏰
- [ ] Project/tag grouping 🏷️
- [ ] CLI flags mode (non-interactive) 🚩
- [ ] Batch import/export (CSV/JSON) 📦
- [ ] Automated test suite 🧪
- [ ] Colorized terminal output 🎨

---

## 🤝 Contributing

Contributions make open source great! Here's how to get involved:

1. **Fork** the repository
2. **Create** your feature branch: `git checkout -b feature/your-feature`
3. **Commit** your changes: `git commit -m "Add your feature"`
4. **Push** to the branch: `git push origin feature/your-feature`
5. **Open** a Pull Request — and describe what you built!

Please keep code clean, commented, and consistent with the existing style.

---

## 📄 License

Distributed under the **MIT License**. See [`LICENSE`](LICENSE) for more information.

---

<div align="center">

## 👨‍💻 Author

**Misagh Momeni Bashusqeh**
*Software Developer*

[![GitHub](https://img.shields.io/badge/GitHub-MisaghMomeniB-181717?style=for-the-badge&logo=github)](https://github.com/MisaghMomeniB)

---

*Built with ❤️ and Go — because great tools deserve a great terminal.*

</div>
