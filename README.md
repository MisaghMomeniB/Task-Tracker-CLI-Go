### Enhanced Task Tracker CLI 🚀  
**A feature-rich command-line tool for managing your tasks efficiently!**  
With support for priorities, searching, sorting, and much more, this CLI tool helps you organize your tasks like a pro.  

---

### Features ✨  
1. **Add Tasks** 📝  
   Easily create new tasks with a title and priority (`High`, `Medium`, or `Low`).  
   
2. **Update Task Status** 🔄  
   Change task statuses between `todo`, `in-progress`, and `done` effortlessly.  

3. **List Tasks** 📋  
   View tasks filtered by status or see them all in one go, with clear details about priority.  

4. **Delete Tasks** 🗑️  
   Safely remove tasks with confirmation prompts to avoid accidental deletions.  

5. **Search Tasks** 🔍  
   Quickly locate tasks by keywords in their titles.  

6. **Sort Tasks** ⬆️⬇️  
   Organize tasks by:
   - **ID**
   - **Title**
   - **Status**
   - **Priority**  

7. **Persistent Storage** 💾  
   Save tasks to a file (`tasks.json`) and reload them on startup.  

8. **User-Friendly Menu** 🎯  
   Intuitive, easy-to-navigate menu for managing tasks.  

---

### Installation 🛠️  

1. **Clone the repository**:  
   ```bash
   git clone https://github.com/your-username/enhanced-task-tracker.git
   cd enhanced-task-tracker
   ```

2. **Run the program**:  
   ```bash
   go run enhanced_task_tracker.go
   ```

3. **(Optional)** Compile for standalone use:  
   ```bash
   go build -o task-tracker
   ./task-tracker
   ```

---

### Usage 🧑‍💻  

Upon running the program, you'll see a menu:  

```text
Enhanced Task Tracker CLI
1. Add Task
2. Update Task Status
3. List Tasks
4. Delete Task
5. Search Tasks
6. Sort Tasks
7. Exit
Choose an option:
```

Simply select an option and follow the prompts to manage your tasks.  

---

### Examples 🌟  

#### Adding a Task  
```text
Enter task title: Learn Go
Enter priority (High, Medium, Low): High
Task 'Learn Go' added with ID 1.
```

#### Listing All Tasks  
```text
Enter status to filter by (todo, in-progress, done, all): all

Tasks:
ID: 1 | Title: Learn Go | Status: todo | Priority: High
```

#### Updating a Task's Status  
```text
Enter task ID to update: 1
Enter new status (todo, in-progress, done): in-progress
Task ID 1 updated to status 'in-progress'.
```

#### Searching Tasks  
```text
Enter keyword to search for tasks: learn

Search Results:
ID: 1 | Title: Learn Go | Status: in-progress | Priority: High
```

#### Sorting Tasks  
```text
Sort tasks by:
1. ID
2. Title
3. Status
4. Priority
Choose an option: 4
Tasks sorted successfully.
```

---

### Contributing 🤝  

Want to improve the Enhanced Task Tracker? Contributions are welcome!  

1. Fork the repository.  
2. Create a feature branch.  
3. Commit your changes.  
4. Submit a pull request.  

---

### License 📜  
This project is licensed under the MIT License.  

Enjoy productive task tracking! 🎉
