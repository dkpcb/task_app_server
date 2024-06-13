package main

import (
	"github.com/dkpcb/TaskList-server/model"
	"github.com/dkpcb/TaskList-server/router"
	"github.com/labstack/echo/v4"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
}
