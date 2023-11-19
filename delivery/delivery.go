package delivery

import (
	"github.com/labstack/echo/v4"
	"log"
	"mentoring-golang-rest-api/service"
)

func Route() {
	e := echo.New()

	e.POST("/todos", service.CreateTodo)
	e.GET("/todos", service.GetTodos)
	e.PATCH("/todos/:id", service.UpdateTodo)
	e.PATCH("/todos/:id", service.UpdateCheckFlag)
	e.DELETE("/todos/:id", service.DeleteTodo)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
