CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
);

INSERT INTO categories (name, description) VALUES
    ('Technology', 'Tech'),
    ('Food', 'Food'),
    ('Fashion', 'Clothing');

INSERT INTO products (name, price, category_id) VALUES
    ('iPhone 15', 999.99, 1),
    ('MacBook Pro', 2499.99, 1),
    ('Nasi Goreng', 25000.00, 2),
    ('T-Shirt', 150000.00, 3);
