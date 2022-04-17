package products

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
)

func ProductToResponse(p *models.Product) *api.Products {

	return &api.Products{
		ID:          p.ID.String(),
		Name:        p.Name,
		Price:       int64(p.Price),
		Stock:       int64(p.Stock),
		Sku:         p.SKU,
		Description: p.Description,
	}
}

func ProductsToResponseList(products *[]models.Product) []*api.Products {
	var productList []*api.Products
	for _, p := range *products {
		productList = append(productList, ProductToResponse(&p))
	}
	return productList
}

func ResponseToProduct(p *api.Products) *models.Product {
	return &models.Product{
		Name:        p.Name,
		Price:       float64(p.Price),
		Stock:       int(p.Stock),
		SKU:         p.Sku,
		Description: p.Description,
	}
}
