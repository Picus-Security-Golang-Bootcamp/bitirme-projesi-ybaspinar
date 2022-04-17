package categories

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"github.com/google/uuid"
)

func CategoryToResponse(category *models.Category) *api.Category {
	return &api.Category{
		ID: category.ID.String(),
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
	id, _ := uuid.Parse(category.ID)
	return &[]models.Category{
		{
			ID: id,
			//Name:     category.Name,
			//Products: category.Products,
		},
	}
}
