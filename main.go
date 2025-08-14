package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AdamAzuddin/tgl/internal/todo"
)

func main() {
	toDoList := todo.NewToDoList()

	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		msg := os.Args[2]
		handleAdd(toDoList, msg)
	case "list":
		handleList(toDoList)
	case "tick":
		handleTick(toDoList)
	case "rmv":
		handleRmv(toDoList)
	case "help", "-h", "--help":
		showHelp()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}

}

func handleAdd(toDoList *todo.ToDoList, title string) {
	task := todo.AddTask(toDoList, title)
	fmt.Printf("%s added as a task with id %d\n", title, task.ID)
}

func handleList(t *todo.ToDoList) {
	fmt.Println("Handle List")
	for _, task := range t.Tasks{
		fmt.Printf("%s		%d\n", task.Title, task.ID)
	}
}

func handleTick(t *todo.ToDoList) {
	fmt.Println("Handle Tick")
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid id of %s\n", os.Args[2])
		return
	}

	err = todo.TickTask(t, int(id))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Task with id %s removed successfully\n", os.Args[2])
}

func handleRmv(toDoList *todo.ToDoList) {
	fmt.Println("Handle Remove")
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid id of %s\n", os.Args[2])
		return
	}

	err = todo.RemoveTask(toDoList, int(id))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Task with id %s removed successfully\n", os.Args[2])
}

func showHelp() {
	fmt.Println("tgl - A simple todo list CLI")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  tgl add \"task description\"  - Add a new task")
	fmt.Println("  tgl list                    - List all tasks")
	fmt.Println("  tgl tick [id]               - Mark task as completed")
	fmt.Println("  tgl rmv [id]                - Remove a task")
	fmt.Println("  tgl help                    - Show this help")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  tgl add \"Buy groceries\"")
	fmt.Println("  tgl list")
	fmt.Println("  tgl tick 1")
	fmt.Println("  tgl rmv 2")
}
