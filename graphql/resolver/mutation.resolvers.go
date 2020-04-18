package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/tsuki42/graphql-meetup/graphql/generated"
	"github.com/tsuki42/graphql-meetup/models"
)

var (
	ErrInput = errors.New("input error")
)

func (r *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, ErrInput
	}

	return r.Domain.Register(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, ErrInput
	}

	return r.Domain.Login(ctx, input)
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetupInput) (*models.Meetup, error) {
	return r.Domain.CreateMeetup(ctx, input)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetupInput) (*models.Meetup, error) {
	return r.Domain.UpdateMeetup(ctx, id, input)
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return r.Domain.DeleteMeetup(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
