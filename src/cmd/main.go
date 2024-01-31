package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gleblagov/tagtour-events/config"
	"github.com/gleblagov/tagtour-events/data"
	"github.com/gleblagov/tagtour-events/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		LogLatency:  true,
		LogMethod:   true,
		LogRemoteIP: true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("host", v.RemoteIP),
					slog.String("uri", v.URI),
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.Duration("latency", v.Latency),
					slog.String("error", v.Error.Error()),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("host", v.RemoteIP),
					slog.String("uri", v.URI),
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.Duration("latency", v.Latency),
				)
			}
			return nil
		},
	}))
	config, err := config.NewStorageConfig()
	if err != nil {
		logger.Error("reading config", "error", err.Error())
	}
	postgres, err := data.NewPostgreStorage(config)
	if err != nil {
		logger.Error("connecting to db", "error", err.Error())
	}
	handler := handlers.NewEventsHandler(postgres)

	app.GET("/version", handler.HealthCheckVersion)

	app.GET("/", handler.GetAllEvents)
	app.GET("/:id", handler.GetEventById)
	app.POST("/", handler.CreateEvent)
	app.PATCH("/:id", handler.UpdateEvent)
	app.DELETE("/:id", handler.DeleteEvent)

	logger.Info("starting server", "port", 1234)
	err = app.Start(":1234")
	if err != nil {
		logger.Error("starting server", "error", err.Error())
	}
}
