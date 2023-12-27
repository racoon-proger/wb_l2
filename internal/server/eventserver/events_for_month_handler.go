package eventserver

import (
	"encoding/json"
	"net/http"
)

type eventForMonthHandler struct {
	service service
}

// ServeHTTP реализует интерфейс http.Handler
func (h *eventForMonthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := h.service.GetEventsForMonth(r.Context())
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

// NewEventForMonthHandler возвращает новый экземпляр хэндлера
func NewEventForMonthHandler(service service) *eventForMonthHandler {
	return &eventForMonthHandler{
		service: service,
	}
}
