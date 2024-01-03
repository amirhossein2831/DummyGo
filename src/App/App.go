package App

import (
	"gorm.io/gorm"
	"os"
)

var App Application

type Application struct {
	Domain  string
	AppName string
	DB      *gorm.DB
}

func InitApp(app *Application) {
	app.Domain = os.Getenv("APP_DOMAIN")
	app.Domain = os.Getenv("APP_NAME")
}
