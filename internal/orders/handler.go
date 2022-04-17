package orders

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	_ "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ordersHandler struct {
	repo *OrdersRepo
}

//TODO: Add authentication
func (h ordersHandler) create(context *gin.Context) {
	var order models.Order
	if err := context.Bind(order); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.repo.Create(&order)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, order)
}

//TODO: Add authentication
func (h ordersHandler) getAll(context *gin.Context) {
	userid := context.Param("userid")
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	orders, totalCount := h.repo.GetUsersOrders(userid, pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = &orders
	context.JSON(http.StatusOK, paginatedResponse)
}

//TODO: Add authentication
func (h ordersHandler) confirm(context *gin.Context) {
	orderid := context.Param("orderid")
	order := h.repo.ConfirmOrder(orderid)
	context.JSON(http.StatusOK, order)
}

// Cancel order if its not past 14 days
//TODO: Add authentication
func (h ordersHandler) cancel(context *gin.Context) {
	orderid := context.Param("orderid")
	order := h.repo.CancelOrder(orderid)
	context.JSON(http.StatusOK, order)
}

// Complete order if user confirms
//TODO: Add authentication
func (h ordersHandler) complete(context *gin.Context) {
	orderid := context.Param("orderid")
	order := h.repo.CompleteOrder(orderid)
	context.JSON(http.StatusOK, order)
}
func NewOrdersHandler(r *gin.RouterGroup, repo *OrdersRepo) {
	h := &ordersHandler{repo: repo}

	r.GET("/", h.getAll)
	r.POST("/create", h.create)
	r.POST("/confirm", h.confirm)
	r.POST("/cancel", h.cancel)
	r.POST("/complete", h.complete)
}
