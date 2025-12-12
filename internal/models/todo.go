package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Todo struct {
		gorm.Model
		Title       string `gorm:"not null"`
		UserID      uint   `gorm:"not null"`
		User        User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		Description string
		CompletedAt *time.Time
		Status      string `gorm:"not null; default:'PENDING'"`
	}

	RTodo struct {
		ID          string
		PrimaryID   uint
		Title       string
		UserID      uint
		Description string
		Status      string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		CompletedAt *time.Time
	}
)
