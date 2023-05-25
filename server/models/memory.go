package models

import "time"

type Memory struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"userId" db:"user_id"`
	CoverUrl  string    `json:"coverUrl" db:"cover_url"`
	Content   string    `json:"content" db:"content"`
	IsPublic  bool      `json:"isPublic" db:"is_public"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
