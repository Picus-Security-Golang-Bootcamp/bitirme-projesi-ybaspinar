package basket

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	jwtHelper "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/JWT"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type basketHandler struct {
	repo *BasketRepo
	cfg  *config.Config
}

//create Creates new basket
func (h basketHandler) create(context *gin.Context) {
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	var basket models.Basket
	if err := context.ShouldBindJSON(&basket); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if basket.UserID == decodedClaims.UserID {
		if err := h.repo.Create(&basket); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, basket)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

}

//getAll Gets all baskets with given ID
func (h basketHandler) getAll(context *gin.Context) {
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)

	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, totalCount := h.repo.GetAllByUserID(decodedClaims.UserID.String(), pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = &products
	context.JSON(http.StatusOK, paginatedResponse)

}

//update Updates basket with given ID
func (h basketHandler) update(context *gin.Context) {
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	var basket models.Basket
	if err := context.ShouldBindJSON(&basket); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if basket.UserID == decodedClaims.UserID {
		if err := h.repo.Update(&basket); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, basket)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

}

//delete Deletes basket with given ID
func (h basketHandler) delete(context *gin.Context) {
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	var basket models.Basket
	if err := context.ShouldBindJSON(&basket); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if basket.UserID == decodedClaims.UserID {
		if err := h.repo.Delete(&basket); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, basket)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}

}

func NewBasketHandler(r *gin.RouterGroup, repo *BasketRepo, cfg *config.Config) {
	h := &basketHandler{repo: repo, cfg: cfg}
	r.POST("/create", h.create)
	r.GET("/", h.getAll)
	r.POST("/update", h.update)
	r.DELETE("/", h.delete)

}
