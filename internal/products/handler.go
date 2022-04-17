package products

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	jwtHelper "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/JWT"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	_ "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productHandler struct {
	repo *ProductRepo
	cfg  *config.Config
}

func (h *productHandler) create(context *gin.Context) {
	var product models.Product
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	if decodedClaims.IsAdmin {
		if err := context.ShouldBindJSON(&product); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := h.repo.Create(&product); err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, product)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func (h *productHandler) delete(context *gin.Context) {
	id := context.Param("id")
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	if decodedClaims.IsAdmin {
		if err := h.repo.Delete(id); err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func (h *productHandler) search(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, totalCount := h.repo.FuzzySearchSkuAndNameAndId(context.Param("q"), pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = products
	context.JSON(http.StatusOK, products)
}

func (h *productHandler) update(context *gin.Context) {
	var product models.Product
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	if decodedClaims.IsAdmin {
		if err := context.ShouldBindJSON(&product); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := h.repo.Update(&product); err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, product)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

}

func (h *productHandler) getAll(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, totalCount := h.repo.GetAll(pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = &products
	context.JSON(http.StatusOK, paginatedResponse)
}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepo, cfg *config.Config) {
	h := &productHandler{repo: repo, cfg: cfg}

	r.POST("/create", h.create)
	r.GET("/", h.getAll)
	r.GET("/search/", h.search)
	r.PUT("/", h.update)
	r.DELETE("/:id", h.delete)
}
