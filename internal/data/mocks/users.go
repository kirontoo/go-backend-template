package mocks

import "github.com/kirontoo/go-backend-template/internal/data"

type UserModel struct{}

func (u UserModel) Insert(user *data.User) error {
	return nil
}

func (u UserModel) GetByEmail(email string) (*data.User, error) {
	return &data.User{}, nil
}

func (u UserModel) Update(user *data.User) error {
	return nil
}

func (u UserModel) GetForToken(tokenScope, tokenPlaintext string) (*data.User, error) {
	return &data.User{}, nil
}
