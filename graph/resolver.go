package graph

import "github.com/tsuki42/graphql-meetup/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	MeetupRepo postgres.MeetupRepo
	UserRepo postgres.UserRepo
}
