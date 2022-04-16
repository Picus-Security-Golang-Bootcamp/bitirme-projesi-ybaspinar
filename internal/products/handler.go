package products

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productHandler struct {
	repo *ProductRepo
}

func (h *productHandler) create(context *gin.Context) {

}

func (h *productHandler) delete(context *gin.Context) {

}

func (h *productHandler) search(context *gin.Context) {

}

func (h *productHandler) update(context *gin.Context) {

}

func (h *productHandler) getAll(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, totalCount := h.repo.GetAll(pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = ProductsToResponseList(&products)

	context.JSON(http.StatusOK, paginatedResponse)

}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepo) {
	h := &productHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/", h.getAll)
	r.GET("/:id", h.search)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}
