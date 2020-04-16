package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/tsuki42/graphql-meetup/graph/generated"
	"github.com/tsuki42/graphql-meetup/graph/model"
	"github.com/tsuki42/graphql-meetup/models"
)

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "Meetup 1",
		Description: "First Meetup",
		User:        users[0],
	},
	{
		ID:          "2",
		Name:        "Meetup 2",
		Description: "Second meetup",
		User:        users[1],
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "bob",
		Email:    "bob@bob.com",
	},
	{
		ID:       "2",
		Username: "alice",
		Email:    "alice@alice.com",
	},
}

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	for _, user := range users {
		if user.ID == obj.User.ID {
			return user, nil
		}
	}
	return nil, errors.New("user does not exist")
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var userMeetups []*models.Meetup

	for _, m := range meetups {
		if m.User.ID == obj.ID {
			userMeetups = append(userMeetups, m)
		}
	}
	return userMeetups, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
