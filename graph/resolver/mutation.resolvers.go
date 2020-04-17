package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsuki42/graphql-meetup/middleware"

	"github.com/tsuki42/graphql-meetup/graph/generated"
	"github.com/tsuki42/graphql-meetup/logging"
	"github.com/tsuki42/graphql-meetup/models"
)

var (
	ErrorBadCredentials  = errors.New("entered email/password is wrong")
	ErrorUnauthenticated = errors.New("unauthenticated")
)

func (r *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	_, err := r.UserRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in user")
	}

	_, err = r.UserRepo.GetUserByUsername(input.Username)
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

	tx := r.UserRepo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if _, err := r.UserRepo.CreateUser(tx, user); err != nil {
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

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := r.UserRepo.GetUserByEmail(input.Email)
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

func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetupInput) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrorUnauthenticated
	}

	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}
	if len(input.Description) < 5 {
		return nil, errors.New("description not long enough")
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	return r.MeetupRepo.CreateMeetup(meetup)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetupInput) (*models.Meetup, error) {
	meetup, err := r.MeetupRepo.GetMeetupByID(id)

	updated := false

	if err != nil || meetup == nil {
		return nil, errors.New("meetup with given id not found")
	}

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name not long enough")
		}
		meetup.Name = *input.Name
		updated = true
	}

	if input.Description != nil {
		if len(*input.Description) < 5 {
			return nil, errors.New("description not long enough")
		}
		meetup.Description = *input.Description
		updated = true
	}

	if !updated {
		return nil, errors.New("nothing to update")
	}

	meetup, err = r.MeetupRepo.UpdateMeetup(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}
	return meetup, nil
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupRepo.GetMeetupByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup with given id doesn't exist")
	}

	if err = r.MeetupRepo.DeleteMeetup(id); err != nil {
		return false, err
	}
	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
