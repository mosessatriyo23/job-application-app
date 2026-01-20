package database

import (
	"log"

	"job-application-app/backend/internal/config"
	"job-application-app/backend/internal/model"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&model.User{},
		&model.Company{},
		&model.JobCategory{},
		&model.Job{},
		&model.Application{},
	)

	if err != nil {
		log.Fatal("migration failed:", err)
	}

	log.Println("database migrated")
}
