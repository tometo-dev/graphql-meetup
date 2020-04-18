package domain

import (
	"context"
	"errors"
	"github.com/tsuki42/graphql-meetup/logging"
	"github.com/tsuki42/graphql-meetup/models"
)

var (
	ErrorBadCredentials  = errors.New("entered email/password is wrong")
	ErrorUnauthenticated = errors.New("unauthenticated")
)

func (d *Domain) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := d.UserRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in user")
	}

	_, err = d.UserRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("username already in user")
	}

	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		logging.ERROR.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	// TODO: send verification code

	tx := d.UserRepo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if _, err := d.UserRepo.CreateUser(tx, user); err != nil {
		logging.ERROR.Printf("error creating user: %v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		logging.ERROR.Printf("error while commiting: %v", err)
		return nil, errors.New("something went wrong")
	}

	token, err := user.GenerateToken()
	if err != nil {
		logging.ERROR.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (d *Domain) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := d.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrorBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrorBadCredentials
	}

	token, err := user.GenerateToken()

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
