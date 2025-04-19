package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"database/sql"
	"product-service/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
)

type App struct {
	DB  *sql.DB
	RDB *redis.Client
	TTL time.Duration
}

func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := models.GetProductByID(a.DB, a.RDB, id, a.TTL)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (a *App) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.ListProducts(a.DB)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := models.CreateProduct(a.DB, p.Name, p.Price)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var p models.Product
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updated, err := models.UpdateProduct(a.DB, id, p.Name, p.Price)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	if !updated {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	deleted, err := models.DeleteProduct(a.DB, a.RDB, id)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	if !deleted {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

