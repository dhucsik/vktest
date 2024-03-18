package entity

import "time"

type Actor struct {
	ID          int
	FullName    string
	Gender      string
	DateOfBirth time.Time
}

type ActorWithMovies struct {
	*Actor
	Movies []*Movie
}

type UpdateActorParams struct {
	FullName    *string
	Gender      *string
	DateOfBirth *time.Time
}

func (a *Actor) MergeParams(params *UpdateActorParams) {
	if params.FullName != nil {
		a.FullName = *params.FullName
	}

	if params.Gender != nil {
		a.Gender = *params.Gender
	}

	if params.DateOfBirth != nil {
		a.DateOfBirth = *params.DateOfBirth
	}
}
