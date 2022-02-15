package models

import (
	"strings"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/helper"
	"github.com/ardhihdra/chirpbird/app/repository"
	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserModel interface {
	Register(u *datautils.User) (*datautils.User, error)
	UsernameValid(u *datautils.User) error
	EmailValid(u *datautils.User) error
	PasswordValid(u *datautils.User) error
	ByUsername(username string, exactmatch bool) (*[]datautils.User, error)
	ByEmail(email string) (*datautils.User, error)
	ByID(ID string) (*datautils.User, error)
	Auth(userPassword, password string) bool
	UsernameAvailable(user datautils.User) bool
	CheckExpiry(id string) (*datautils.User, error)
	DeleteByID(id string)
}
type userModel struct {
	UsernameRE string
}

var (
	repo repository.UserRepository
)

func NewUsersHandler(repos repository.UserRepository) UserModel {
	repo = repos
	return &userModel{
		UsernameRE: "^[A-Za-z0-9_]{1,15}$",
	}
}

func (h *userModel) Register(u *datautils.User) (*datautils.User, error) {
	u.ID = uuid.NewV4().String()
	u.Username = strings.TrimSpace(u.Username)
	// u.Email = strings.TrimSpace(u.Email)
	// u.Password = strings.TrimSpace(u.Password)
	u.CreatedAt = helper.Timestamp()
	u.UpdatedAt = helper.Timestamp()

	if err := h.UsernameValid(u); err != nil {
		return nil, err
	}

	// if err := h.EmailValid(u); err != nil {
	// 	return nil, err
	// }

	// if err := h.PasswordValid(u); err != nil {
	// 	return nil, err
	// }

	// hpass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }
	// u.Password = string(hpass)

	if err := repo.CreateUser(*u); err != nil {
		return nil, err
	}

	return u, nil
}

func (h *userModel) UsernameValid(u *datautils.User) error {
	if u.Username == "" {
		return errors.New("username required")
	}

	//if matched, _ := regexp.MatchString(h.UsernameRE, u.Username); !matched {
	//	return errors.New("username invalid")
	//}

	if !repo.UsernameAvailable(*u) {
		return errors.New("username exists!")
	}
	return nil
}

func (h *userModel) EmailValid(u *datautils.User) error {
	if u.Email == "" {
		return errors.New("email required")
	}

	if !govalidator.IsEmail(u.Email) {
		return errors.New("email invalid")
	}

	if repo.EmailAvailable(*u) {
		return errors.New("email exists")
	}
	return nil
}

func (h *userModel) PasswordValid(u *datautils.User) error {
	if u.Password == "" {
		return errors.New("password required")
	}
	return nil
}

func (h *userModel) ByUsername(username string, exactmatch bool) (*[]datautils.User, error) {
	return repo.FindByUsername(username, exactmatch)
}

func (h *userModel) ByEmail(email string) (*datautils.User, error) {
	return repo.FindByEmail(email)
}

func (h *userModel) ByID(ID string) (*datautils.User, error) {
	return repo.FindByID(ID)
}

func (h *userModel) Auth(userPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func (h *userModel) UsernameAvailable(user datautils.User) bool {
	return repo.UsernameAvailable(user)
}

func (h *userModel) CheckExpiry(id string) (*datautils.User, error) {
	return repo.CheckExpiry(id)
}

func (h *userModel) DeleteByID(id string) {
	repo.DeleteByID(id)
}
