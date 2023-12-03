package models

"database/sql"

type Employee struct {
	ID          sql.NullInt64    `json:id`
	Name        sql.NullString `json:name`
	Address     sql.NullString `json:address`
	NumberPhone sql.NullString `json:number_phone`
	Position    sql.NullString `json:position`
	IsUser      sql.NullBool `json:is_user`
	IsSuperuser sql.NullBool   `json:is_superuser`
}


func (e *Employee) GetById(id int) error {
	db := config.ConnectDB()
	var query string = "SELECT id, name, address, number_phone, position, is_user, is_superuser WHERE id = ?"
	
}