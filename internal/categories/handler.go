package categories

import "github.com/gin-gonic/gin"

type categoriesHandler struct {
	repo *CategoriesRepo
}

func (h categoriesHandler) create(context *gin.Context) {

}

func (h categoriesHandler) getAll(context *gin.Context) {

}

func NewCategoriesHandler(r *gin.RouterGroup, repo *CategoriesRepo) {
	h := &categoriesHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/", h.getAll)
}
