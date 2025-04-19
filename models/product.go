package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func GetProductByID(db *sql.DB, rdb *redis.Client, id int, ttl time.Duration) (*Product, error) {
	cacheKey := fmt.Sprintf("product: %d,", id)
	val, err := rdb.Get(context.Background(), cacheKey).Result()

	if err == nil {
		var p Product
		json.Unmarshal([]byte(val), &p)
		fmt.Println("Cache hit")
		return &p, nil
	}

	fmt.Println("Cache miss")

	row := db.QueryRow("SELECT id, name, price FROM products WHERE id=$1", id)
	var p Product
	err = row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	b, _ := json.Marshal(p)
	rdb.Set(context.Background(), cacheKey, b, ttl)
	return &p, nil
}

// ListProducts retrieves all products
func ListProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// CreateProduct adds a new product to the database
func CreateProduct(db *sql.DB, name string, price float64) (int, error) {
	var id int
	err := db.QueryRow(
		"INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id",
		name, price,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateProduct modifies an existing product
func UpdateProduct(db *sql.DB, id int, name string, price float64) (bool, error) {
	res, err := db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", name, price, id)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := res.RowsAffected()
	return rowsAffected > 0, nil
}

// DeleteProduct removes a product by ID
func DeleteProduct(db *sql.DB, rdb *redis.Client, id int) (bool, error) {
	res, err := db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		return false, err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected > 0 {
		cacheKey := fmt.Sprintf("product:%d", id)
		rdb.Del(context.Background(), cacheKey)
		return true, nil
	}
	return false, nil
}
