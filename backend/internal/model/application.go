package model

import (
	"time"
)

type Application struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	JobID        int       `json:"job_id"`
	UserID       int       `json:"user_id"`
	FullName     string    `json:"full_name"`
	BirthDate    time.Time `json:"birth_date"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	CVURL        string    `json:"cv_url"`
	PortfolioURL string    `json:"portfolio_url"`
	CoverLetter  string    `json:"cover_letter"`
	Status       string    `json:"status"`
	AppliedAt    time.Time `json:"applied_at"`

	Job  *Job  `json:"job,omitempty"`
	User *User `json:"user,omitempty"`
}
