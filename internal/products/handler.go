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

//create Creates new product if user is admin
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

//delete Deletes product if user is admin
func (h *productHandler) delete(context *gin.Context) {
	var product models.Product
	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	if decodedClaims.IsAdmin {
		if err := h.repo.Delete(product.ID.String()); err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

//search Searches products by given query
func (h *productHandler) search(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, totalCount := h.repo.FuzzySearchSkuAndNameAndId(context.Param("q"), pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = products
	context.JSON(http.StatusOK, products)
}

//update Updates product if user is admin
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

//getAll Gets all products
func (h *productHandler) getAll(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, totalCount := h.repo.GetAll(pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = &products
	context.JSON(http.StatusOK, paginatedResponse)
}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepo, cfg *config.Config) {
	h := &productHandler{repo: repo, cfg: cfg}
	r.GET("/", h.getAll)
	r.PUT("/", h.update)
	r.DELETE("/", h.delete)
	r.POST("/create", h.create)
	r.GET("/search", h.search)

}
