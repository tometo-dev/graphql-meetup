package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/tsuki42/graphql-meetup/graph/dataloader"
	"github.com/tsuki42/graphql-meetup/graph/generated"
	"github.com/tsuki42/graphql-meetup/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return dataloader.GetUserLoader(ctx).Load(obj.UserID)
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

type meetupResolver struct{ *Resolver }
