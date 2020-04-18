package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tsuki42/graphql-meetup/graphql/generated"
	"github.com/tsuki42/graphql-meetup/models"
)

func (r *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilterInput, limit *int, offset *int) ([]*models.Meetup, error) {
	return r.Domain.MeetupRepo.GetMeetups(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.Domain.UserRepo.GetUserByID(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
