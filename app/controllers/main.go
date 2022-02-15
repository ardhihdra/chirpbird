package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/helper/jwt"
)

type BaseController struct {
	User *datautils.User
}

func (c *BaseController) Authenticate(r *http.Request) (err error) {
	token, err := jwt.Parse(r)
	if err != nil {
		return errors.New("failed to parse token")
	}
	user, err := userModel.CheckExpiry(token.UserID)

	if err != nil || user == nil {
		fmt.Println(err)
		return errors.New("Unauthenticate")
	}
	c.User = user
	return nil
}
