package models

import (
	"time"

	"github.com/google/uuid"
)

// Category - Describe Category data structure
type Category struct {
	ID string           `bson:"ID"`
	Name string         `bson:"name"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

// NewCategory - create new instace of Category
func NewCategory(name string) (category Category) {
	id := uuid.New().String()
	updatedAt := time.Now()
	return Category{
		ID: id, 
		Name: name, 
		CreatedAt: updatedAt,
		UpdatedAt: updatedAt,
	}
}

// CategoryRepository - Describe interface that must be implemented to handle persistence methods for Category
type CategoryRepository interface {
	FindByAttributeFilters(filters interface{}) (category []Category, err error)
	Create(category Category) (created Category, err error)
	Delete(id string) (err error)
	Override(category Category) (updated Category, err error)
}