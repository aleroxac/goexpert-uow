package repository

import (
	"context"
	"database/sql"

	"github.com/aleroxac/goexpert-uow/internal/db"
	"github.com/aleroxac/goexpert-uow/internal/entity"
)

type CategoryRespositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRespository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dtb *sql.DB) *CategoryRespository {
	return &CategoryRespository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CategoryRespository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
