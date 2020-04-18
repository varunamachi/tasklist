package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/varunamachi/tasklist/srv/todo"
)

func Run() {
	e := echo.New()

	listTasksHandler := func(ctx echo.Context) error {
		offset := getInt(ctx, "offset")
		limit := getInt(ctx, "limit")
		tl, err := todo.GetStorage().RetrieveAll(offset, limit)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, tl)
		// ctx.HTML(http.StatusOK, "<b>Hello</b>")
		return nil
	}
	e.GET("/api/tasks/:offset/:limit", listTasksHandler)

	createTaskHandler := func(ctx echo.Context) error {
		var ti todo.TaskItem
		err := ctx.Bind(&ti)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		if ti.Heading == "" {
			return fmt.Errorf("Invalid data given")
		}
		err = todo.GetStorage().Add(&ti)
		if err == nil {
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"message": "Task item created",
			})
		}
		return err
	}
	e.POST("/api/task", createTaskHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func getInt(ctx echo.Context, key string) int {
	val := ctx.Param(key)
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Error("Failed convert %s to int", key)
	}
	return num
}
