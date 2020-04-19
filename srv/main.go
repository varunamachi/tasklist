package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/varunamachi/tasklist/srv/api"
	"github.com/varunamachi/tasklist/srv/db"
	"github.com/varunamachi/tasklist/srv/todo"

	_ "github.com/lib/pq"
)

func main() {
	app := createCliApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("App failed! - %s", err.Error())
	}
}

func createCliApp() *cli.App {
	return &cli.App{
		Name: "tasklist",
		Before: func(ctx *cli.Context) error {
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

			storage := &db.PostgresStorage{}
			err = storage.Init()
			if err != nil {
				log.Fatalf("Failed to initialize data source")
			}
			todo.SetStorage(storage)

			return err
		},
		Commands: []cli.Command{
			cli.Command{
				Name:  "add",
				Usage: "Add a new task",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "heading",
						Required: true,
						Usage:    "Heading of the task",
					},
					cli.StringFlag{
						Name:     "desc",
						Required: true,
						Usage:    "Description of the task",
					},
					cli.UintFlag{
						Name:     "deadline",
						Required: true,
						Usage:    "Deadline in number of days",
					},
				},
				Action: func(ctx *cli.Context) error {
					heading := ctx.String("heading")
					desc := ctx.String("desc")
					deadline := ctx.Int("deadline")
					numHrs := 24 * time.Hour * time.Duration(deadline)
					err := todo.GetStorage().Add(todo.NewTaskItem(
						heading, desc, time.Now().Add(numHrs)))
					if err == nil {
						fmt.Println("Added...")
					}
					return err
				},
			},
			cli.Command{
				Name:  "list",
				Usage: "List task items",
				Action: func(ctx *cli.Context) error {
					tl, err := todo.GetStorage().RetrieveAll(0, 10)
					if err != nil {
						return err
					}
					for _, ti := range tl {
						fmt.Println(ti)
					}
					return nil
				},
			},
			cli.Command{
				Name:  "serve",
				Usage: "Start the API server",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:  "port",
						Usage: "Port for the server to run",
						Value: 8080,
					},
				},
				Action: func(ctx *cli.Context) error {
					port := ctx.Int("port")
					return api.Run(port)
				},
			},
		},
	}
}
