package main

import (
	"fmt"
	"github.com/amirhossein2831/DummyGo/Error"
	"github.com/amirhossein2831/DummyGo/src/App"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//instantiate the app
	App.InitApp(&App.App)

	//load env file
	err := godotenv.Load()
	Error.CheckError(err)

	log.Printf("The app start in port %v:", os.Getenv("APP_PORT"))
	err = http.ListenAndServe(
		fmt.Sprintf(":%v", os.Getenv("APP_PORT")),
		App.App.Router,
	)
	Error.CheckError(err)
}
