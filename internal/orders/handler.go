package orders

import "github.com/gin-gonic/gin"

type ordersHandler struct {
	repo *OrdersRepo
}

func (h ordersHandler) create(context *gin.Context) {

}

func (h ordersHandler) getAll(context *gin.Context) {

}

func (h ordersHandler) delete(context *gin.Context) {

}

func (h ordersHandler) update(context *gin.Context) {

}

func (h ordersHandler) complete(context *gin.Context) {

}
func NewOrdersHandler(r *gin.RouterGroup, repo *OrdersRepo) {
	h := &ordersHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/", h.getAll)
	r.DELETE("/", h.delete)
	r.POST("/update", h.update)
	r.POST("/complete", h.complete)
}
