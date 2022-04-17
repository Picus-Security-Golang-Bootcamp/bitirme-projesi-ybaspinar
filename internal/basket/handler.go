package basket

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type basketHandler struct {
	repo *BasketRepo
}

//TODO: authenication
func (h basketHandler) create(context *gin.Context) {
	var basket models.Basket
	if err := context.ShouldBindJSON(&basket); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(&basket); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, basket)
}

//TODO: authenticate
func (h basketHandler) list(context *gin.Context) {
	var basket models.Basket
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, totalCount := h.repo.GetAllByUserID(basket.ID.String(), pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = &products
	context.JSON(http.StatusOK, paginatedResponse)
}

//TODO: authenticate
func (h basketHandler) update(context *gin.Context) {
	var basket models.Basket
	if err := context.ShouldBindJSON(&basket); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Update(&basket); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, basket)
}

//TODO: authenticate
func (h basketHandler) delete(context *gin.Context) {
	var basket models.Basket
	if err := context.ShouldBindJSON(&basket); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Delete(&basket); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, basket)
}

func NewBasketHandler(r *gin.RouterGroup, repo *BasketRepo) {
	h := &basketHandler{
		repo: repo,
	}
	r.POST("/create", h.create)
	r.GET("/", h.list)
	r.POST("/update", h.update)
	r.DELETE("/", h.delete)

}
