package tasks

import (
	"log"
	"runtime"
	"strconv"
)

type ToDo struct {
	Title, Status string
}

type EditTodo struct {
	Title, NewTitle, NewStatus string
}

type Task int

var todoSlice []ToDo

func (t *Task) GetToDo(title string, reply *ToDo) error {
	var found ToDo

	for _, v := range todoSlice {
		if v.Title == title {
			found = v
		}
	}

	*reply = found
	return nil
}

func (t *Task) GetSlice(title string, reply *[]ToDo) error {
	*reply = todoSlice
	log.Printf("GetSlice : ")
	log.Println(todoSlice)
	return nil
}

func (t *Task) MakeToDo(todo ToDo, reply *ToDo) error {
	if todo.Title == "Mem usage" {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		todo.Status = strconv.Itoa(int(m.Sys/1024/1024)) + " MiB"
	}
	todoSlice = append(todoSlice, todo)
	*reply = todo
	log.Printf("MakeToDo : ")
	log.Println(todo)
	return nil
}

func (t *Task) EditToDo(todo EditTodo, reply *ToDo) error {
	var edited ToDo

	for i, v := range todoSlice {
		if v.Title == todo.Title {
			todoSlice[i] = ToDo{todo.NewTitle, todo.NewStatus}
			edited = ToDo{todo.NewTitle, todo.NewStatus}
		}
	}

	*reply = edited
	log.Printf("EditToDo : ")
	log.Println(edited)
	return nil
}

// DeleteToDo takes a ToDo type and deletes it from todoArray
func (t *Task) DeleteToDo(todo ToDo, reply *ToDo) error {
	var deleted ToDo
	for i, v := range todoSlice {
		if v.Title == todo.Title && v.Status == todo.Status {
			todoSlice = append(todoSlice[:i], todoSlice[i+1:]...)
			deleted = todo
			break
		}
	}
	*reply = deleted
	log.Printf("DeleteToDo : ")
	log.Println(deleted)
	return nil
}
