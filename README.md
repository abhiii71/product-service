#  ProductService — Go REST Microservice

A simple, Dockerized RESTful microservice built with Golang, PostgreSQL, and Redis caching to manage product data via CRUD operations.

---

##  Overview

**ProductService** allows clients to **Create**, **Read**, **Update**, and **Delete** product records stored in a **PostgreSQL database**.  
To enhance performance, the service uses **Redis** as a caching layer for individual product retrievals.

---

##  Features

✅ CRUD Product Management via HTTP API  
✅ Redis caching for fetching products by ID  
✅ PostgreSQL as persistent data store  
✅ Dockerized environment with Docker Compose  
✅ Seeded with sample product data  
✅ Clean error handling and status codes  

---

##  Functional Requirements

- **Database:**  
  - PostgreSQL  
  - `products` table:

    | Field | Type  |
    |:-------|:--------|
    | id      | SERIAL PRIMARY KEY |
    | name    | TEXT    |
    | price   | REAL    |

  - Seeded with at least **5 sample products** via `init.sql`

- **Redis Caching:**  
  - Before querying the database, check if the product exists in Redis  
  - If cached, return from Redis  
  - If not cached:
    - Fetch from database  
    - Cache in Redis with a TTL of **1 minute**  
    - Return to client  

- **HTTP API Endpoints:**

    | Method | Endpoint          | Description                 | Caching |
    |:--------|:----------------|:-----------------------------|:----------|
    | GET     | `/products/{id}`  | Fetch product by ID            | ✅ |
    | GET     | `/products`       | Fetch all products             | ❌ |
    | POST    | `/products`       | Create new product             | ❌ |
    | PUT     | `/products/{id}`  | Update existing product        | ❌ |
    | DELETE  | `/products/{id}`  | Delete product by ID           | ❌ |

---

##  How to Run 


### 1. Clone Repository

```bash
git clone https://github.com/abhiii71/product-service.git
cd product-service
```
### 2. Run with Docker Compose
```bash 
docker-compose up --build
```

The ProductService will be available at:
```bash
http://localhost:8080
```

####  API Usage Examples
Create Product:
```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name": "New Product", "price": 99.99}'
```

Get Product by ID:
```bash
curl http://localhost:8080/products/1
```

Update Product:
```bash
curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{"name": "Updated Product", "price": 150.00}'
```

Delete Product:
```bash
curl -X DELETE http://localhost:8080/products/1
```

Get All Products:
```bash
curl http://localhost:8080/products
```
####  Tech Stack
```bash
Go 1.20+

PostgreSQL

Redis

Docker / Docker Compose

Gorilla Mux
```
