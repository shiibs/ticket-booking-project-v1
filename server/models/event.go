package models

import (
	"context"
	"time"
)

type Event struct {
	ID        uint
	Name      string
	Location  string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	CreateOne(ctx context.Context, event Event) (*Event, error)
}
