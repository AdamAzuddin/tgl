package todo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ToDoList struct {
	Tasks  []*Task
	NextID int
}

func NewToDoList() *ToDoList {
	return &ToDoList{
		Tasks:  make([]*Task, 0),
		NextID: 1,
	}
}

func AddTask(t *ToDoList, message string, file *os.File) (*Task, error) {
	task := NewTask(t.NextID, message)
	t.Tasks = append(t.Tasks, task)
	if _, err := file.Seek(0, io.SeekEnd); err != nil {
		return nil, err
	}
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(task.ToString()); err != nil {
		return nil, err
	}
	if err := writer.Flush(); err != nil {
		return nil, err
	}
	t.NextID++
	return task, nil
}

func RemoveTask(t *ToDoList, id int, file *os.File) error {
	found := false
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Task with id %d not found", id)
	}

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	removedTask := ""
	lines := []string{}
	i := 1
	for scanner.Scan() {
		if i != id {
			lines = append(lines, scanner.Text())
		} else {
			removedTask = scanner.Text()
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}
	file.Seek(0, 0)

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	fmt.Printf("Removed task: %s\n", removedTask)
	return writer.Flush()
}

func TickTask(t *ToDoList, id int, file *os.File) error {
	for _, task := range t.Tasks {
		if task.ID == id {
			task.Completed = true
		}
	}

	// update the file
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// split by a space. [0] = id, [1] = completed
		parts := strings.SplitN(line, " ", 3)
		idInLine, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		if idInLine != id {
			lines = append(lines, strings.Join(parts, " "))
		} else {
			// change the - to x
			newLine := fmt.Sprintf("%d - %s", idInLine, parts[2])
			lines = append(lines, newLine)
		}
	}
	if err:= file.Truncate(0); err != nil{
		return err
	}
	file.Seek(0,0)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line+"\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}

func UnTickTask(t *ToDoList, id int, file *os.File) error {
	for _, task := range t.Tasks {
		if task.ID == id {
			task.Completed = true
		}
	}

	// update the file
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// split by a space. [0] = id, [1] = completed
		parts := strings.SplitN(line, " ", 3)
		idInLine, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		if idInLine != id {
			lines = append(lines, strings.Join(parts, " "))
		} else {
			// change the - to x
			newLine := fmt.Sprintf("%d x %s", idInLine, parts[2])
			lines = append(lines, newLine)
		}
	}
	if err:= file.Truncate(0); err != nil{
		return err
	}
	file.Seek(0,0)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line+"\n"); err != nil {
			return err
		}
	}
	return writer.Flush()
}
