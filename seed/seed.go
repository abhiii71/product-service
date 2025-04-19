package seed

import (
	"database/sql"
)

func SeedProducts(db *sql.DB) error {
	products := []struct {
		name  string
		price float64
	}{
		{"Laptop", 999.99},
		{"Phone", 499.99},
		{"Headphones", 79.99},
		{"Keyboard", 49.99},
		{"Mouse", 39.99},
	}

	for _, p := range products {
		_, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", p.name, p.price)
		if err != nil {
			return err
		}
	}
	return nil
}
