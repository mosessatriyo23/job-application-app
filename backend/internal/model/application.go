package model

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	JobID        uuid.UUID `json:"job_id" gorm:"type:uuid;not null"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	FullName     string    `json:"full_name"`
	BirthDate    time.Time `json:"birth_date"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	CVURL        string    `json:"cv_url"`
	PortfolioURL string    `json:"portfolio_url"`
	CoverLetter  string    `json:"cover_letter"`
	Status       string    `json:"status"`
	AppliedAt    time.Time `json:"applied_at"`

	Job  *Job  `json:"job,omitempty" gorm:"foreignKey:JobID"`
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
