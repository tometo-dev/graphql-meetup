package domain

import (
	"errors"
	"github.com/tsuki42/graphql-meetup/models"
	"github.com/tsuki42/graphql-meetup/postgres"
)

var (
	ErrBadCredentials  = errors.New("entered email/password is wrong")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("unauthorized")
)

type Domain struct {
	UserRepo   postgres.UserRepo
	MeetupRepo postgres.MeetupRepo
}

func NewDomain(userRepo postgres.UserRepo, meetupRepo postgres.MeetupRepo) *Domain {
	return &Domain{UserRepo: userRepo, MeetupRepo: meetupRepo}
}

type Ownable interface {
	IsOwner(user *models.User) bool
}

func checkOwnership(o Ownable, user *models.User) bool {
	return o.IsOwner(user)
}
