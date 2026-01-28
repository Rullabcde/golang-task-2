package service

import (
	"context"
	"errors"

	"category-api/internal/entity"
	"category-api/internal/repository"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryService interface {
	GetAll(ctx context.Context) ([]entity.Category, error)
	GetByID(ctx context.Context, id int) (*entity.Category, error)
	Create(ctx context.Context, req entity.CreateCategoryRequest) (*entity.Category, error)
	Update(ctx context.Context, id int, req entity.UpdateCategoryRequest) (*entity.Category, error)
	Delete(ctx context.Context, id int) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll(ctx context.Context) ([]entity.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *categoryService) GetByID(ctx context.Context, id int) (*entity.Category, error) {
	cat, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrCategoryNotFound
	}
	return cat, nil
}

func (s *categoryService) Create(ctx context.Context, req entity.CreateCategoryRequest) (*entity.Category, error) {
	category := &entity.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) Update(ctx context.Context, id int, req entity.UpdateCategoryRequest) (*entity.Category, error) {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrCategoryNotFound
	}

	category := &entity.Category{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.Update(ctx, id, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) Delete(ctx context.Context, id int) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return ErrCategoryNotFound
	}

	return s.repo.Delete(ctx, id)
}
