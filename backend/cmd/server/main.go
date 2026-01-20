package main

import (
	"log"
	"net/http"

	"job-application-app/backend/internal/config"
	"job-application-app/backend/internal/database"
)

func main() {
	cfg := config.LoadConfig()
	config.ConnectDatabase(cfg)

	database.Migrate()

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
