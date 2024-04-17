package usecase

import (
	"context"

	"github.com/aleroxac/goexpert-uow/internal/entity"
	"github.com/aleroxac/goexpert-uow/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCase struct {
	CourseRepository    repository.CourseRespositoryInterface
	CategoryRespository repository.CategoryRespositoryInterface
}

func NewAddCourseUseCase(courseRespository repository.CourseRespositoryInterface, categoryRepository repository.CategoryRespositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{
		CourseRepository:    courseRespository,
		CategoryRespository: categoryRepository,
	}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := a.CategoryRespository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
