package orders

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	jwtHelper "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/JWT"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	_ "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ordersHandler struct {
	repo *OrdersRepo
	cfg  *config.Config
}

func (h ordersHandler) create(context *gin.Context) {
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	var order models.Order
	if err := context.Bind(order); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if decodedClaims.UserID == order.UserID {
		err := h.repo.Create(&order)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, order)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func (h ordersHandler) getAll(context *gin.Context) {
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	orders, totalCount := h.repo.GetUsersOrders(decodedClaims.UserID.String(), pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = &orders
	context.JSON(http.StatusOK, paginatedResponse)
}

// Cancel order if its not past 14 days
func (h ordersHandler) cancel(context *gin.Context) {
	var order models.Order
	if error := context.Bind(&order); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	h.repo.CancelOrder(order.ID.String(), decodedClaims.UserID.String())
	context.JSON(http.StatusOK, order)
}

// Complete order if user confirms
func (h ordersHandler) complete(context *gin.Context) {
	var order models.Order
	if error := context.Bind(&order); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	h.repo.CompleteOrder(order.ID.String(), decodedClaims.UserID.String())
	context.JSON(http.StatusOK, order)
}
func NewOrdersHandler(r *gin.RouterGroup, repo *OrdersRepo, cfg *config.Config) {
	h := &ordersHandler{repo: repo, cfg: cfg}

	r.GET("/", h.getAll)
	r.POST("/create", h.create)
	r.POST("/cancel", h.cancel)
	r.POST("/complete", h.complete)
}
