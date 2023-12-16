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
	db := config.ConnectDB()
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
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var products []map[string]interface{}

	for rows.Next() {
		var productData Product // Use a separate variable to scan data into

		err := rows.Scan(
			&productData.ID,
			&productData.Name,
			&productData.SerialNumber,
			&productData.ManufacturerID,
			&productData.Stocks,
			&productData.Price,
			&productData.Weight,
			&productData.CategoryID,
			&productData.ManufacturerName,
			&productData.CategoryName,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		product := map[string]interface{}{
			"id":              productData.ID.Int64,
			"name":            utility.Capitalize(productData.Name.String),
			"serial_number":   utility.CapitalizeAll(productData.SerialNumber.String),
			"manufacturer_id": productData.ManufacturerID.Int64,
			"manufacturer":    utility.CapitalizeAll(productData.ManufacturerName.String),
			"stocks":          productData.Stocks.Int64,
			"price":           utility.RupiahFormat(productData.Price.Int64),
			"weight":          productData.Weight.Int64,
			"category_id":     productData.CategoryID.Int64,
			"category":        utility.Capitalize(productData.CategoryName.String),
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *Product) Count() (int, error) {
	var db = config.ConnectDB()
	defer db.Close()

	var totalProducts int
	var query string = "SELECT COUNT(*) AS total FROM products"
	err := db.QueryRow(query).Scan(&totalProducts)
	if err != nil {
		return 0, err
	}

	return totalProducts, nil
}
