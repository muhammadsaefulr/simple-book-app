package models

import (
	"time"

	"gorm.io/gorm"
)

type Detail struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	BookTitle    string         `json:"BookTitle"`
	BookCategory string         `json:"BookCategory"`
	BookAuthor   string         `json:"BookAuthor"`
	BookRating   string         `json:"BookRating"`
	AddedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
