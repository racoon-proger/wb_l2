package eventserver

import (
	"encoding/json"
	"net/http"
)

type eventForWeekHandler struct {
	service service
}

// ServeHTTP реализует интерфейс http.Handler
func (h *eventForWeekHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := h.service.GetEventsForWeek(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.NewEncoder(w).Encode(&commonResponse{
		Result: &eventsResponse{
			Events: events,
		},
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}

// NewEventForWeekHandler возвращает новый экземпляр хэндлера
func NewEventForWeekHandler(service service) *eventForWeekHandler {
	return &eventForWeekHandler{
		service: service,
	}
}
