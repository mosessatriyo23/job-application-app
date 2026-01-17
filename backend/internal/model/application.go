package model

import "time"

type Application struct {
	ID           int       `json:"id"`
	JobID        int       `json:"job_id"`
	UserID       int       `json:"user_id"`
	CVURL        string    `json:"cv_url"`
	PortfolioURL string    `json:"portfolio_url"`
	CoverLetter  string    `json:"cover_letter"`
	Status       string    `json:"status"`
	AppliedAt    time.Time `json:"applied_at"`

	Job  *Job  `json:"job,omitempty"`
	User *User `json:"user,omitempty"`
}
