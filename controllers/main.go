package controllers

import (
	"fmt"
	"net/http"

	"github.com/ardhihdra/chirpbird/datautils"
	"github.com/ardhihdra/chirpbird/helper/jwt"
)

type BaseController struct {
	User *datautils.User
}

func (c *BaseController) Authenticate(r *http.Request) (err error) {
	token, err := jwt.Parse(r)
	if err != nil {
		return
	}
	u, err := users.ByID(token.UserID)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.User = u
	return nil
}
