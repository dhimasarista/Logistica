package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Earning struct {
	ID   sql.NullInt64  `gorm:"primaryKey;column:id" json:"id"`
	Name sql.NullString `gorm:"column:name" json:"name"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
