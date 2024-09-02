package mocks

import (
	"time"

	"github.com/kirontoo/go-backend-template/internal/data"
)

var mockMovie = &data.Movie{
	ID:        1,
	CreatedAt: time.Now(),
	Title:     "Beetlejuice Beetlejuice",
	Year:      2024,
	Runtime:   104,
	Genres:    []string{"dark comedy", "dark fantasy", "comdey", "horror", "fantasy", "supernatural horror"},
	Version:   1,
}

type MovieModel struct{}

func (m MovieModel) Insert(movie *data.Movie) error {
	return nil
}

func (m MovieModel) Get(id int64) (*data.Movie, error) {
	switch id {
	case 1:
		return mockMovie, nil
	default:
		return nil, data.ErrRecordNotFound
	}
}

func (m MovieModel) Update(movie *data.Movie) error {
	return nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
}

func (m MovieModel) GetAll(title string, genres []string, filters data.Filters) ([]*data.Movie, data.Metadata, error) {
	return []*data.Movie{mockMovie}, data.Metadata{}, nil
}
