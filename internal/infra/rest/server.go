package rest

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router        chi.Router
	Handlers      []Route
	WebServerPort string
}

type Route struct {
	path     string
	httpVerb string
	handler  http.HandlerFunc
}

func NewRoute(path string, httpVerb string, handler http.HandlerFunc) Route {
	return Route{
		path:     path,
		httpVerb: httpVerb,
		handler:  handler,
	}
}

func NewServer(serverPort string) *Server {
	return &Server{
		Router:        chi.NewRouter(),
		Handlers:      make([]Route, 0),
		WebServerPort: serverPort,
	}
}

func (s *Server) AddHandler(route Route) {
	s.Handlers = append(s.Handlers, route)
}

func (s *Server) Start() {
	s.Router.Use(middleware.Logger)
	for r := range s.Handlers {
		switch s.Handlers[r].httpVerb {
		case "GET":
			s.Router.Get(s.Handlers[r].path, s.Handlers[r].handler)
		case "POST":
			s.Router.Post(s.Handlers[r].path, s.Handlers[r].handler)
		default:
			panic("Invalid HTTP verb")
		}
	}
	err := http.ListenAndServe(s.WebServerPort, s.Router)
	if err != nil {
		log.Panic(err)
	}
}
