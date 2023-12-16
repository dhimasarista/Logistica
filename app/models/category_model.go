package models

import (
	"context"
	"database/sql"
	"logistica/app/config"
	"logistica/app/utility"
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
			"name": utility.Capitalize(c.Name.String),
		}

		categories = append(categories, category)
	}

	return categories, nil
}
