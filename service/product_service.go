package service

import (
	"context"
	"errors"

	"category-api/internal/entity"
	"category-api/internal/repository"
)

var ErrProductNotFound = errors.New("product not found")

type ProductService interface {
	GetAll(ctx context.Context) ([]entity.ProductWithCategory, error)
	GetByID(ctx context.Context, id int) (*entity.ProductWithCategory, error)
	Create(ctx context.Context, req entity.CreateProductRequest) (*entity.Product, error)
	Update(ctx context.Context, id int, req entity.UpdateProductRequest) (*entity.Product, error)
	Delete(ctx context.Context, id int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll(ctx context.Context) ([]entity.ProductWithCategory, error) {
	return s.repo.GetAll(ctx)
}

func (s *productService) GetByID(ctx context.Context, id int) (*entity.ProductWithCategory, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

func (s *productService) Create(ctx context.Context, req entity.CreateProductRequest) (*entity.Product, error) {
	product := &entity.Product{
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) Update(ctx context.Context, id int, req entity.UpdateProductRequest) (*entity.Product, error) {
	// Check if exists
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrProductNotFound
	}

	product := &entity.Product{
		ID:         id,
		Name:       req.Name,
		Price:      req.Price,
		CategoryID: req.CategoryID,
	}

	if err := s.repo.Update(ctx, id, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) Delete(ctx context.Context, id int) error {
	// Check if exists
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return ErrProductNotFound
	}

	return s.repo.Delete(ctx, id)
}
