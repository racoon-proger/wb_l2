package eventserver

import (
	"encoding/json"
	"net/http"
)

type eventsForDayHandler struct {
	service service
}

// ServeHTTP реализует интерфейс http.Handler
func (h *eventsForDayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := h.service.GetEventsForDay(r.Context())
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

// NewEventsForDayHandler возвращает новый экземпляр хэндлера
func NewEventsForDayHandler(service service) *eventsForDayHandler {
	return &eventsForDayHandler{
		service: service,
	}
}
