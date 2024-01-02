package models

import (
	"database/sql"
	"logistica/app/config"
	"logistica/app/utility"
	"time"

	"gorm.io/gorm"
)

type Manufacturer struct {
	ID   sql.NullInt64  `gorm:"primaryKey" json:"id"`
	Name sql.NullString `gorm:"column:name" json:"name"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (m *Manufacturer) FindAll() ([]map[string]interface{}, error) {
	db := config.ConnectGormDB()
	query := "SELECT * FROM manufacturers"

	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	var manufacturers []map[string]interface{}
	for rows.Next() {
		err = rows.Scan(
			&m.ID,
			&m.Name,
			&m.UpdatedAt,
			&m.CreatedAt,
			&m.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		var manufacturer = map[string]any{
			"id":   m.ID.Int64,
			"name": utility.CapitalizeAll(m.Name.String),
		}

		manufacturers = append(manufacturers, manufacturer)
	}

	return manufacturers, nil
}

func (m *Manufacturer) NewManufacturer(id int, name string) error {
	mutex.Lock()
	defer mutex.Unlock()

	db := config.ConnectGormDB()
	var query string = "INSERT INTO manufacturers VALUES(?, ?, NOW(), NOW(), NULL)"

	if m.ID.Int64 <= 9100 {
		m.ID.Int64 = 9000
	}
	results := db.Exec(query, m.ID.Int64, m.Name.String)
	if results.Error != nil {
		return results.Error
	}

	return nil

}

func (m *Manufacturer) LastId() (int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	db := config.ConnectGormDB()

	var lastId int
	var query string = "SELECT COALESCE(MAX(id), 9100) FROM manufacturers;"
	results := db.Raw(query).Scan(&lastId)

	if results.Error != nil {
		return 0, results.Error
	}

	return lastId, nil
}
