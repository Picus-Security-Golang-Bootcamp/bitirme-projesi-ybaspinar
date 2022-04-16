package categories

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoriesHandler struct {
	repo *CategoriesRepo
}

func (h categoriesHandler) create(context *gin.Context) {
}

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
