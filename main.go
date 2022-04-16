package main

import (
	"category-rest-service/database"
	"category-rest-service/errors"
	"category-rest-service/repositories"
	"category-rest-service/routes/v1"
	"category-rest-service/services"
	"category-rest-service/utils"
	"context"

	"github.com/gin-gonic/gin"

	_ "category-rest-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Category API documentation
// @version 0.1.0
// @host localhost:8080
// @BasePath /categories
func main() {
	var ctx = context.TODO()

	db, err := database.NewProvider().MongoDB()
	if err != nil {
		utils.PanicIfError(err)
	}

	app := gin.Default()

	repo := repositories.NewMongoDBRepository(ctx, db)
	service := services.NewCategoryService(repo)
	httpErrorHandler := errors.GinHTTPErrorHandler{}
	router := routes.NewCategoryRouter(service, &httpErrorHandler)

	appGroup := app.Group("/categories/v1")
	{
		appGroup.GET("/home", router.Home)
		appGroup.GET("/", router.FindByAttributeFilters)
		appGroup.POST("/", router.Create)
		appGroup.PATCH("/", router.Patch)
		appGroup.DELETE("/", router.Delete)
		appGroup.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	utils.PanicIfError(app.Run("0.0.0.0:8080"))
}
