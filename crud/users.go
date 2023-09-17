package crud

import (
	"github.com/DavidEsdrs/template-server/connections"
	"github.com/DavidEsdrs/template-server/models"
)

func GetUser(id int) (*models.User, error) {
	db, err := connections.Connect()
	if err != nil {
		return &models.User{}, err
	}
	var user models.User
	result := db.First(&user, id)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers() ([]models.User, error) {
	db, err := connections.Connect()
	if err != nil {
		return []models.User{}, err
	}
	var users []models.User
	result := db.Limit(10).Find(&users)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
