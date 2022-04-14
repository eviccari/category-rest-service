package repositories

import (
	"category-rest-service/models"
	"category-rest-service/payloads"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBRepository -
type MongoDBRepository struct {
	ctx context.Context
	db  *mongo.Collection
}

// NewMongoDBRepository - Create new MongoDBRepository instace
func NewMongoDBRepository(context context.Context, db *mongo.Collection) (repo *MongoDBRepository) {
	return &MongoDBRepository{ctx: context, db: db}
}

// FindByAttributeFilters - Find categories searching by filter 
func (repo *MongoDBRepository) FindByAttributeFilters(filters interface{}) (categories []models.Category, err error){
	cursor, err := repo.db.Find(repo.ctx, filters)
	if err != nil {
		return nil, err
	}

	var cursorDetails []bson.M
	if err := cursor.All(repo.ctx, &cursorDetails); err != nil {
		return nil, err
	}

	var result []models.Category
	for _, detail := range cursorDetails {
		category, err := convert(detail)
		if err !=  nil {
			return nil, err
		}
		result = append(result, category)
	}
	return result, nil
}

// Create - Create new category into database
func (repo *MongoDBRepository) Create(category models.Category) (created models.Category, err error) {
	_, err = repo.db.InsertOne(repo.ctx, category)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

// Delete - Delete category that corresponds with the given id
func (repo *MongoDBRepository) Delete(id string) (err error) {
	filter := bson.D{{"ID", id}}
	_, err = repo.db.DeleteOne(repo.ctx, filter)

	if err != nil {
		return err
	}
	return nil
}

// Override - Override category with new data that have been sent by client
func (repo *MongoDBRepository) Override(category models.Category) (updated models.Category, err error) {
	filter := payloads.AttributeFilters{ID: category.ID}
	filterByID := filter.Clean()
	result, err := repo.FindByAttributeFilters(filterByID) // remove unnecessary fields  

	if err != nil {
		return models.Category{}, err
	}
	if len(result) == 0 {
		return models.Category{}, errors.New("not found")
	}

	actual := result[0]

	actual.UpdatedAt = time.Now()
	actual.Name = category.Name

	updateBody := bson.D{{"$set", bson.D{{"name", actual.Name}, {"updatedAt", actual.UpdatedAt}}}}

	_, err = repo.db.UpdateMany(repo.ctx, filterByID, updateBody)
	if err != nil {
		return actual, err
	}

	return actual, nil
}

// convert - Convert bson.M interface into models.Category structure 
func convert(b bson.M) (category models.Category, err error) {
	var c models.Category

	bsonBytes, err := bson.Marshal(b)
	if err != nil {
		return models.Category{}, err
	}
	_ = bson.Unmarshal(bsonBytes, &c)
	return c, nil
}
