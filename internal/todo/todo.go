package todo

import "fmt"

type ToDoList struct {
	Tasks  []*Task
	nextID int
}

func NewToDoList() *ToDoList {
	return &ToDoList{
		Tasks:  make([]*Task, 0),
		nextID: 1,
	}
}

func AddTask(t *ToDoList, message string) *Task{
	task := NewTask(t.nextID, message)
	t.Tasks = append(t.Tasks, task)
	t.nextID++
	return task
}

func RemoveTask(t *ToDoList, id int) error{
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task with id %d not found", id)
}

func TickTask(t *ToDoList, id int) error{
	for _, task := range t.Tasks {
		if task.ID == id {
			task.Completed = true
			return nil
		}
	}
	return fmt.Errorf("Task with id %d not found", id)
}