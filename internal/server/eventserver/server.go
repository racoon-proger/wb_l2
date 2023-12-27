package eventserver

import (
	"context"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"wb_l2/internal/domain"
)

const (
	endpointCreateEvent    = "/create_event"
	endpointUpdateEvent    = "/update_event"
	endpointDeleteEvent    = "/delete_event"
	endpointEventsForDay   = "/events_for_day"
	endpointEventsForWeek  = "/events_for_week"
	endpointEventsForMonth = "/events_for_month"
)

type service interface {
	CreateEvent(ctx context.Context, event *domain.Event) (id string, err error)
	UpdateEvent(ctx context.Context, event *domain.Event) (err error)
	DeleteEvent(ctx context.Context, id string) (err error)
	GetEventsForDay(ctx context.Context) (events []domain.Event, err error)
	GetEventsForWeek(ctx context.Context) (events []domain.Event, err error)
	GetEventsForMonth(ctx context.Context) (events []domain.Event, err error)
}

type logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type server struct {
	service service
	logger  logger
	port    int
}

// Run runs the server
func (s *server) Run() error {
	r := mux.NewRouter()
	r.Handle(endpointCreateEvent, NewCreateEventHandler(s.service))
	r.Handle(endpointUpdateEvent, NewUpdateEventHandler(s.service))
	r.Handle(endpointDeleteEvent, NewDeleteEventHandler(s.service))
	r.Handle(endpointEventsForDay, NewEventsForDayHandler(s.service))
	r.Handle(endpointEventsForWeek, NewEventForWeekHandler(s.service))
	r.Handle(endpointEventsForMonth, NewEventForMonthHandler(s.service))
	r.Use(s.requestLoggerMiddleware(r, s.logger))
	return http.ListenAndServe(net.JoinHostPort("", strconv.Itoa(s.port)), r)
}

func (s *server) requestLoggerMiddleware(r *mux.Router, logger logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(w, req)
			logger.Info(
				"new request",
				"method",
				req.Method,
				"host",
				req.Host,
				"url",
				req.URL.Path,
			)
		})
	}
}

// NewServer returns a new server
func NewServer(service service, port int, logger logger) *server {
	return &server{
		service: service,
		logger:  logger,
		port:    port,
	}
}
