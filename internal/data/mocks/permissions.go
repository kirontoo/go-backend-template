package mocks

import "github.com/kirontoo/go-backend-template/internal/data"

type PermisionModel struct{}

func (p PermisionModel) AddForUser(userID int64, codes ...string) error {

	return nil
}

func (p PermisionModel) GetAllForUser(userID int64) (data.PermissionList, error) {
	return data.PermissionList{}, nil
}
