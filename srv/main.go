package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/varunamachi/tasklist/srv/db"
	"github.com/varunamachi/tasklist/srv/todo"
	"github.com/varunamachi/tasklist/srv/util"

	_ "github.com/lib/pq"
)

func main() {
	connect()

	storage := &db.PostgresStorage{}
	err := storage.Init()
	if err != nil {
		log.Fatalf("Failed to initialize data source")
	}

	cmd := ""
	if len(os.Args) >= 2 {
		cmd = os.Args[1]
	} else {
		fmt.Println("Valid commands are: ")
		fmt.Println("\t 'add' - Add a task")
		fmt.Println("\t 'list' - List tasks in the tasks list")
		os.Exit(-1)
	}

	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to load tasklist %s", err.Error())
	}
	switch cmd {
	case "add":
		addTaskAction(storage, os.Args[2:])
	case "list":
		listTaskAction(storage, os.Args[2:])
	case "ping-db":
		pingDbAction()
	default:
		log.Fatalf("Invalid command %s", cmd)
	}

}

func addTaskAction(st todo.Storage, args []string) {
	if len(args) < 2 {
		log.Fatalf("'add': Invalid number of arguments provided")
	}
	item := todo.NewTaskItem(args[0], args[1], time.Now().Add(24*time.Hour))
	err := st.Add(item)
	if err != nil {
		log.Fatalf("Failed to add %s", err.Error())
	}
	log.Print("Added...")
}

func listTaskAction(st todo.Storage, args []string) {
	items, err := st.RetrieveAll(0, 10)
	if err != nil {
		log.Fatalf("Failed to retrieve todo items: %s", err.Error())
	}
	util.DumpJSON(items)
}

func pingDbAction() {
	err := db.Conn().Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %s", err.Error())
	}

	log.Println("Database Connection Works!!")
}

func connect() error {

	// prompt := fmt.Sprintf("Password for %s@%s", user, host)
	// passwd, err := util.AskSecret(prompt)
	// if err != nil {
	// 	log.Fatalf("Failed to read password from terminal: %s", err.Error())
	// }

	err := db.Connect(&db.ConnOpts{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "tasklist",
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	return err
}
