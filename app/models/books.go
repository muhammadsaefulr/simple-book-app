package models

import (
	"time"

	"gorm.io/gorm"
)

type Detail struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Book         string         `json:"book"`
	BookCategory string         `json:"BookCategory"`
	AddedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
