package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"logistica/app/config"
	"logistica/app/utility"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var mutex sync.Mutex

type Employee struct {
	ID          sql.NullInt64  `json:"id"`
	Name        sql.NullString `json:"name"`
	Address     sql.NullString `json:"address"`
	NumberPhone sql.NullString `json:"number_phone"`
	Position    sql.NullInt64  `json:"position"`
	IsUser      sql.NullBool   `json:"is_user"`
	IsSuperuser sql.NullBool   `json:"is_superuser"`
}

func (e *Employee) GetById(id int64) error {
	var db = config.ConnectDB()
	defer db.Close()

	var query string = "SELECT id, name, address, number_phone, position_id, is_user, is_superuser FROM employees WHERE id = ?"
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

func (e *Employee) NewEmployee(id int, name, address, numberPhone string, position int) (sql.Result, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectDB()
	defer db.Close()

	var query string = "INSERT INTO employees VALUES(?, ?, ?, ?, ?, 0, 0)"
	result, err := db.Exec(query, id, name, address, numberPhone, position)
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

func (e *Employee) DeleteEmployee(id int) error {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectDB()
	defer db.Close()

	var query string = "DELETE FROM employees WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return errors.New("race condition, id has been taken")
			}
		}
		return err
	}

	return nil
}

func (e *Employee) FindAll() ([]map[string]any, error) {
	var db = config.ConnectDB()
	defer db.Close()

	var query string = "SELECT e.id, e.name as employee_name, e.address, e.number_phone, e.position_id, e.is_user, e.is_superuser, p.name AS position_name FROM employees e JOIN positions p ON e.position_id = p.id WHERE e.id > 1;"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		cancel()
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var employees []map[string]any
	var positionName sql.NullString

	for rows.Next() {
		err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Address,
			&e.NumberPhone,
			&e.Position,
			&e.IsUser,
			&e.IsSuperuser,
			&positionName,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		var employee = map[string]any{
			"id":           e.ID.Int64,
			"name":         utility.Capitalize(e.Name.String),
			"address":      utility.Capitalize(e.Address.String),
			"positionName": utility.Capitalize(positionName.String),
			"numberPhone":  utility.Capitalize(e.NumberPhone.String),
			"isUser":       e.IsUser.Bool,
			"isSuperuser":  e.IsSuperuser.Bool,
		}

		employees = append(employees, employee)
	}
	return employees, nil
}

func (e *Employee) LastId() (int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectDB()
	defer db.Close()

	var lastId int
	var query string = "SELECT COALESCE(MAX(id), 100020) FROM employees"
	err := db.QueryRow(query).Scan(
		&lastId,
	)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (e *Employee) Count() (int, error) {
	var db = config.ConnectDB()
	defer db.Close()

	var totalEmployee int
	var query = "SELECT COUNT(*) AS total FROM employees WHERE id > 1"
	err := db.QueryRow(query).Scan(&totalEmployee)
	if err != nil {
		return 0, err
	}

	return totalEmployee, nil
}
