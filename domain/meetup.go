package domain

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsuki42/graphql-meetup/middleware"
	"github.com/tsuki42/graphql-meetup/models"
)

func (d *Domain) CreateMeetup(ctx context.Context, input models.NewMeetupInput) (*models.Meetup, error) {
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

	return d.MeetupRepo.CreateMeetup(meetup)
}

func (d *Domain) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetupInput) (*models.Meetup, error) {
	meetup, err :=d.MeetupRepo.GetMeetupByID(id)

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

	meetup, err =d.MeetupRepo.UpdateMeetup(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}
	return meetup, nil
}

func (d *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err :=d.MeetupRepo.GetMeetupByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup with given id doesn't exist")
	}

	if err =d.MeetupRepo.DeleteMeetup(id); err != nil {
		return false, err
	}
	return true, nil
}
