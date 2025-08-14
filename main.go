package main

import (
	"fmt"
	"os"

	"github.com/AdamAzuddin/tgl/internal/todo"
)

func main() {
	fmt.Println("Welcome to ToGoList")

	todo.List()

	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "tick":
		handleTick()
	case "rmv":
		handleRmv()
	case "help", "-h", "--help":
		showHelp()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}

}

func handleAdd() {
	fmt.Println("Handle Add")
}

func handleList() {
	fmt.Println("Handle List")
}

func handleTick() {
	fmt.Println("Handle Tick")
}

func handleRmv() {
	fmt.Println("Handle Remove")
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
