package models

import (
	"database/sql"
	"logistica/app/config"
	"time"

	"gorm.io/gorm"
)

type StockRecord struct {
	ID          sql.NullInt64  `gorm:"primaryKey" json:"id"`
	Amount      sql.NullInt64  `gorm:"column:amount" json:"amount"`
	Before      sql.NullInt64  `gorm:"column:before_record" json:"before_record"`
	After       sql.NullInt64  `gorm:"column:after_record" json:"after_record"`
	Description sql.NullString `gorm:"column:description" json:"description"`
	IsAddition  sql.NullInt16  `gorm:"column:is_addition" json:"is_addition"`
	// Foreign Key
	Product   Product       `gorm:"foreignKey:ProductID" json:"product"`
	ProductID sql.NullInt64 `gorm:"column:product_id" json:"product_id"`
	// Timestamp
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (sr *StockRecord) FindAll() ([]StockRecord, error) {
	var stockRecords []StockRecord
	var db = config.ConnectGormDB()

	query := "SELECT * FROM stock_records;"
	results := db.Raw(query).Scan(&stockRecords)
	if results.Error != nil {
		return nil, results.Error
	}

	return stockRecords, nil
}

func (sr *StockRecord) NewRecord() error {
	var db = config.ConnectGormDB()

	result := db.Exec(
		`INSERT INTO stock_records(amount, before_record, after_record, description, is_addition, product_id) VALUES(?, ?, ?, ?, ?, ?);`,
		sr.Amount.Int64, sr.Before.Int64, sr.After.Int64, sr.Description.String, sr.IsAddition.Int16, sr.ProductID.Int64,
	)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

/*
func (sr *StockRecord) FindAll() ([]map[string]any, error) {
	var stockRecords []map[string]any
	var db = config.ConnectGormDB()

	query := "SELECT * FROM stock_records;"
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&sr.ID,
			&sr.Amount,
			&sr.Before,
			&sr.After,
			&sr.Description,
			&sr.IsAddition,
			&sr.ProductID,
			&sr.CreatedAt,
			&sr.DeletedAt,
			&sr.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		var data = map[string]interface{}{
			"id":          sr.ID.Int64,
			"amount":      sr.Amount.Int64,
			"before":      sr.Before.Int64,
			"after":       sr.After.Int64,
			"is_addition": IsAddition(int(sr.IsAddition.Int16)),
			"product":     sr.ProductID.Int64,
			"description": sr.Description.String,
		}

		stockRecords = append(stockRecords, data)
	}

	return stockRecords, nil
}
*/
