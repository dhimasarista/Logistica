package models

import (
	"context"
	"database/sql"
	"log"
	"logistica/app/config"
)

type Employee struct {
	ID          sql.NullInt64  `json:"id"`
	Name        sql.NullString `json:"name"`
	Address     sql.NullString `json:"address"`
	NumberPhone sql.NullString `json:"number_phone"`
	Position    sql.NullString `json:"position"`
	IsUser      sql.NullBool   `json:"is_user"`
	IsSuperuser sql.NullBool   `json:"is_superuser"`
}

func (e *Employee) GetById(id int) error {
	var db = config.ConnectDB()
	defer db.Close()

	var query string = "SELECT id, name, address, number_phone, position, is_user, is_superuser FROM employees WHERE id = ?"
	err := db.QueryRow(query, id).Scan(
		&e.ID,
		&e.Name,
		&e.Address,
		&e.NumberPhone,
		&e.Position,
		&e.IsUser,
		&e.IsSuperuser,
	)
	if err != nil {
		return err
	}
	return nil
}

func (e *Employee) FindAll() ([]map[string]any, error) {
	var db = config.ConnectDB()
	defer db.Close()

	var query string = "SELECT id, name, address, number_phone, position, is_user, is_superuser FROM employees WHERE id > 1"
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var employees []map[string]any

	for rows.Next() {
		err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Address,
			&e.NumberPhone,
			&e.Position,
			&e.IsUser,
			&e.IsSuperuser,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		var employee = map[string]any{
			"id":          e.ID.Int64,
			"name":        e.Name.String,
			"address":     e.Address.String,
			"position":    e.Position.String,
			"numberPhone": e.NumberPhone.String,
			"isUser":      e.IsUser.Bool,
			"isSuperuser": e.IsSuperuser.Bool,
		}

		employees = append(employees, employee)
	}
	return employees, nil
}

func (e *Employee) LastId() error {
	var db = config.ConnectDB()
	defer db.Close()

	var query string = "SELECT MAX(id) as maxId FROM employees"
	err := db.QueryRow(query).Scan(
		&e.ID,
	)

	if err != nil {
		return err
	}

	return nil
}
