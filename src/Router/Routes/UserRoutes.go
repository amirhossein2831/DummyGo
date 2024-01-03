package Router

import (
	"github.com/amirhossein2831/DummyGo/src/Controller"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(router chi.Router) {
	//Add User Routes
	router.Post("/user", Controller.StoreUser)
}
