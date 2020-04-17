package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/tsuki42/graphql-meetup/graph/generated"
	"github.com/tsuki42/graphql-meetup/models"
)

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupRepo.GetMeetups()
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UserRepo.GetUserByID(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
