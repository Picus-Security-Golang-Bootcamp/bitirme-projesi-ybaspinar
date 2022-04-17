package categories

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoriesHandler struct {
	repo *CategoriesRepo
}

// create new categories with given data
func (h categoriesHandler) create(context *gin.Context) {
	categories := &api.Category{}
	if err := context.Bind(categories); err != nil {
		context.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}
	cCategories := ResponseToCategory(categories)
	h.repo.CreateFromCSV(cCategories)

	context.JSON(http.StatusCreated, categories)

}

// getAll categories
func (h categoriesHandler) getAll(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	category, totalCount := h.repo.GetAll(pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = CategoriesToResponse(&category)

	context.JSON(http.StatusOK, paginatedResponse)
}

func NewCategoriesHandler(r *gin.RouterGroup, repo *CategoriesRepo) {
	h := &categoriesHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/", h.getAll)
}
