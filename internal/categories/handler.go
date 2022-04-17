package categories

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"net/http"
)

type categoriesHandler struct {
	repo *CategoriesRepo
}

// create new categories with given data
//TODO: add validation
func (h categoriesHandler) create(context *gin.Context) {
	var Categories models.Category
	if err := gocsv.Unmarshal(context.Request.Body, &Categories); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.CreateFromCSV(&Categories); err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, Categories)
}

// getAll categories
//TODO: add validation
func (h categoriesHandler) getAll(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	category, totalCount := h.repo.GetAll(pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = &category

	context.JSON(http.StatusOK, paginatedResponse)
}

func NewCategoriesHandler(r *gin.RouterGroup, repo *CategoriesRepo) {
	h := &categoriesHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/", h.getAll)
}
