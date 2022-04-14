package services

import (
	"category-rest-service/models"
)

// CategoryService - Describe CategoryService structure
type CategoryService struct {
	repo models.CategoryRepository
}

// NewCategoryService - Create new CategoryService instance
func NewCategoryService(repo models.CategoryRepository) (service *CategoryService) {
	return &CategoryService{
		repo: repo,
	}
}

// FindByAttributeFilters - Find categories searching by filter that have ben sent by clients
func (service *CategoryService) FindByAttributeFilters(filters interface{}) (category []models.Category, err error) {
	categories, err := service.repo.FindByAttributeFilters(filters)
	if err != nil {
		return []models.Category{}, err
	}
	return categories, nil
}

// Create - Create new category into database
func (service *CategoryService) Create(name string) (created models.Category, err error) {
	category := models.NewCategory(name)
	c, err := service.repo.Create(category)

	if err != nil {
		return models.Category{}, err
	}
	return c, nil
}

// Delete - Delete category that corresponds with the given id
func (service *CategoryService) Delete(id string) (err error) {
	err = service.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Override - Override category with new data that have been sent by client
func (service *CategoryService) Override(category models.Category) (updated models.Category, err error) {
	c, err := service.repo.Override(category)
	if err != nil {
		return models.Category{}, err
	}
	return c, nil
}
