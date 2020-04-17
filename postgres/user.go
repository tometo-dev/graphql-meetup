package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tsuki42/graphql-meetup/models"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) GetUserByField(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Table("USER").Where(fmt.Sprintf("%v = ?", field), value).Find(&user).Error
	return &user, err
}

func (u *UserRepo) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}

func (u *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	return u.GetUserByField("username", username)
}

func (u *UserRepo) CreateUser(tx *gorm.DB, user *models.User) (*models.User, error) {
	user.ID = uuid.New().String()
	err := tx.Table("USER").Create(user).Error
	return user, err
}
