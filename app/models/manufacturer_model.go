package models

import (
	"context"
	"database/sql"
	"errors"
	"logistica/app/config"
	"logistica/app/utility"

	"github.com/go-sql-driver/mysql"
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
			"name": utility.CapitalizeAll(m.Name.String),
		}

		manufacturers = append(manufacturers, manufacturer)
	}

	return manufacturers, nil
}

func (m *Manufacturer) NewManufacturer(id int, name string) (sql.Result, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectDB()
	defer db.Close()

	// Jika id yang diterima di bawah 9100
	if id <= 9100 {
		id = 9000 // sebagai nilai set otomatis jika row belum ada
	}
	var query string = "INSERT INTO manufacturer VALUES(?, ?)"
	result, err := db.Exec(query, id, name)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return nil, errors.New("race condition, id has been taken")
			}
		}
		return result, err
	}

	return result, nil
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
