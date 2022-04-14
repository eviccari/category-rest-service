package database

import (
	"category-rest-service/utils"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Provider - Describe Provider structure
type Provider struct {}

// NewProvider - Create new Provider instance
func NewProvider() (provider *Provider) {
	return &Provider{}
}

// MongoDB - Get database instance for mongodb
func (p Provider) MongoDB() (db *mongo.Collection, err error) {
	var ctx = context.TODO()

	if err := godotenv.Load(); err != nil {
		utils.PanicIfError(err)
	}

	credential := options.Credential{
		Username: os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASS"),
	}

	uri := fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGODB_HOST"), os.Getenv("MONGODB_PORT"))
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)
	client, openError := mongo.Connect(ctx, clientOptions)
	utils.PanicIfError(openError)

	collection := client.Database("repositories").Collection("categories")
	return collection, nil
}