package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Handler http.HandlerFunc
	Method  string
	Path    string
}

type WebServer struct {
	Router          chi.Router
	Handlers        map[string]Handler
	WebServerPort   string
	AuthRealm       string
	AuthCredentials map[string]string
	AuthSecret      string
}

func NewWebServer(webServerPort, authRealm, authClientId, authClientSecret string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]Handler),
		WebServerPort: webServerPort,
		AuthRealm:     authRealm,
		AuthCredentials: map[string]string{
			authClientId: authClientSecret,
		},
	}
}

func (ws *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	ws.Handlers[path+"-"+method] = Handler{
		Handler: handler,
		Method:  method,
		Path:    path,
	}
}

func (ws *WebServer) Start() {
	ws.Router.Use(middleware.Logger)
	fmt.Println(ws.AuthRealm, ws.AuthCredentials)
	ws.Router.Use(middleware.BasicAuth(ws.AuthRealm, ws.AuthCredentials))

	for _, handler := range ws.Handlers {
		switch handler.Method {
		case "GET":
			ws.Router.Get(handler.Path, handler.Handler)
		case "POST":
			ws.Router.Post(handler.Path, handler.Handler)
		case "PUT":
			ws.Router.Put(handler.Path, handler.Handler)
		case "DELETE":
			ws.Router.Delete(handler.Path, handler.Handler)
		default:
			ws.Router.Head(handler.Path, handler.Handler)
		}
	}

	fmt.Println("🚀 Starting web server on port", ws.WebServerPort)
	err := http.ListenAndServe(":"+ws.WebServerPort, ws.Router)
	if err != nil {
		panic(err)
	}
}
