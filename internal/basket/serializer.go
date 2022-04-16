package basket

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"github.com/google/uuid"
)

func BasketToResponca(basket *models.Basket) *api.Cart {
	return &api.Cart{
		ID:     basket.ID.String(),
		Userid: basket.UserID.String(),
	}
}
func ResponseToBasket(basket *api.Cart) *models.Basket {
	id, _ := uuid.Parse(basket.ID)
	userid, _ := uuid.Parse(basket.Userid)
	return &models.Basket{
		ID:     id,
		UserID: userid,
	}
}
