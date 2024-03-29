package Router

import (
	"github.com/amirhossein2831/DummyGo/src/Controller"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(router chi.Router) {
	//Add User Routes
	router.Get("/users", Controller.UserIndex)
	router.Get("/users/{userID}", Controller.ShowUser)
	router.Post("/users", Controller.StoreUser)
}
