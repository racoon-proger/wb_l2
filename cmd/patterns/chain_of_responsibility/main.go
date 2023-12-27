package main

import "log"

func main() {
	router := NewRouter([]handler{
		&logHandler{},
		&metricsHandler{},
		&notificationHandler{},
	})
	err := router.Route("some task")
	if err != nil {
		log.Fatal(err)
	}
}

type handler interface {
	Handle(task string) (err error)
}

type router struct {
	handlers []handler
}

func (r *router) Route(task string) (err error) {
	for i := range r.handlers {
		err = r.handlers[i].Handle(task)
		if err != nil {
			return
		}
	}
	return
}

func NewRouter(handlers []handler) *router {
	return &router{
		handlers: handlers,
	}
}

type logHandler struct {
}

func (h *logHandler) Handle(task string) (err error) {
	// write logs
	return nil
}

type metricsHandler struct {
}

func (h *metricsHandler) Handle(task string) (err error) {
	// write metrics
	return nil
}

type notificationHandler struct {
}

func (h *notificationHandler) Handle(task string) (err error) {
	// send notifications
	return nil
}
