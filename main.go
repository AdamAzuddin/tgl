package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/AdamAzuddin/tgl/internal/storage"
	"github.com/AdamAzuddin/tgl/internal/todo"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}
	command := os.Args[1]

	file, err := os.OpenFile("tgl.txt", os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	
	toDoList, err := storage.LoadTasksFromFile(file)
	if err != nil{
		fmt.Printf("Error: %s\n", err)
		return
	}

	switch command {
	case "add":
		msg := os.Args[2]
		handleAdd(toDoList, msg, file)
	case "list":
		handleList(file)
	case "tick":
		handleTick(toDoList, file)
	case "untick":
		handleUnTick(toDoList, file)
	case "rmv":
		handleRmv(toDoList, file)
	case "clear":
		handleClear(file)
	case "help", "-h", "--help":
		showHelp()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}

}

func handleAdd(toDoList *todo.ToDoList, title string, file *os.File) {
	task, err := todo.AddTask(toDoList, title, file)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf("%s added as a task with id %d\n", title, task.ID)
}

func handleList(file *os.File) {
	file.Seek(0,0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}

func handleTick(t *todo.ToDoList, file *os.File) {
	fmt.Println("Handle Tick")
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid id of %s\n", os.Args[2])
		return
	}

	err = todo.TickTask(t, int(id), file)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Task with id %s ticked successfully\n", os.Args[2])
}

func handleUnTick(t *todo.ToDoList, file *os.File) {
	fmt.Println("Handle Tick")
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid id of %s\n", os.Args[2])
		return
	}

	err = todo.UnTickTask(t, int(id), file)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Task with id %s unticked successfully\n", os.Args[2])
}

func handleRmv(toDoList *todo.ToDoList, file *os.File) {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid id of %s\n", os.Args[2])
		return
	}

	err = todo.RemoveTask(toDoList, int(id), file)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

func handleClear(file *os.File) error{
	if err:= file.Truncate(0); err != nil{
		return err
	}
	if _, err := file.Seek(0,0); err != nil{
		return err
	}
	return nil
}

func showHelp() {
	fmt.Println("tgl - A simple todo list CLI")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  tgl add \"task description\"  - Add a new task")
	fmt.Println("  tgl list                    - List all tasks")
	fmt.Println("  tgl tick [id]               - Mark task as completed")
	fmt.Println("  tgl untick [id]               - Mark task as uncompleted")
	fmt.Println("  tgl rmv [id]                - Remove a task")
	fmt.Println("  tgl clear                - Remove all tasks")
	fmt.Println("  tgl help                    - Show this help")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  tgl add \"Buy groceries\"")
	fmt.Println("  tgl list")
	fmt.Println("  tgl tick 1")
	fmt.Println("  tgl rmv 2")
}
