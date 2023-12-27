package domain

import "time"

// Event это событие в календаре
type Event struct {
	ID    string    `json:"id,omitempty"`
	Title string    `json:"title,omitempty"`
	Date  time.Time `json:"date,omitempty"`
}
