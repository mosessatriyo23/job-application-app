package model

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Name      string    `json:"name"`
	LogoURL   string    `json:"logo_url"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	CreatedAt time.Time `json:"created_at"`

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
