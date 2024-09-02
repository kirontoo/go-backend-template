package data

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrEditConflict    = errors.New("edit conflict")
	ThreeSecondTimeout = 3 * time.Second
)

type Models struct {
	Movies      Movies
	Permissions Permissions
	Tokens      Tokens
	Users       Users
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Users:       UserModel{DB: db},
	}
}
