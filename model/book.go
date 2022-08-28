package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	CreatedAt time.Time      `gorm:"<-:create;column:created_at;type:datetime;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index"`
}
