package models

import (
	"context"
	"database/sql"
	"logistica/app/config"
	"logistica/app/utility"
)

type Manufacturer struct {
	ID   sql.NullInt64  `json:"id"`
	Name sql.NullString `json:"name"`
}

func (m *Manufacturer) FindAll() ([]map[string]interface{}, error) {
	db := config.ConnectDB()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := "SELECT id, name FROM manufacturer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	// Map untuk menyimpan daftar manufacturer
	var manufacturers []map[string]interface{}

	for rows.Next() {
		err := rows.Scan(&m.ID, &m.Name)
		if err != nil {
			return nil, err
		}
		var manufacturer = map[string]interface{}{
			"id":   m.ID.Int64,
			"name": utility.Capitalize(m.Name.String),
		}

		manufacturers = append(manufacturers, manufacturer)
	}

	return manufacturers, nil
}

func (m *Manufacturer) LastId() (int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectDB()
	defer db.Close()

	var lastId int
	var query string = "SELECT MAX(id) FROM manufacturer;"
	err := db.QueryRow(query).Scan(
		&lastId,
	)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}
