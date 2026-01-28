package repository

import (
	"context"

	"category-api/config"
	"category-api/entity"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]entity.ProductWithCategory, error)
	GetByID(ctx context.Context, id int) (*entity.ProductWithCategory, error)
	Create(ctx context.Context, product *entity.Product) error
	Update(ctx context.Context, id int, product *entity.Product) error
	Delete(ctx context.Context, id int) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetAll(ctx context.Context) ([]entity.ProductWithCategory, error) {
	query := `
		SELECT p.id, p.name, p.price, p.category_id, c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		ORDER BY p.id
	`
	rows, err := config.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.ProductWithCategory
	for rows.Next() {
		var p entity.ProductWithCategory
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID, &p.CategoryName); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepository) GetByID(ctx context.Context, id int) (*entity.ProductWithCategory, error) {
	query := `
		SELECT p.id, p.name, p.price, p.category_id, c.name as category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1
	`
	var p entity.ProductWithCategory
	err := config.DB.QueryRow(ctx, query, id).
		Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID, &p.CategoryName)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) error {
	err := config.DB.QueryRow(ctx,
		"INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3) RETURNING id",
		product.Name, product.Price, product.CategoryID).Scan(&product.ID)
	return err
}

func (r *productRepository) Update(ctx context.Context, id int, product *entity.Product) error {
	_, err := config.DB.Exec(ctx,
		"UPDATE products SET name = $1, price = $2, category_id = $3 WHERE id = $4",
		product.Name, product.Price, product.CategoryID, id)
	return err
}

func (r *productRepository) Delete(ctx context.Context, id int) error {
	_, err := config.DB.Exec(ctx, "DELETE FROM products WHERE id = $1", id)
	return err
}
