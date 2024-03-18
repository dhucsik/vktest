package actor

import (
	"time"

	"github.com/dhucsik/vktest/internal/entity"
)

type actorModel struct {
	ID          int       `db:"id"`
	FullName    string    `db:"full_name"`
	Gender      string    `db:"age"`
	DateOfBirth time.Time `db:"date_of_birth"`
}

func convertActor(actr *entity.Actor) *actorModel {
	return &actorModel{
		ID:          actr.ID,
		FullName:    actr.FullName,
		Gender:      actr.Gender,
		DateOfBirth: actr.DateOfBirth,
	}
}

func (m *actorModel) convert() *entity.Actor {
	return &entity.Actor{
		ID:          m.ID,
		FullName:    m.FullName,
		Gender:      m.Gender,
		DateOfBirth: m.DateOfBirth,
	}
}
