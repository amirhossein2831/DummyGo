package Router

import (
	"github.com/amirhossein2831/DummyGo/src/Controller"
	"github.com/go-chi/chi/v5"
)

func HomeRoutes(router chi.Router) {
	//Add Home Routes
	router.Get("/", Controller.Home)
}
