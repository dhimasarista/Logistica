package models

import (
	"context"
	"database/sql"
	"log"
	"logistica/app/config"
	"logistica/app/utility"
	"time"
)

type Product struct {
	ID               sql.NullInt64  `json:"id"`
	Name             sql.NullString `json:"name"`
	SerialNumber     sql.NullString `json:"serial_number"`
	ManufacturerID   sql.NullInt64  `json:"manufacturer_id"`
	ManufacturerName sql.NullString `json:"manufacturer_name"`
	Stocks           sql.NullInt64  `json:"stocks"`
	Price            sql.NullInt64  `json:"price"`
	Weight           sql.NullInt64  `json:"weight"`
	CategoryID       sql.NullInt64  `json:"category_id"`
	CategoryName     sql.NullString `json:"category_name"`
}

func (p *Product) GetById(id int) error {
	var db = config.ConnectDB()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
	SELECT
		p.id, 
		p.name, 
		p.serial_number, 
		p.manufacturer_id, 
		p.stocks, 
		p.price, 
		p.weight, 
		p.category_id AS products, 
		m.name AS manufacturer_name, 
		c.name AS category_name 
	FROM 
		products p 
	JOIN 
		manufacturer m ON p.manufacturer_id = m.id 
	JOIN 
		product_category c ON p.category_id = c.id
	WHERE p.id = ?;
	`
	err := db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Name,
		&p.SerialNumber,
		&p.ManufacturerID,
		&p.Stocks,
		&p.Price,
		&p.Weight,
		&p.CategoryID,
		&p.ManufacturerName,
		&p.CategoryName,
	)
	if err != nil {

		return err
	}

	return nil
}

func (p *Product) FindAll() ([]map[string]interface{}, error) {
	var db = config.ConnectDB()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `
	SELECT
		p.id, 
		p.name, 
		p.serial_number, 
		p.manufacturer_id, 
		p.stocks, 
		p.price, 
		p.weight, 
		p.category_id AS products, 
		m.name AS manufacturer_name, 
		c.name AS category_name 
	FROM 
		products p 
	JOIN 
		manufacturer m ON p.manufacturer_id = m.id 
	JOIN 
		product_category c ON p.category_id = c.id;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		cancel()
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var products []map[string]interface{}

	for rows.Next() {
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.SerialNumber,
			&p.ManufacturerID,
			&p.Stocks,
			&p.Price,
			&p.Weight,
			&p.CategoryID,
			&p.ManufacturerName,
			&p.CategoryName,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		var product = map[string]interface{}{
			"id":              p.ID.Int64,
			"name":            utility.Capitalize(p.Name.String),
			"serial_number":   utility.Capitalize(p.SerialNumber.String),
			"manufacturer_id": p.ManufacturerID.Int64,
			"manufacturer":    utility.Capitalize(p.ManufacturerName.String),
			"stocks":          p.Stocks.Int64,
			"price":           p.Price.Int64,
			"weight":          p.Weight.Int64,
			"category_id":     p.CategoryID.Int64,
			"category":        utility.Capitalize(p.CategoryName.String),
		}

		products = append(products, product)
	}

	return products, nil
}
