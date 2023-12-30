package App

import (
	"github.com/amirhossein2831/DummyGo/src/Router"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"os"
)

var App Application

type Application struct {
	Domain  string
	AppName string
	Router  *chi.Mux
	DB      *gorm.DB
}

func InitApp(app *Application) {
	app.Domain = os.Getenv("APP_DOMAIN")
	app.Domain = os.Getenv("APP_NAME")
	app.Router = Router.Route()
}
