package models

import (
	"database/sql"
	"logistica/app/config"
	"logistica/app/utility"
	"sync"
	"time"

	"gorm.io/gorm"
)

var mutex sync.Mutex

type Employee struct {
	// gorm.Model
	ID           sql.NullInt64  `json:"id" gorm:"primaryKey;column:id"`
	Name         sql.NullString `json:"name" gorm:"column:name"`
	Address      sql.NullString `json:"address" gorm:"column:address"`
	NumberPhone  sql.NullString `json:"number_phone" gorm:"column:number_phone"`
	IsUser       sql.NullBool   `json:"is_user" gorm:"column:is_user"`
	IsSuperuser  sql.NullBool   `json:"is_superuser" gorm:"column:is_superuser"`
	IdentityCard sql.NullByte   `json:"identity_card" gorm:"column:identity_card"`

	// Foreign key, memiliki relasi dengan Position model
	Position   Position      `gorm:"foreignKey:PositionID" json:"position"`
	PositionID sql.NullInt64 `gorm:"column:position_id" json:"position_id"`

	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (e *Employee) GetById(id int64) error {
	db := config.ConnectGormDB()
	query := "SElECT * FROM employees WHERE id = ?;"

	results := db.Raw(query, id).Scan(&e)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (e *Employee) NewEmployee() error {
	mutex.Lock()
	defer mutex.Unlock()

	db := config.ConnectGormDB()
	query := "INSERT INTO employees(id, name, address, number_phone, position_id, is_user, is_superuser, created_at, updated_at, deleted_at) VALUES(?, ?, ?, ?, ?, 0, 0, NOW(), NOW(), NULL);"

	results := db.Exec(
		query,
		e.ID.Int64,
		e.Name.String,
		e.Address.String,
		e.NumberPhone.String,
		e.PositionID.Int64,
	)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (e *Employee) UpdateEmployee() error {
	mutex.Lock()
	defer mutex.Unlock()

	db := config.ConnectGormDB()
	query := "UPDATE employees SET name = ?, address = ?, number_phone = ?, position_id = ? WHERE id = ?;"

	results := db.Exec(
		query,
		e.Name.String,
		e.Address.String,
		e.NumberPhone.String,
		e.PositionID.Int64,
		e.ID.Int64,
	)
	if results != nil {
		return results.Error
	}

	return nil
}

// Soft Delete
func (e *Employee) DeleteEmployee(id int64) error {
	db := config.ConnectGormDB()
	query := "UPDATE employees SET deleted_at = NOW() WHERE id = ?;"

	results := db.Exec(query, id)
	if results.Error != nil {
		return results.Error
	}
	return nil
}

// Restore from Soft Delete
// func (e *Employee) RestoreEmployee(id int) error {}

// Hard Delete
// func (e *Employee) DeleteEmployeePermanent(id int) error {}
func (e *Employee) FindAll() ([]map[string]any, error) {
	db := config.ConnectGormDB()
	query := `
		SELECT 
			e.id, 
			e.name as employee_name, 
			e.address, 
			e.number_phone, 
			e.position_id, 
			e.is_user, 
			e.is_superuser, 
			p.name AS position_name 
		FROM 
			employees e 
		JOIN positions p ON e.position_id = p.id 
		WHERE e.id > 1 AND e.deleted_at IS NULL;`

	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []map[string]interface{}
	for rows.Next() {
		err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Address,
			&e.NumberPhone,
			&e.PositionID,
			// &e.IdentityCard,
			&e.IsUser,
			&e.IsSuperuser,
			&e.Position.Name,
			// &e.CreatedAt,
			// &e.UpdatedAt,
			// &e.DeletedAt,

		)

		if err != nil {
			return nil, err
		}

		var employee = map[string]any{
			"id":            e.ID.Int64,
			"name":          utility.Capitalize(e.Name.String),
			"address":       utility.Capitalize(e.Address.String),
			"position_name": utility.Capitalize(e.Position.Name.String),
			"number_phone":  utility.Capitalize(e.NumberPhone.String),
			"is_user":       e.IsUser.Bool,
			"is_superuser":  e.IsSuperuser.Bool,
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (e *Employee) LastId() (int, error) {
	db := config.ConnectGormDB()
	query := "SELECT COALESCE(MAX(id), 100020) FROM employees;"

	var lastId int
	results := db.Raw(query).Scan(&lastId)
	if results.Error != nil {
		return -1, results.Error
	}

	return lastId, nil
}

func (e *Employee) Count() (int, error) {
	db := config.ConnectGormDB()
	query := "SELECT COUNT(*) AS total FROM employees WHERE id > 1;"

	var totalEmployee int
	results := db.Raw(query).Scan(&totalEmployee)
	if results.Error != nil {
		return -1, results.Error
	}

	return totalEmployee, nil
}
