package models

import (
	"database/sql"
	"logistica/app/config"
	"logistica/app/utility"
	"time"

	"gorm.io/gorm"
)

type Position struct {
	ID   sql.NullInt64  `gorm:"primaryKey;column:id" json:"id"`
	Name sql.NullString `gorm:"column:name" json:"name"`
	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (p *Position) FindAll() ([]map[string]any, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	var query string = "SELECT id, name FROM positions"
	rows, err := db.Query(query)

	var positions []map[string]interface{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&p.ID,
			&p.Name,
		)
		if err != nil {
			return nil, err
		}

		var position = map[string]any{
			"id":   p.ID.Int64,
			"name": utility.Capitalize(p.Name.String),
		}

		positions = append(positions, position)
	}
	return positions, nil
}

func (p *Position) NewPosition(id int, name string) (sql.Result, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	var query string = "INSERT INTO positions VALUES(?, ?)"
	result, err := db.Exec(query, id, name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Position) LastId() (int, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	var lastId int
	var query string = "SELECT COALESCE(MAX(id), 2000) FROM positions"
	err := db.QueryRow(query).Scan(
		&lastId,
	)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}
