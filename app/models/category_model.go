package models

import (
	"database/sql"
	"logistica/app/config"
	"logistica/app/utility"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID   sql.NullInt64  `gorm:"primaryKey;column:id" json:"id"`
	Name sql.NullString `gorm:"column:name" json:"name"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (c *Category) FindAll() ([]map[string]interface{}, error) {
	db := config.ConnectGormDB()
	query := "SELECT id, name FROM product_categories;"

	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	// Map untuk menyimpan daftar manufacturer
	var categories []map[string]interface{}

	for rows.Next() {
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		var category = map[string]interface{}{
			"id":   c.ID.Int64,
			"name": utility.CapitalizeAll(c.Name.String),
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) NewCategory(id int64, name string) error {
	mutex.Lock()
	defer mutex.Unlock()

	db := config.ConnectGormDB()
	query := "INSERT INTO product_categories VALUES(?, ?, NOW(), NOW(), NULL)"

	if id <= 889 {
		id = 889
	}

	results := db.Exec(query)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (c *Category) LastId() (int64, error) {
	mutex.Lock()
	defer mutex.Unlock()

	db := config.ConnectGormDB()
	query := "SELECT COALESCE(MAX(id), 890) FROM product_categories;"

	var lastId int64
	resutls := db.Raw(query).Scan(&lastId)
	if resutls.Error != nil {
		return -1, resutls.Error
	}
	return lastId, nil
}
