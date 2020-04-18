package domain

import "github.com/tsuki42/graphql-meetup/postgres"

type Domain struct {
	UserRepo postgres.UserRepo
	MeetupRepo postgres.MeetupRepo
}

func NewDomain(userRepo postgres.UserRepo, meetupRepo postgres.MeetupRepo) *Domain {
	return &Domain{UserRepo: userRepo, MeetupRepo: meetupRepo}
}
