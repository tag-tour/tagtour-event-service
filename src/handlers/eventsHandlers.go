package handlers

import (
	"net/http"
	"strconv"

	"github.com/gleblagov/tagtour-events/data"
	"github.com/labstack/echo"
)

type eventsHandler struct {
	db data.Storage
}

func NewEventsHandler(s data.Storage) *eventsHandler {
	return &eventsHandler{
		db: s,
	}
}

func (h *eventsHandler) HealthCheckVersion(c echo.Context) error {
	vers, err := h.db.CheckVersion()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, vers)
}

func (h *eventsHandler) CreateEvent(c echo.Context) error {
	eb := new(data.EventBase)
	if err := c.Bind(eb); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	event := data.NewEvent(eb)
	id, err := h.db.CreateEvent(event)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	event.Id = id
	return c.JSON(http.StatusOK, event)
}

func (h *eventsHandler) GetAllEvents(c echo.Context) error {
	events, err := h.db.GetAllEvents()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, events)
}

func (h *eventsHandler) UpdateEvent(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	id := int32(id64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	eb := new(data.EventBase)
	if err := c.Bind(eb); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	updatedEvent, err := h.db.UpdateEvent(id, eb)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedEvent)

}
func (h *eventsHandler) DeleteEvent(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	id := int32(id64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.db.DeleteEvent(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusAccepted)
}
