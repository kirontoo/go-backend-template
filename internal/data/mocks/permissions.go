package mocks

import (
	"fmt"

	"github.com/kirontoo/go-backend-template/internal/data"
)

type PermisionModel struct{}

func (p PermisionModel) AddForUser(userID int64, codes ...string) error {

	return nil
}

func (p PermisionModel) GetAllForUser(userID int64) (data.PermissionList, error) {
	switch userID {
	case 1:
		return data.PermissionList{"movies:read"}, nil
	case 2:
		return data.PermissionList{"movies:read", "movies:write"}, nil
	case 3:
		return data.PermissionList{}, nil
	default:
		return nil, fmt.Errorf("user does not exist")
	}
}
