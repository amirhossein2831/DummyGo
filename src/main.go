package main

import (
	"fmt"
	"github.com/amirhossein2831/DummyGo/src/App"
	"github.com/amirhossein2831/DummyGo/src/Error"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//Instantiate the app
	App.InitApp(&App.App)

	//Load env file
	err := godotenv.Load("./../.env")
	Error.CheckError(err)

	//Connect to DB

	//Migrate DB Table

	//Start HttpServer
	log.Printf("The app start in port %v:", os.Getenv("APP_PORT"))
	err = http.ListenAndServe(
		fmt.Sprintf(":%v", os.Getenv("APP_PORT")),
		App.App.Router,
	)
	Error.CheckError(err)
}
