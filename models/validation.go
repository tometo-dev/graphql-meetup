package models

import "github.com/tsuki42/graphql-meetup/validator"

func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()
	v.Required("email", r.Email)
	v.IsEmail("email", r.Email)

	v.Required("password", r.Password)
	v.MinLength("password", r.Password, 8)

	v.Required("confirmPassword", r.ConfirmPassword)
	v.EqualToField("confirmPassword", r.ConfirmPassword, "password", r.Password)

	v.Required("username", r.Username)
	v.MinLength("username", r.Username, 3)

	v.Required("firstName", r.FirstName)
	v.MinLength("firstName", r.FirstName, 3)

	v.Required("lastName", r.LastName)
	v.MinLength("lastName", r.LastName, 3)

	return v.IsValid(), v.Errors
}

func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()
	v.Required("email", l.Email)
	v.IsEmail("email", l.Email)

	v.Required("password", l.Password)

	return v.IsValid(), v.Errors
}
