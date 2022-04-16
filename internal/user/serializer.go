package user

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/api"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-ybaspinar/internal/models"
	"github.com/google/uuid"
)

func UserToResponse(user *models.User) *api.User {
	return &api.User{
		ID:        user.ID.String(),
		Firstname: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func ResponseToUser(user *api.User) *models.User {
	id, _ := uuid.Parse(user.ID)
	return &models.User{
		ID:        id,
		FirstName: user.Firstname,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
