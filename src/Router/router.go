package Router

import (
	"github.com/amirhossein2831/DummyGo/src/Middleware"
	Router "github.com/amirhossein2831/DummyGo/src/Router/Routes"
	"github.com/go-chi/chi/v5"
)

func Route() *chi.Mux {
	router := chi.NewRouter()

	//Add Middleware Here
	Middleware.CorsMiddleware(router)

	//Api Routes
	router.Route("/api/v1", func(r chi.Router) {
		//Add Routes Here
		Router.HomeRoutes(r)
	})
	return router
}
