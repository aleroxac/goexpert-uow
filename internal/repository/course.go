package repository

import (
	"context"
	"database/sql"

	"github.com/aleroxac/goexpert-uow/internal/db"
	"github.com/aleroxac/goexpert-uow/internal/entity"
)

type CourseRespositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRespository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRespository(dtb *sql.DB) *CourseRespository {
	return &CourseRespository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CourseRespository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name:       course.Name,
		CategoryID: int32(course.CategoryID),
	})
}
