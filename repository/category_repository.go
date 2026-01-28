package repository

import (
	"context"

	"category-api/config"
	"category-api/entity"
)

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]entity.Category, error)
	GetByID(ctx context.Context, id int) (*entity.Category, error)
	Create(ctx context.Context, category *entity.Category) error
	Update(ctx context.Context, id int, category *entity.Category) error
	Delete(ctx context.Context, id int) error
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]entity.Category, error) {
	rows, err := config.DB.Query(ctx, "SELECT id, name, description FROM categories ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var cat entity.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.Description); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id int) (*entity.Category, error) {
	var cat entity.Category
	err := config.DB.QueryRow(ctx, "SELECT id, name, description FROM categories WHERE id = $1", id).
		Scan(&cat.ID, &cat.Name, &cat.Description)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *categoryRepository) Create(ctx context.Context, category *entity.Category) error {
	err := config.DB.QueryRow(ctx,
		"INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id",
		category.Name, category.Description).Scan(&category.ID)
	return err
}

func (r *categoryRepository) Update(ctx context.Context, id int, category *entity.Category) error {
	_, err := config.DB.Exec(ctx,
		"UPDATE categories SET name = $1, description = $2 WHERE id = $3",
		category.Name, category.Description, id)
	return err
}

func (r *categoryRepository) Delete(ctx context.Context, id int) error {
	_, err := config.DB.Exec(ctx, "DELETE FROM categories WHERE id = $1", id)
	return err
}
