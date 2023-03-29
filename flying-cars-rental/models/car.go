package models

import "time"

type Car struct {
	ID        int64     `json:"id"`
	Model     string    `json:"model"`
	Make      string    `json:"make"`
	Year      int       `json:"year"`
	Available bool      `json:"available"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
