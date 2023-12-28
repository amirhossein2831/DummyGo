package main

import (
	"github.com/amirhossein2831/DummyGo/Error"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	Error.CheckError(err)

}
