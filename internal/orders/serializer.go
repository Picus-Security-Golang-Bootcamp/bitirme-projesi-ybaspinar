package orders

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"github.com/google/uuid"
)

func OrderToResponse(order *models.Order) *api.Orders {
	return &api.Orders{
		ID:     order.ID.String(),
		Userid: order.UserID.String(),
		Status: order.Status,
		//Cart:       order.BasketID.ID,
	}
}

func OrdersToResponse(orders *[]models.Order) []*api.Orders {
	var response []*api.Orders
	for _, order := range *orders {
		response = append(response, OrderToResponse(&order))
	}
	return response
}
func ResponseToOrder(order *api.Orders) *models.Order {
	id, _ := uuid.Parse(order.ID)
	//uid, _ := uuid.Parse(order.Userid)
	return &models.Order{
		ID: id,
		//UserID:        uid,
		Status: order.Status,
		//BasketID:      order.Cart,
		//UserAddressID: order.Useraddres,
	}
}
