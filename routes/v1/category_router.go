package routes

import (
	"category-rest-service/errors"
	"category-rest-service/models"
	"category-rest-service/payloads"
	"category-rest-service/services"
	"category-rest-service/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoryRouter - Describe CategoryRouter structure
type CategoryRouter struct {
	service *services.CategoryService
}

// NewCategoryRouter - Create new CategoryRouter instance
func NewCategoryRouter(service *services.CategoryService, httpErrorHandler errors.HTTPErrorHandler) (categoryRouter *CategoryRouter){
	return &CategoryRouter{service: service}
}

// Home - Home router to return simple message
// Home ... Show welcome message
// @Summary Show welcome message
// @Description Show welcome message
// @Tags Categories Home
// @Success 200 {object} models.Category
// @Router /v1/home [get]
func (categoryRouter *CategoryRouter) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hey!"})
}

// FindByAttributeFilters - Find categories searching by filter that have ben sent by clients
// GetCategoriesByAttributeFilters ... Get categories searching by supported params
// @Summary Get list of categories
// @Description get categories by params
// @Tags Categories
// @Param ID query string false "Category ID"
// @Param name query string false "Category Name"
// @Success 200 {object} models.Category
// @Failure 400,404,500 {object} errors.HTTPError
// @Router /v1 [get]
func (categoryRouter *CategoryRouter) FindByAttributeFilters(c *gin.Context) {
	errorHandler := errors.NewGinHTTPErrorHandler(c)
	filters := payloads.AttributeFilters{}
	
	if err := c.ShouldBindQuery(&filters); err != nil{
		errorHandler.Response(http.StatusBadRequest, "invalid parameters")
	} 

	categories, err := categoryRouter.service.FindByAttributeFilters(filters.Clean())
	if err != nil {
		errorHandler.Response(http.StatusInternalServerError, err.Error())
		return
 	}
	if len(categories) == 0 {
		errorHandler.Response(http.StatusNotFound, "not found")
		return
	}

	c.IndentedJSON(http.StatusOK, categories)
}

// Patch - Update category by ID 
// Patch ... Update category by ID
// @Summary Update category  using given ID
// @Description Update category
// @Tags Categories
// @Accept json
// @Param ID query string true "ID of Category"
// @Param category body payloads.Category true "Category Data"
// @Success 200 {object} models.Category
// @Failure 400,422,500 {object} errors.HTTPError
// @Router /v1 [patch]
func (categoryRouter *CategoryRouter) Patch(c *gin.Context) {
	errorHandler := errors.NewGinHTTPErrorHandler(c)
	payload := payloads.Category{}
	attributeFilters := payloads.AttributeFilters{}
	
	if err := c.ShouldBindJSON(&payload); err != nil {
		errorHandler.Response(http.StatusBadRequest, err.Error())
		return 
	}
	if err := c.ShouldBindQuery(&attributeFilters); err != nil {
		errorHandler.Response(http.StatusBadRequest, err.Error())
		return
	}

    toUpdate := models.Category{
		ID: attributeFilters.ID, 
		Name: payload.Name,
	}

	updated, err := categoryRouter.service.Override(toUpdate)
	if err != nil {
		errorHandler.Response(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, updated)
}

// Create - Create new category into repository
// Create ... Create new category into repository
// @Summary Create new category using sent name 
// @Description Create new category
// @Tags Categories
// @Accept json
// @Param user body models.Category true "Category Data"
// @Success 200 {object} models.Category
// @Failure 400,422,500 {object} errors.HTTPError
// @Router / [post]
func (categoryRouter *CategoryRouter) Create(c *gin.Context) {
	errorHandler := errors.NewGinHTTPErrorHandler(c)
	newCategory := payloads.Category{}
	err := c.ShouldBindJSON(&newCategory)

	if err !=nil {
		errorHandler.Response(http.StatusBadRequest, err.Error())
		return
	}

	created, err := categoryRouter.service.Create(newCategory.Name)
	if err != nil {
		errorHandler.Response(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, created)
}

// Delete - Delete category from repository by ID
// Delete ... Delete category from repository
// @Summary Delete category using given ID 
// @Description Delete category
// @Tags Categories
// @Param ID query string true "Category ID"
// @Success 200 {object} object
// @Failure 400,422,500 {object} errors.HTTPError
// @Router /v1 [delete]
func (categoryRouter *CategoryRouter) Delete(c *gin.Context){
	errorHandler := errors.NewGinHTTPErrorHandler(c)
	filters := payloads.AttributeFilters{}
	if err := c.ShouldBindQuery(&filters); err != nil {
		errorHandler.Response(http.StatusBadRequest, err.Error())
		return
	}
	if utils.IsEmpty(filters.ID) {
		errorHandler.Response(http.StatusBadRequest, "parameter ID  is required")
		return
	}

	err := categoryRouter.service.Delete(filters.ID)
	if err != nil {
		errorHandler.Response(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, fmt.Sprintf("entity with id %v deleted", filters.ID))
}



