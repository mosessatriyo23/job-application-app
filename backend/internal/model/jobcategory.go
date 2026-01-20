package model

import "github.com/google/uuid"

type JobCategory struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `json:"name"`
}
