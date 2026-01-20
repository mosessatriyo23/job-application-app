package main

import (
	"job-application-app/backend/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	config.ConnectDatabase(cfg)
}
