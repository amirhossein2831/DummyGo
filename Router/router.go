package Router

import (
	"github.com/amirhossein2831/DummyGo/Error/Middleware"
	"github.com/go-chi/chi/v5"
)

func Route() *chi.Mux {
	router := chi.NewRouter()
	Middleware.CorsMiddleware(router)

	return router
}
