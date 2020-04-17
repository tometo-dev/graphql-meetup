package resolver

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

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}
	if len(input.Description) < 5 {
		return nil, errors.New("description not long enough")
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}

	return r.MeetupRepo.CreateMeetup(meetup)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*models.Meetup, error) {
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
