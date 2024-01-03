package main

import (
	"fmt"
	"github.com/amirhossein2831/DummyGo/src/App"
	"github.com/amirhossein2831/DummyGo/src/DB"
	"github.com/amirhossein2831/DummyGo/src/Error"
	"github.com/amirhossein2831/DummyGo/src/Model/Migrator"
	"github.com/amirhossein2831/DummyGo/src/Router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//Instantiate the app
	App.InitApp(&App.App)

	//instantiate router
	router := Router.Route()

	//Load env file
	err := godotenv.Load("./../.env")
	Error.CheckError(err)

	//Connect to DB
	App.App.DB, err = DB.Connect()
	Error.CheckError(err)

	//Migrate DB Table
	err = Migrator.MigrateUp()
	Error.CheckError(err)

	//Start HttpServer
	log.Println("Connected to DB successfully.")
	log.Printf("The app start in port %v: \n", os.Getenv("APP_PORT"))
	err = http.ListenAndServe(
		fmt.Sprintf(":%v", os.Getenv("APP_PORT")),
		router,
	)
	Error.CheckError(err)
}
