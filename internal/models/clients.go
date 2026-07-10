package models

import "time"

type Client struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Document  string    `json:"document"`
	CreatedAt time.Time `json:"created_at"`
}
