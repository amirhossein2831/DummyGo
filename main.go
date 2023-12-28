package main

import (
	"github.com/amirhossein2831/DummyGo/Error"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	Error.CheckError(err)

	router := chi.NewRouter()

	log.Printf("The app start in port %v", os.Getenv("APP_PORT"))
	err = http.ListenAndServe(":"+os.Getenv("APP_PORT"), router)
	Error.CheckError(err)
}
