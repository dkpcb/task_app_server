package router

import (
	"os"

	"github.com/labstack/echo/v4/middleware"

	_ "net/http"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo) error {

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/tasks", GetTasksHandler)
	e.POST("/api/tasks", AddTaskHandler)
	e.PUT("/api/tasks/:taskID", ChangeFinishedTaskHandler)
	e.DELETE("/api/tasks/:taskID", DeleteTaskHandler)
	e.POST("/api/register", RegisterUserHandler)
	e.POST("/api/login", LoginUserHandler)

	err := e.Start(":8000")
	return err
}
