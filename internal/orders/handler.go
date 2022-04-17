package orders

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ordersHandler struct {
	repo *OrdersRepo
}

func (h ordersHandler) create(context *gin.Context) {
	orders := &api.Orders{}
	if err := context.Bind(orders); err != nil {
		context.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}
	order := h.repo.Create(ResponseToOrder(orders))

	context.JSON(http.StatusOK, order)
}

func (h ordersHandler) getAll(context *gin.Context) {
	userid := context.Param("userid")
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	orders, totalCount := h.repo.GetUsersOrders(userid, pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = OrdersToResponse(&orders)

	context.JSON(http.StatusOK, paginatedResponse)
}

func (h ordersHandler) confirm(context *gin.Context) {
	orderid := context.Param("orderid")
	order := h.repo.ConfirmOrder(orderid)
	context.JSON(http.StatusOK, order)
}

// Cancel order if its not past 14 days
func (h ordersHandler) cancel(context *gin.Context) {
	orderid := context.Param("orderid")
	order := h.repo.CancelOrder(orderid)
	context.JSON(http.StatusOK, order)
}

// Complete order if user confirms
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
