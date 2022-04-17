package categories

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	jwtHelper "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/JWT"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"net/http"
)

type categoriesHandler struct {
	repo *CategoriesRepo
	cfg  *config.Config
}

// create new categories with given data
func (h categoriesHandler) create(context *gin.Context) {
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	if decodedClaims.IsAdmin {
		var Categories models.Category
		if err := gocsv.Unmarshal(context.Request.Body, &Categories); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := h.repo.CreateFromCSV(&Categories); err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, Categories)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
	}

}

// getAll categories
func (h categoriesHandler) getAll(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	category, totalCount := h.repo.GetAll(pageIndex, pageSize)
	paginatedResponse := pagination.NewFromGinRequest(context, totalCount)
	paginatedResponse.Items = &category

	context.JSON(http.StatusOK, paginatedResponse)
}

func NewCategoriesHandler(r *gin.RouterGroup, repo *CategoriesRepo, cfg *config.Config) {
	h := &categoriesHandler{repo: repo, cfg: cfg}

	r.POST("/create", h.create)
	r.GET("/", h.getAll)
}
