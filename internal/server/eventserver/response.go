package eventserver

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/exp/slog"

	"wb_l2/internal/domain"
)

type commonResponse struct {
	Result any `json:"result,omitempty"`
}

type errorResponse struct {
	Error string `json:"error,omitempty"`
}

type createEventResponse struct {
	ID string `json:"id,omitempty"`
}

type updateEventResponse struct {
	ID    string    `json:"id,omitempty"`
	Title string    `json:"title,omitempty"`
	Date  time.Time `json:"date,omitempty"`
}

type eventsResponse struct {
	Events []domain.Event `json:"events,omitempty"`
}

func writeError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(&errorResponse{
		Error: err.Error(),
	})
	if err != nil {
		slog.Error(err.Error())
	}
}
