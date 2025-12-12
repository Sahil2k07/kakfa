package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	Profile struct {
		gorm.Model
		UserID    uint
		FirstName string `gorm:"not null"`
		LastName  string `gorm:"not null"`
		Phone     string
		Country   string
	}

	User struct {
		gorm.Model
		Email    string  `gorm:"not null; uniqueIndex"`
		UserName string  `gorm:"not null"`
		Password string  `gorm:"not null"`
		Profile  Profile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}

	RUser struct {
		ID        string
		Email     string
		UserName  string
		Password  string
		PrimaryID uint
		FirstName string
		LastName  string
		Phone     string
		Country   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
