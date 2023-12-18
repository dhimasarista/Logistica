package models

import (
	"context"
	"database/sql"
	"errors"
	"logistica/app/config"
	"logistica/app/utility"

	"github.com/go-sql-driver/mysql"
)

type Category struct {
	ID   sql.NullInt64  `json:"id"`
	Name sql.NullString `json:"name"`
}

func (c *Category) FindAll() ([]map[string]interface{}, error) {
	db := config.ConnectDB()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := "SELECT id, name FROM product_category;"

	rows, err := db.QueryContext(ctx, query)
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

func (c *Category) NewCategory(id int, name string) (sql.Result, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectDB()
	defer db.Close()

	// Jika id yang diterima di bawah 9100
	if id <= 889 {
		id = 890 // sebagai nilai set otomatis jika row belum ada
	}
	var query string = "INSERT INTO product_category VALUES(?, ?)"
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
func (c *Category) LastId() (int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectDB()
	defer db.Close()

	var lastId int
	var query string = "SELECT MAX(id) FROM product_category;"
	err := db.QueryRow(query).Scan(
		&lastId,
	)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}
