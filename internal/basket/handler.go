package basket

import "github.com/gin-gonic/gin"

type basketHandler struct {
	repo *BasketRepo
}

//TODO: implement
func (h basketHandler) create(context *gin.Context) {

}

//TODO: implement
func (h basketHandler) list(context *gin.Context) {

}

//TODO: implement
func (h basketHandler) update(context *gin.Context) {

}

//TODO: implement delete
func (h basketHandler) delete(context *gin.Context) {

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
