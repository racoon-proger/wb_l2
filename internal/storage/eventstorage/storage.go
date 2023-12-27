package eventstorage

import (
	"context"
	"fmt"
	"sync"
	"time"

	"wb_l2/internal/domain"

	"github.com/google/uuid"
)

type storage struct {
	sync.RWMutex
	events map[string]domain.Event
}

// CreateEvents создает новое событие
func (s *storage) CreateEvent(ctx context.Context, event *domain.Event) (id string, err error) {
	s.Lock()
	id = uuid.New().String()
	event.ID = id
	s.events[id] = *event
	s.Unlock()
	return
}

// UpdateEvent обновляет событие
func (s *storage) UpdateEvent(ctx context.Context, event *domain.Event) (err error) {
	s.Lock()
	if _, ok := s.events[event.ID]; !ok {
		return fmt.Errorf("event not found")
	}
	s.events[event.ID] = *event
	s.Unlock()
	return
}

// DeleteEvent удаляет событие
func (s *storage) DeleteEvent(ctx context.Context, id string) (err error) {
	s.Lock()
	delete(s.events, id)
	s.Unlock()
	return
}

// GetEventsByDateRange возвращет события входящие во временной диапазон
func (s *storage) GetEventsByDateRange(
	ctx context.Context,
	start, end time.Time,
) (
	events []domain.Event,
	err error,
) {
	s.RLock()
	for _, event := range s.events {
		if event.Date.After(start) && event.Date.Before(end) {
			events = append(events, event)
		}
	}
	s.RUnlock()
	return
}

// NewStorage returns a new storage instance
func NewStorage() *storage {
	return &storage{
		events: make(map[string]domain.Event),
	}
}
