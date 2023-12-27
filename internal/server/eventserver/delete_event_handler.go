package eventserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type deleteEventHandler struct {
	service service
}

// ServeHTTP реализует интерфейс http.Handler
func (h *deleteEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request = &deleteEventRequest{}
	err := decodePOSTRequest(r.Body, request)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	err = h.validateDeleteRequest(request)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteEvent(r.Context(), request.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.NewEncoder(w).Encode(&commonResponse{})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *deleteEventHandler) validateDeleteRequest(d *deleteEventRequest) (err error) {
	if d.ID == "" {
		return fmt.Errorf("id is empty")
	}
	return nil
}

// NewDeleteEventHandler возвращает новый экземпляр хэндлера
func NewDeleteEventHandler(service service) *deleteEventHandler {
	return &deleteEventHandler{
		service: service,
	}
}
