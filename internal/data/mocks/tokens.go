package mocks

import (
	"time"

	"github.com/kirontoo/go-backend-template/internal/data"
)

type TokenModel struct{}

func (t TokenModel) DeleteAllForUser(scope string, userID int64) error {
	return nil
}

func (t TokenModel) Insert(token *data.Token) error {
	return nil
}

func (t TokenModel) New(userID int64, ttl time.Duration, scope string) (*data.Token, error) {
	return &data.Token{}, nil
}

func (t TokenModel) NewActivationToken(userID int64) (*data.Token, error) {
	return &data.Token{}, nil
}
