package categories

import (
	"encoding/csv"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	jwtHelper "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/JWT"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoriesHandler struct {
	repo *CategoriesRepo
	cfg  *config.Config
}

// create new categories with given data
func (h categoriesHandler) create(context *gin.Context) {
	var categories []models.Category
	token := context.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, h.cfg.JWTConfig.SecretKey)
	if decodedClaims.IsAdmin {
		file, _, err := context.Request.FormFile("file")
		defer file.Close()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		csvLines, err := csv.NewReader(file).ReadAll()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, line := range csvLines {
			categories = append(categories, models.Category{
				ID:   line[0],
				Name: line[1],
			})
		}
		categories = categories[1:]
		for _, category := range categories {
			h.repo.Create(&category)
		}
		context.JSON(http.StatusOK, categories)
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
