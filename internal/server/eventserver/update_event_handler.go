package eventserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type updateEventHandler struct {
	service service
}

// ServeHTTP реализует интерфейс http.Handler
func (h *updateEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request = &updateEventRequest{}
	err := decodePOSTRequest(r.Body, request)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	err = h.validateRequest(request)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.UpdateEvent(r.Context(), request.GetEvent())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.NewEncoder(w).Encode(&commonResponse{
		Result: &updateEventResponse{
			ID:    request.ID,
			Title: request.Title,
			Date:  request.Date,
		},
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *updateEventHandler) validateRequest(r *updateEventRequest) (err error) {
	if r.ID == "" {
		return fmt.Errorf("id is empty")
	}
	if r.Title == "" {
		return fmt.Errorf("title is empty")
	}
	if r.Date.IsZero() {
		return fmt.Errorf("date is empty")
	}
	return nil
}

// NewUpdateEventHandler возвращает новый экземпляр хэндлера
func NewUpdateEventHandler(service service) *updateEventHandler {
	return &updateEventHandler{
		service: service,
	}
}
