package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID            uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CompanyID     uuid.UUID  `json:"company_id" gorm:"type:uuid;not null"`
	CategoryID    *uuid.UUID `json:"category_id,omitempty"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Qualification string     `json:"qualification"`
	JobType       string     `json:"job_type"`
	Salary        int        `json:"salary"`
	Location      string     `json:"location"`
	Deadline      *time.Time `json:"deadline,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`

	Company  *Company     `json:"company,omitempty" gorm:"foreignKey:CompanyID"`
	Category *JobCategory `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}
