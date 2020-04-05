package main

import (
	"time"

	"github.com/varunamachi/tasklist/srv/todo"
	"github.com/varunamachi/tasklist/srv/util"
)

func main() {
	tasksList := todo.NewTaskList()
	task1 := todo.NewTaskItem("task1", "blah blah",
		time.Now().Add(24*time.Hour))
	task2 := todo.NewTaskItem("task2", "blah blah blah",
		time.Now().Add(24*time.Hour))

	tasksList.Add(task1, task2)

	util.DumpJSON(tasksList)
}