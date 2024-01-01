package models

import (
	"database/sql"
	"errors"
	"log"
	"logistica/app/config"
	"logistica/app/utility"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID               sql.NullInt64  `gorm:"primaryKey" json:"id"`
	Name             sql.NullString `gorm:"column:name" json:"name"`
	SerialNumber     sql.NullString `gorm:"column:serial_number" json:"serial_number"`
	ManufacturerID   sql.NullInt64  `gorm:"column:manufacturer_id" json:"manufacturer_id"`
	ManufacturerName sql.NullString `gorm:"column:manufacturer_name" json:"manufacturer_name"`
	Stocks           sql.NullInt64  `gorm:"column:stocks" json:"stocks"`
	Price            sql.NullInt64  `gorm:"column:price" json:"price"`
	Weight           sql.NullInt64  `gorm:"column:weight" json:"weight"`
	CategoryID       sql.NullInt64  `gorm:"column:category_id" json:"category_id"`
	CategoryName     sql.NullString `gorm:"column:category_id" json:"category_name"`

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// SLOW SQL >= 200ms
func (p *Product) GetById(id int) error {
	db := config.ConnectGormDB()
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
			manufacturers m ON p.manufacturer_id = m.id 
		JOIN 
			product_categories c ON p.category_id = c.id
		WHERE p.id = ?;`

	results := db.Raw(query, id).Scan(&p)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (p *Product) FindAll() ([]map[string]interface{}, error) {
	db := config.ConnectGormDB()

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
			manufacturers m ON p.manufacturer_id = m.id 
		JOIN 
			product_categories c ON p.category_id = c.id
		WHERE p.deleted_at IS NULL;`

	rows, err := db.Raw(query).Rows()
	if err != nil {
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

		product := map[string]interface{}{
			"id":              p.ID.Int64,
			"name":            utility.Capitalize(p.Name.String),
			"serial_number":   utility.CapitalizeAll(p.SerialNumber.String),
			"manufacturer_id": p.ManufacturerID.Int64,
			"manufacturer":    utility.CapitalizeAll(p.ManufacturerName.String),
			"stocks":          p.Stocks.Int64,
			"price":           utility.RupiahFormat(p.Price.Int64),
			"weight":          p.Weight.Int64,
			"category_id":     p.CategoryID.Int64,
			"category":        utility.Capitalize(p.CategoryName.String),
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *Product) NewProduct() error {
	mutex.Lock()
	defer mutex.Unlock()
	db := config.ConnectGormDB()
	var query string = `INSERT INTO products VALUES(?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), NULL);`
	results := db.Exec(
		query,
		p.ID.Int64,
		p.Name.String,
		p.SerialNumber.String,
		p.ManufacturerID.Int64,
		p.Stocks.Int64,
		p.Price.Int64,
		p.Weight.Int64,
		p.CategoryID.Int64,
	)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
func (p *Product) UpdateProduct() error {
	mutex.Lock()
	defer mutex.Unlock()

	db := config.ConnectGormDB()

	var query string = "UPDATE products SET name = ?, serial_number = ?, manufacturer_id = ?, price = ?, weight = ?, category_id = ? WHERE id = ?;"
	results := db.Exec(
		query,
		p.Name.String,
		p.SerialNumber.String,
		p.ManufacturerID.Int64,
		p.Price.Int64,
		p.Weight.Int64,
		p.CategoryID.Int64,
		p.ID.Int64,
	)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

// Soft Delete
func (p *Product) DeleteProduct(id int) error {
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectGormDB()
	var query string = "UPDATE products SET deleted_at = NOW() where id = ?;"
	results := db.Exec(query, id)
	if results.Error != nil {
		return results.Error
	}

	return nil
}
func (p *Product) CheckStock(id int) (int, error) {
	db := config.ConnectGormDB()
	var totalStocks int
	var query string = "SELECT stocks FROM products WHERE id = ?;"
	results := db.Raw(query, id).Scan(&totalStocks)
	if results.Error != nil {
		return 0, results.Error
	}
	return totalStocks, nil
}

func (p *Product) Count() (int, error) {
	db := config.ConnectGormDB()
	var totalProducts int
	var query string = "SELECT COUNT(*) AS total FROM products"
	results := db.Raw(query).Scan(&totalProducts)
	if results.Error != nil {
		return 0, results.Error
	}
	return totalProducts, nil
}

func (p *Product) UpdateStocks(id int, stocks int) (sql.Result, error) {
	// Gunakan mutex yang sama jika diperlukan
	mutex.Lock()
	defer mutex.Unlock()

	var db = config.ConnectSQLDB()
	defer db.Close()

	// Periksa apakah produk dengan ID yang diberikan ada dan belum dihapus
	var checkQuery = "SELECT id FROM products WHERE id = ? AND deleted_at IS NULL"
	err := db.QueryRow(checkQuery, id).Scan(new(int)) // Scan ke variabel baru untuk memeriksa keberadaan produk
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Produk tidak ditemukan atau sudah dihapus
			return nil, errors.New("product not found or deleted")
		}
		return nil, err
	}

	query := "UPDATE products SET stocks = ? WHERE id = ? AND deleted_at IS NULL;"

	result, err := db.Exec(query, stocks, id)
	if err != nil {
		return nil, err
	}

	// Periksa apakah result tidak nil
	if result == nil {
		return nil, errors.New("update operation failed")
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
