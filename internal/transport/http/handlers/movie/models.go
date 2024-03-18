package movie

import (
	"time"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/samber/lo"
)

type createMovieRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
	ActorIDs    []int   `json:"actor_ids"`
}

func (r createMovieRequest) convert() (*entity.Movie, error) {
	date, err := time.Parse(time.DateOnly, r.ReleaseDate)
	if err != nil {
		return nil, err
	}

	return &entity.Movie{
		Title:       r.Title,
		Description: r.Description,
		ReleaseDate: date,
		Rating:      r.Rating,
	}, nil
}

type createMovieResponse struct {
	ID int `json:"id"`
}

type updateMovieRequest struct {
	Title       *string  `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	ReleaseDate *string  `json:"release_date,omitempty"`
	Rating      *float64 `json:"rating,omitempty"`
	ActorIDs    []int    `json:"actor_ids,omitempty"`
}

func (r updateMovieRequest) convert() (*entity.UpdateMovieParams, error) {
	var releaseDate *time.Time
	if r.ReleaseDate != nil {
		date, err := time.Parse(time.DateOnly, *r.ReleaseDate)
		if err != nil {
			return nil, err
		}

		releaseDate = &date
	}

	return &entity.UpdateMovieParams{
		Title:       r.Title,
		Description: r.Description,
		ReleaseDate: releaseDate,
		Rating:      r.Rating,
		ActorIDs:    r.ActorIDs,
	}, nil
}

type getMovieResponse struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ReleaseDate string          `json:"release_date"`
	Rating      float64         `json:"rating"`
	Actors      []actorResponse `json:"actors"`
}

func convertMovie(m *entity.MovieWithActors) getMovieResponse {
	actors := lo.Map(m.Actors, func(actor *entity.Actor, _ int) actorResponse {
		return convertActor(actor)
	})

	return getMovieResponse{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		ReleaseDate: m.ReleaseDate.Format(time.DateOnly),
		Rating:      m.Rating,
		Actors:      actors,
	}
}

type actorResponse struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

func convertActor(a *entity.Actor) actorResponse {
	return actorResponse{
		ID:          a.ID,
		FullName:    a.FullName,
		Gender:      a.Gender,
		DateOfBirth: a.DateOfBirth.Format(time.DateOnly),
	}
}

type getMoviesResponse struct {
	Movies []getMovieResponse `json:"movies"`
}

func convertMovies(movies []*entity.MovieWithActors) getMoviesResponse {
	return getMoviesResponse{
		Movies: lo.Map(movies, func(movie *entity.MovieWithActors, _ int) getMovieResponse {
			return convertMovie(movie)
		}),
	}
}

type errorResponse struct {
	Error string `json:"error"`
}
