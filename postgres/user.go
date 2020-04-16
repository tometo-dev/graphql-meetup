package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/tsuki42/graphql-meetup/models"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Table("USER").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}
