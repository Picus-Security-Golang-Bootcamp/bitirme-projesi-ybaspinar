package categories

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
)

func CategoryToResponse(category *models.Category) *api.Category {
	return &api.Category{
		ID: int64(category.ID),
		//Name:     category.Name,
		//Products: category.Products,
	}
}

func CategoriesToResponse(categories *[]models.Category) []*api.Category {
	var response []*api.Category
	for _, category := range *categories {
		response = append(response, CategoryToResponse(&category))
	}
	return response
}

func ResponseToCategory(category *api.Category) *[]models.Category {
	return &[]models.Category{
		{
			//ID: int(category.ID),
			//Name:     category.Name,
			//Products: category.Products,
		},
	}
}
