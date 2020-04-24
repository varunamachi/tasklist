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
	defer func() {
		if todo.GetStorage() != nil {
			storage := todo.GetStorage()
			switch storage.(type) {
			case *todo.JSONStorage:
				jstor := storage.(*todo.JSONStorage)
				err = jstor.Close()
				if err != nil {
					log.Fatalln("Storage file close error:", err)
				} else {
					log.Println("Storage safely closed")
				}
			case *db.PostgresStorage:
				if db.Conn() != nil {
					err = db.Conn().Close()
					if err != nil {
						log.Fatalln("DB not closed:", err)
					} else {
						log.Println("DB safely closed")
					}
				}
			}
		}
	}()
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

			// storage := &db.PostgresStorage{}
			storage := &todo.JSONStorage{}
			err = storage.Init()
			if err != nil {
				log.Fatalf("Failed to initialize data source %s", err.Error())
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
				Name:  "remove",
				Usage: "Remove a existing task using ID",
				Flags: []cli.Flag{
					cli.UintFlag{
						Name:     "id",
						Required: true,
						Usage:    "ID of the Task item",
					},
				},
				Action: func(ctx *cli.Context) error {
					id := ctx.Int("id")
					err := todo.GetStorage().Remove(id)
					if err == nil {
						fmt.Println("The task, ", id,
							" has removed from the storage")
					}
					return err
				},
			},
			cli.Command{
				Name:  "update",
				Usage: "Update a new task",
				Flags: []cli.Flag{
					cli.UintFlag{
						Name:     "id",
						Required: true,
						Usage:    "ID of the Task item",
					},
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
					cli.StringFlag{
						Name:     "status",
						Required: true,
						Usage:    "Status of the task",
					},
					cli.UintFlag{
						Name:     "deadline",
						Required: true,
						Usage:    "Deadline in number of days",
					},
				},
				Action: func(ctx *cli.Context) error {
					id := ctx.Int("id")
					heading := ctx.String("heading")
					desc := ctx.String("desc")
					deadline := ctx.Int("deadline")
					status := ctx.String("status")
					numHrs := 24 * time.Hour * time.Duration(deadline)
					updatedItem := &todo.TaskItem{
						ID:          id,
						Heading:     heading,
						Description: desc,
						Status:      status,
						Deadline:    time.Now().Add(numHrs),
					}
					err := todo.GetStorage().Update(updatedItem)
					if err == nil {
						fmt.Println("Updated...")
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
