package entity

import "time"

type Movie struct {
	ID          int
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      float64
}

type MovieWithActors struct {
	*Movie
	Actors []*Actor
}

type UpdateMovieParams struct {
	Title       *string
	Description *string
	ReleaseDate *time.Time
	Rating      *float64
	ActorIDs    []int
}

func (m *Movie) MergeParams(params *UpdateMovieParams) []int {
	if params.Title != nil {
		m.Title = *params.Title
	}

	if params.Description != nil {
		m.Description = *params.Description
	}

	if params.ReleaseDate != nil {
		m.ReleaseDate = *params.ReleaseDate
	}

	if params.Rating != nil {
		m.Rating = *params.Rating
	}

	return params.ActorIDs
}
