package eventservice

import (
	"context"
	"time"

	"wb_l2/internal/domain"
)

type storage interface {
	CreateEvent(ctx context.Context, event *domain.Event) (id string, err error)
	UpdateEvent(ctx context.Context, event *domain.Event) (err error)
	DeleteEvent(ctx context.Context, id string) (err error)
	GetEventsByDateRange(
		ctx context.Context,
		start, end time.Time,
	) (
		events []domain.Event,
		err error,
	)
}

type service struct {
	storage storage
}

func (s *service) CreateEvent(ctx context.Context, event *domain.Event) (id string, err error) {
	id, err = s.storage.CreateEvent(ctx, event)
	return
}

func (s *service) UpdateEvent(ctx context.Context, event *domain.Event) (err error) {
	err = s.storage.UpdateEvent(ctx, event)
	return
}

func (s *service) DeleteEvent(ctx context.Context, id string) (err error) {
	err = s.storage.DeleteEvent(ctx, id)
	return
}

func (s *service) GetEventsForDay(ctx context.Context) (events []domain.Event, err error) {
	// start := 2023-12-26 00:00:00
	start := time.Now().Truncate(24 * time.Hour)
	// end := 2023-12-27 00:00:00
	end := time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour)
	events, err = s.storage.GetEventsByDateRange(ctx, start, end)
	return
}

func (s *service) GetEventsForWeek(ctx context.Context) (events []domain.Event, err error) {
	start := time.Now().Truncate(24 * time.Hour)
	end := time.Now().AddDate(0, 0, 8).Truncate(24 * time.Hour)
	events, err = s.storage.GetEventsByDateRange(ctx, start, end)
	return
}

func (s *service) GetEventsForMonth(ctx context.Context) (events []domain.Event, err error) {
	start := time.Now().AddDate(0, 0, 0).Truncate(24 * time.Hour)
	end := time.Now().AddDate(0, 1, 1).Truncate(24 * time.Hour)
	events, err = s.storage.GetEventsByDateRange(ctx, start, end)
	return
}

func NewService(storage storage) *service {
	return &service{
		storage: storage,
	}
}
