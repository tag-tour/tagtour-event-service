package main

import (
	"github.com/gleblagov/tagtour-events/data"
	"github.com/gleblagov/tagtour-events/handlers"
	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	postgres, err := data.NewPostgreStorage("admin", "admin", "events")
	if err != nil {
		panic(err)
	}
	handler := handlers.NewEventsHandler(postgres)

	app.GET("/version", handler.HealthCheckVersion)

	app.GET("/", handler.GetAllEvents)
	app.GET("/:id", handler.GetEventById)
	app.POST("/", handler.CreateEvent)
	app.PATCH("/:id", handler.UpdateEvent)
	app.DELETE("/:id", handler.DeleteEvent)

	app.Start(":1234")
}
