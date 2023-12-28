package models

import "database/sql"

type User struct {
	ID       sql.NullInt64  `json:"id" gorm:"primaryKey;column:id"`
	Username sql.NullString `json:"username" gorm:"column:username"`
	Password sql.NullString `json:"password" gorm:"column:password"`
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
