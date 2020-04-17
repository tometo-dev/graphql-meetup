package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type User struct {
	ID        string     `json:"id" gorm:"column:id; primary_key"`
	Username  string     `json:"username" gorm:"column:username" validate:"required"`
	Email     string     `json:"email" gorm:"column:email" validate:"required"`
	Password  string     `json:"password" gorm:"column:password" validate:"required"`
	FirstName string     `json:"firstName" gorm:"column:first_name" validate:"required"`
	LastName  string     `json:"lastName" gorm:"column:last_name" validate:"required"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(passwordHash)

	return nil
}

func (u *User) GenerateToken() (*AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "graphql-meetup",
	})

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}
