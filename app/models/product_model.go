package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"logistica/app/config"
	"logistica/app/utility"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
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

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (p *Product) GetById(id int) error {
	var db = config.ConnectSQLDB()
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
	db := config.ConnectSQLDB()
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
		product_category c ON p.category_id = c.id
	WHERE p.deleted_at IS NULL;
	`

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

func (p *Product) NewProduct(id int, name, serialNumber string, manufacturer, stocks, price, weight, category int) (sql.Result, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectSQLDB()
	defer db.Close()

	// Kode ini adalah pernyataan SQL untuk menyisipkan data baru ke dalam tabel 'products'.
	// Nilai pada kolom pertama (index_column) ditentukan dengan menggunakan fungsi CASE:
	//   - Jika nilai maksimum dari kolom 'index_column' kurang dari 1020 atau NULL, maka nilainya diset menjadi 1020.
	//   - Jika nilai maksimum dari kolom 'index_column' tidak kurang dari 1020, maka nilainya diset menjadi nilai maksimum tersebut ditambah 1.
	var query string = `INSERT INTO products VALUES(?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), NULL);`
	result, err := db.Exec(query, id, name, serialNumber, manufacturer, stocks, price, weight, category)
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
func (p *Product) UpdateProduct(id int, name, serialNumber string, manufacturer, price, weight, category int) (sql.Result, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectSQLDB()
	defer db.Close()

	var query string = "UPDATE products SET name = ?, serial_number = ?, manufacturer_id = ?, price = ?, weight = ?, category_id = ? WHERE id = ?;"
	result, err := db.Exec(query, name, serialNumber, manufacturer, price, weight, category, id)
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

// Soft Delete
func (p *Product) DeleteProduct(id int) error {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectSQLDB()
	defer db.Close()

	var query string = "UPDATE products SET deleted_at = NOW() where id = ?;"
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
func (p *Product) CheckStock(id int) (int, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	var totalStocks int
	var query string = "SELECT stocks FROM products WHERE id = ?;"
	err := db.QueryRow(query, id).Scan(&totalStocks)
	if err != nil {
		return 0, err
	}

	return totalStocks, nil
}

func (p *Product) Count() (int, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	var totalProducts int
	var query string = "SELECT COUNT(*) AS total FROM products"
	err := db.QueryRow(query).Scan(&totalProducts)
	if err != nil {
		return 0, err
	}

	return totalProducts, nil
}

func (p *Product) UpdateStocks(id int, stocks int) (sql.Result, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	query := "UPDATE products SET stocks = ? WHERE id = ?"

	result, err := db.Exec(query, stocks, id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return nil, errors.New("race condition, id has been taken")
			}
		}
		return nil, err
	}

	return result, nil
}

func (p *Product) LastStocks(id int) (int, error) {
	var db = config.ConnectSQLDB()
	defer db.Close()

	var lastStock int
	var query string = "SELECT stocks FROM products WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&lastStock)
	if err != nil {
		return 0, err
	}

	return lastStock, nil
}

func (p *Product) LastId() (int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectSQLDB()
	defer db.Close()

	var lastId int
	// Declare a variable named 'query' of type string.
	var query string = "SELECT COALESCE(MAX(id), 1020) FROM products;"
	err := db.QueryRow(query).Scan(
		&lastId,
	)

	if err != nil {
		return 0, err
	}

	return lastId, nil
}
