package web

import (
	"net/http"

	customer "github.com/adrianostankewicz/customer-favorites/internal/customer/service"
	"github.com/adrianostankewicz/customer-favorites/internal/infra/web/handler"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router        *chi.Mux
	WebServerPort string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		WebServerPort: port,
	}
}

func (w *WebServer) AddHandler(s *customer.CustomerService) *chi.Mux {

	w.Router.Use(middleware.Logger)
	w.Router.Use(middleware.Recoverer)

	customerHandler := handler.NewWebCustomerHandler(s)

	w.Router.Route("/customers", func(r chi.Router) {
		r.Post("/", customerHandler.Create)
		//r.Get("/{id}", customerHandler.FindByID)
		//r.Put("/{id}", customerHandler.Update)
		//r.Delete("/{id}", customerHandler.Delete)
	})

	return w.Router
}

func (s *WebServer) Start() {
	http.ListenAndServe(s.WebServerPort, s.Router)
}
