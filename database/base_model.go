package database

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Status    Status         `json:"status" gorm:"status;default:active"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"autoDeleteTime"`
}

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Deleted  Status = "deleted"
)
