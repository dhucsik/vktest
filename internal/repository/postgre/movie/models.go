package movie

import (
	"time"

	"github.com/dhucsik/vktest/internal/entity"
)

type movieModel struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ReleaseDate time.Time `db:"release_date"`
	Rating      float64   `db:"rating"`
}

func convertMovie(m *entity.Movie) *movieModel {
	return &movieModel{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		ReleaseDate: m.ReleaseDate,
		Rating:      m.Rating,
	}
}

func (m *movieModel) convert() *entity.Movie {
	return &entity.Movie{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		ReleaseDate: m.ReleaseDate,
		Rating:      m.Rating,
	}
}
