package models

import (
	"database/sql"
	"logistica/app/config"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       sql.NullInt64  `json:"id" gorm:"primaryKey;column:id"`
	Username sql.NullString `json:"username" gorm:"column:username"`
	Password sql.NullString `json:"password" gorm:"column:password"`

	// Employee   Employee      `gorm:"foreignKey:EmployeeeID" json:"employees"`
	// EmployeeID sql.NullInt64 `gorm:"column:employee_id" json:"employee_id"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (u *User) FindAll() []User {
	var users = []User{
		{
			Username: sql.NullString{String: "dhimasarista"},
			Password: sql.NullString{String: "01052002"},
		},
	}
	return users
}

func (u *User) GetByUsername(username string) error {
	var db = config.ConnectGormDB()

	results := db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&u)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (u *User) GetByID(id int) error {
	var db = config.ConnectGormDB()

	results := db.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&u)
	if results.Error != nil {
		return results.Error
	}

	return nil
}
