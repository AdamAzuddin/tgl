package todo

import (
	"fmt"
	"time"
)

// create new task type

type Task struct{
	ID int
	Title string
	Completed bool
	CreatedAt time.Time
}

func NewTask(id int, title string) *Task{
	return &Task{
		ID: id,
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}

func List(){
	fmt.Println("List")
}