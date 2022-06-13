package main

import (
	"os"

	"github.com/tomlister/ragg/web/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := handlers.Router()
	r.Run(":"+port)
}