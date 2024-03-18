package actor

import (
	"time"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/samber/lo"
)

type createActorRequest struct {
	FullName    string `json:"full_name"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

func (r createActorRequest) convert() (*entity.Actor, error) {
	dateOfBirth, err := time.Parse(time.DateOnly, r.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &entity.Actor{
		FullName:    r.FullName,
		Gender:      r.Gender,
		DateOfBirth: dateOfBirth,
	}, nil
}

type createActorResponse struct {
	ID int `json:"id"`
}

type updateActorRequest struct {
	FullName    *string `json:"full_name,omitempty"`
	Gender      *string `json:"gender,omitempty"`
	DateOfBirth *string `json:"date_of_birth,omitempty"`
}

func (r updateActorRequest) convert() (*entity.UpdateActorParams, error) {
	var dateOfBirth *time.Time
	if r.DateOfBirth != nil {
		dob, err := time.Parse(time.DateOnly, *r.DateOfBirth)
		if err != nil {
			return nil, err
		}

		dateOfBirth = &dob
	}

	return &entity.UpdateActorParams{
		FullName:    r.FullName,
		Gender:      r.Gender,
		DateOfBirth: dateOfBirth,
	}, nil
}

type getActorResponse struct {
	ID          int             `json:"id"`
	FullName    string          `json:"full_name"`
	Gender      string          `json:"gender"`
	DateOfBirth string          `json:"date_of_birth"`
	Movies      []movieResponse `json:"movies"`
}

func convertActor(actor *entity.ActorWithMovies) getActorResponse {
	movies := lo.Map(actor.Movies, func(movie *entity.Movie, _ int) movieResponse {
		return convertMovie(movie)
	})

	return getActorResponse{
		ID:          actor.ID,
		FullName:    actor.FullName,
		Gender:      actor.Gender,
		DateOfBirth: actor.DateOfBirth.Format(time.DateOnly),
		Movies:      movies,
	}
}

type movieResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
}

func convertMovie(movie *entity.Movie) movieResponse {
	return movieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate.Format(time.DateOnly),
		Rating:      movie.Rating,
	}
}

func convertActors(actors []*entity.ActorWithMovies) []getActorResponse {
	return lo.Map(actors, func(actor *entity.ActorWithMovies, _ int) getActorResponse {
		return convertActor(actor)
	})
}

type errorResponse struct {
	Error string `json:"error"`
}
