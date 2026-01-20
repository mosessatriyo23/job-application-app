package model

import (
	"time"
)

type Job struct {
	ID             int        `json:"id" gorm:"primaryKey"`
	CompanyID      int        `json:"company_id"`
	CategoryID     *int       `json:"category_id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Qualification  string     `json:"qualification"`
	JobType        string     `json:"job_type"`
	Salary         int        `json:"salary"`
	Location       string     `json:"location"`
	Deadline       *time.Time `json:"deadline"`
	CreatedAt      time.Time  `json:"created_at"`

	Company  *Company     `json:"company,omitempty"`
	Category *JobCategory `json:"category,omitempty"`
}
