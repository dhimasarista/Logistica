package models

import (
	"database/sql"
	"logistica/app/config"
	"logistica/app/utility"
)

type Position struct {
	ID   sql.NullInt64  `json:"id"`
	Name sql.NullString `json:"name"`
}

func (p *Position) FindAll() ([]map[string]any, error) {
	var db = config.ConnectDB()
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
