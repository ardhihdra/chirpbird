package models

import (
	"strings"

	"github.com/ardhihdra/chirpbird/datautils"
	"github.com/ardhihdra/chirpbird/db"
	"github.com/ardhihdra/chirpbird/helper"
	"github.com/asaskevich/govalidator"
	"github.com/twinj/uuid"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type usersHandler struct {
	UsernameRE string
}

func NewUsersHandler() *usersHandler {
	return &usersHandler{
		UsernameRE: "^[A-Za-z0-9_]{1,15}$", // username length copied from Twitter
	}
}

func (h *usersHandler) Register(u *datautils.User) (*datautils.User, error) {
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

	if err := u.CreateUser(); err != nil {
		return nil, err
	}

	return u, nil
}

func (h *usersHandler) UsernameValid(u *datautils.User) error {
	if u.Username == "" {
		return errors.New("username required")
	}

	//if matched, _ := regexp.MatchString(h.UsernameRE, u.Username); !matched {
	//	return errors.New("username invalid")
	//}

	if !u.UsernameAvailable() {
		return errors.New("username exists!")
	}
	return nil
}

func (h *usersHandler) EmailValid(u *datautils.User) error {
	if u.Email == "" {
		return errors.New("email required")
	}

	if !govalidator.IsEmail(u.Email) {
		return errors.New("email invalid")
	}

	if u.EmailAvailable() {
		return errors.New("email exists")
	}
	return nil
}

func (h *usersHandler) PasswordValid(u *datautils.User) error {
	if u.Password == "" {
		return errors.New("password required")
	}
	return nil
}

func (h *usersHandler) ByUsername(username string) (*datautils.User, error) {
	var u *datautils.User
	query := db.MatchCondition(map[string]interface{}{
		"username":        strings.ToLower(username),
		"_source":         true,
		"terminate_after": 1,
	})
	return u, datautils.FindOne(query, db.IdxUsers, &u)
}

func (h *usersHandler) ByEmail(email string) (*datautils.User, error) {
	var u *datautils.User
	query := db.MatchCondition(map[string]interface{}{"email": strings.ToLower(email)})
	return u, datautils.FindOne(query, db.IdxUsers, &u)
}

func (h *usersHandler) ByID(ID string) (*datautils.User, error) {
	var u *datautils.User
	query := db.MatchCondition(map[string]interface{}{"_id": strings.ToLower(ID)})
	return u, datautils.FindOne(query, db.IdxUsers, &u)
}

func (h *usersHandler) Auth(userPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
