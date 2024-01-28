package data

import (
	"time"
)

type Event struct {
	Id          int32     `json:"id"`
	Title       string    `json:"title"`
	Likes       int32     `json:"likes"`
	Media       []string  `json:"media"`
	Author      int32     `json:"author"`
	CreatedAt   time.Time `json:"createdAt"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Members     []int32   `json:"members"`
}

type EventBase struct {
	Title       string    `json:"title"`
	Media       []string  `json:"media"`
	Author      int32     `json:"author"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Members     []int32   `json:"members"`
}

func NewEvent(b *EventBase) *Event {
	return &Event{
		Id:          0,
		Title:       b.Title,
		Likes:       0,
		Media:       b.Media,
		Author:      b.Author,
		CreatedAt:   time.Now(),
		Date:        b.Date,
		Description: b.Description,
		Members:     b.Members,
	}
}
