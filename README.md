# toGoList

A simple and efficient CLI todo list application built with Go.

## Features

- ✅ Add new tasks
- 📋 List all tasks
- ✔️ Mark tasks as complete
- 🗑️ Delete tasks
- 💾 Persistent storage

## Installation

```bash
# Clone the repository
git clone https://github.com/AdamAzuddin/tgl.git
cd tgl

# Build and install the application
go build && go install

```

## Usage

```bash
# Add a new task
tgl add "Buy groceries"

# List all tasks
tgl list

# Complete a task
tgl tick 1

# Delete a task
tgl rmv 1

# Show help
tgl help
```

## Project Structure

```
ToGoList/
├── internal/            # Private application code
│   ├── todo/           # Todo business logic
│   └── storage/        # Data persistence
├── main.go             # Application entry point
├── go.mod              # Go module file
└── README.md           # Project documentation
```