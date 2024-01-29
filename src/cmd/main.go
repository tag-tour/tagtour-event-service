package main

import (
	"log"

	"github.com/gleblagov/tagtour-events/config"
	"github.com/gleblagov/tagtour-events/data"
	"github.com/gleblagov/tagtour-events/handlers"
	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	config, err := config.NewStorageConfig()
	if err != nil {
		log.Fatalf("error while reading config: %v\n", err)
	}
	postgres, err := data.NewPostgreStorage(config)
	if err != nil {
		log.Fatalf("error while connecting to db: %v\n", err)
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
