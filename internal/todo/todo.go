package todo

type ToDoList struct{
	tasks []*Task
	nextID int
}


func NewToDoList() *ToDoList{
	return &ToDoList{
		tasks: make([]*Task, 0),
		nextID: 1,
	}
}