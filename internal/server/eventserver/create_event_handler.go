package eventserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type createEventHandler struct {
	service service
}

// ServeHTTP реализует интерфейс http.Handler
func (h *createEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request = &createEventRequest{}
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
	var id string
	id, err = h.service.CreateEvent(r.Context(), request.GetEvent())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.NewEncoder(w).Encode(&commonResponse{
		Result: &createEventResponse{
			ID: id,
		},
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *createEventHandler) validateRequest(r *createEventRequest) (err error) {
	if r.Title == "" {
		return fmt.Errorf("title is empty")
	}
	if r.Date.IsZero() {
		return fmt.Errorf("date is empty")
	}
	return nil
}

// NewCreateEventHandler возвращает новый экземпляр хэндлера
func NewCreateEventHandler(service service) *createEventHandler {
	return &createEventHandler{
		service: service,
	}
}
