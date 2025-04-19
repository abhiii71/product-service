package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"

	"product-service/handlers"
	"product-service/seed"
)

func main() {
	// Load env variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	ttl := 60 * time.Second

	// Connect to Postgres
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Ping DB to check connection
	if err = db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	defer rdb.Close()

	// Ping Redis
	if err = rdb.Ping(rdb.Context()).Err(); err != nil {
		log.Fatal("Redis unreachable:", err)
	}

	// Seed products (if needed)
	err = seed.SeedProducts(db)
	if err != nil {
		log.Fatal("Seeding products failed:", err)
	}

	// Set up router and handlers
	app := &handlers.App{
		DB:  db,
		RDB: rdb,
		TTL: ttl,
	}
	r := chi.NewRouter()

	r.Get("/products/{id}", app.GetProduct)
	r.Get("/products", app.ListProducts)
	r.Post("/products", app.CreateProduct)
	r.Put("/products/{id}", app.UpdateProduct)
	r.Delete("/products/{id}", app.DeleteProduct)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, r)
}

