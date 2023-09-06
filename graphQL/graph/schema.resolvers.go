package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/luizeduu/graphQl/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	coures, err := r.CourseDB.FindByCategoryId(obj.ID)

	if err != nil {
		return nil, err
	}

	var coursesModel []*model.Course

	for _, course := range coures {
		coursesModel = append(coursesModel, &model.Course{ID: course.ID, Name: course.Name, Description: &course.Description})
	}

	return coursesModel, nil
}

// Category is the resolver for the category field.
func (r *courseResolver) Category(ctx context.Context, obj *model.Course) (*model.Category, error) {
	category, err := r.CategoryDB.FindByCourseId(obj.ID)

	if err != nil {
		return nil, err
	}

	return &model.Category{ID: category.ID, Name: category.Name, Description: &category.Description}, nil
}

// AddCategory is the resolver for the addCategory field.
func (r *mutationResolver) AddCategory(ctx context.Context, category model.NewCategory) (*model.Category, error) {
	categoryDb, err := r.CategoryDB.Create(category.Name, *category.Description)

	if err != nil {
		return nil, err
	}

	return &model.Category{ID: categoryDb.ID, Name: categoryDb.Name, Description: &categoryDb.Description}, nil
}

// AddCourse is the resolver for the addCourse field.
func (r *mutationResolver) AddCourse(ctx context.Context, course model.NewCourse) (*model.Course, error) {
	courseDB, err := r.CourseDB.Create(course.Name, *course.Description, course.CategoryID)

	if err != nil {
		return nil, err
	}

	return &model.Course{ID: courseDB.ID, Name: courseDB.Name, Description: &courseDB.Description}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.FindAll()

	if err != nil {
		return nil, err
	}

	var categoriesModel []*model.Category

	for _, category := range categories {
		categoriesModel = append(categoriesModel, &model.Category{ID: category.ID, Name: category.Name, Description: &category.Description})
	}

	return categoriesModel, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindAll()

	if err != nil {
		return nil, err
	}

	var coursesModel []*model.Course

	for _, course := range courses {
		coursesModel = append(coursesModel, &model.Course{ID: course.ID, Name: course.Name, Description: &course.Description})
	}

	return coursesModel, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Course returns CourseResolver implementation.
func (r *Resolver) Course() CourseResolver { return &courseResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }