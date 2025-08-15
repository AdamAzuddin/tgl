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

func (t Task) ToString() string{
	mark := "x"
	if t.Completed{
		mark="-"
	}
	return fmt.Sprintf("%d %s %s\n",t.ID, mark, t.Title)
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