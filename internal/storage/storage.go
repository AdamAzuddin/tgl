package storage

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/AdamAzuddin/tgl/internal/todo"
)

func LoadTasksFromFile(file *os.File) (*todo.ToDoList, error){
	file.Seek(0,0)
	tdl := todo.NewToDoList()
	// read from file
	scanner := bufio.NewScanner(file)
	lastId := 0
	for scanner.Scan(){
		line := scanner.Text()
		// split by a space. [0] = id, [1] = completed
		id, err := strconv.Atoi(strings.Split(line, "")[0])
		if err != nil{
			return nil,err
		}
		completed := strings.Split(line, "")[1]
		completedBool := completed == "-"
		title := strings.Join(strings.Split(line, " ")[:2], "")
		t := todo.Task{
			Title: title,
			ID: id,
			Completed: completedBool,
		}
		tdl.Tasks = append(tdl.Tasks, &t)
		lastId = id
	}
	tdl.NextID = lastId+1
	return tdl,nil
}