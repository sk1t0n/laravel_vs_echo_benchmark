package main

import (
	"web_app/internal/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	todoHandler := handler.NewTodoHandler()
	defer todoHandler.CloseDB()

	e.GET("/todos", todoHandler.GetTodosHandler)
	e.GET("/todos/:id", todoHandler.GetTodoHandler)
	e.POST("/todos", todoHandler.CreateTodoHandler)
	e.PUT("/todos/:id", todoHandler.UpdateTodoHandler)
	e.DELETE("/todos/:id", todoHandler.DeleteTodoHandler)

	e.Logger.Fatal(e.Start(":9001"))
}
