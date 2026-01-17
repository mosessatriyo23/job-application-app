package model

import "time"

type Company struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	LogoURL   string    `json:"logo_url"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
