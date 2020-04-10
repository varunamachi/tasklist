package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/varunamachi/tasklist/srv/todo"
	"github.com/varunamachi/tasklist/srv/util"
)

func main() {
	cmd := ""
	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	} else {
		fmt.Println("Valid commands are: ")
		fmt.Println("\t 'add' - Add a task")
		fmt.Println("\t 'list' - List tasks in the tasks list")
		os.Exit(-1)
	}

	tasklist, err := todo.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to load tasklist %s", err.Error())
	}
	switch cmd {
	case "add":
		addTaskAction(tasklist, os.Args[2:])
	case "list":
		listTaskAction(tasklist, os.Args[2:])
	default:
		log.Fatalf("Invalid command %s", cmd)
	}

}

func addTaskAction(tl *todo.TaskList, args []string) {
	if len(args) < 2 {
		log.Fatalf("'add': Invalid number of arguments provided")
	}
	item := todo.NewTaskItem(args[0], args[1], time.Now().Add(24*time.Hour))
	tl.Add(item)
	todo.Store(tl)
}

func listTaskAction(tl *todo.TaskList, args []string) {
	util.DumpJSON(tl)
}
