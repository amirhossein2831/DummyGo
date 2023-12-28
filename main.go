package main

import (
	"fmt"
	"github.com/amirhossein2831/DummyGo/Error"
	"github.com/amirhossein2831/DummyGo/Router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	Error.CheckError(err)

	router := Router.Route()

	log.Printf("The app start in port %v:", os.Getenv("APP_PORT"))
	err = http.ListenAndServe(
		fmt.Sprintf(":%v", os.Getenv("APP_PORT")),
		router,
	)
	Error.CheckError(err)
}
