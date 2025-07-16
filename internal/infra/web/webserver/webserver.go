package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandlerInfo struct {
	Method  string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]HandlerInfo
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]HandlerInfo),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	newHandler := HandlerInfo{
		Method:  method,
		Handler: handler,
	}

	s.Handlers[path] = newHandler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Method(handler.Method, path, handler.Handler)
		// s.Router.Handle(path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
