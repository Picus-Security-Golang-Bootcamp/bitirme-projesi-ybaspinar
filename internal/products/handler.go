package products

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productHandler struct {
	repo *ProductRepo
}

func (h *productHandler) create(context *gin.Context) {
	productb := &api.Products{}
	if err := context.Bind(productb); err != nil {
		context.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}
	product := h.repo.Create(ResponseToProduct(productb))

	context.JSON(http.StatusOK, product)
}

func (h *productHandler) delete(context *gin.Context) {
	id := context.Param("id")
	if err := h.repo.Delete(id); err != nil {
		context.JSON(httpErrors.ErrorResponse(err))
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (h *productHandler) search(context *gin.Context) {

}

func (h *productHandler) update(context *gin.Context) {
	id := context.Param("id")
	productb := &api.Products{ID: id}
	if err := context.Bind(productb); err != nil {
		context.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}
	product := h.repo.Update(ResponseToProduct(productb))

	context.JSON(http.StatusOK, product)
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
