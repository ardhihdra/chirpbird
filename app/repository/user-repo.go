package repository

import "github.com/ardhihdra/chirpbird/app/datautils"

type UserRepository interface {
	FindByUsername(username string, exactmatch bool) (*[]datautils.User, error)
	FindByEmail(email string) (*datautils.User, error)
	FindByID(ID string) (*datautils.User, error)
	CheckExpiry(id string) (*datautils.User, error)
	EmailAvailable(user datautils.User) bool
	UsernameAvailable(user datautils.User) bool
	CreateUser(u datautils.User) error
	DeleteByID(id string)
}
