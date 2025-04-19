CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

INSERT INTO products (name, price) VALUES
('Product 1', 10.99),
('Product 2', 20.99),
('Product 3', 30.99),
('Product 4', 40.99),
('Product 5', 50.99);

