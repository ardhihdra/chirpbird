package jwt

import (
	"net/http"

	"github.com/golang-jwt/jwt"

	"errors"
	"fmt"
	"strings"
	"time"
)

const alg = "HS256"

type JWTToken struct {
	UserID string
}

var signingKey = []byte("signingKey")

func Create(userID string) string {
	tokenDuration, _ := time.ParseDuration(fmt.Sprintf("%ds", 1209600))

	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(tokenDuration).Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod(alg), claims)

	tokenStr, err := token.SignedString([]byte("signingKey"))
	if err != nil {
		panic(err)
	}

	return tokenStr
}

func Parse(r *http.Request) (*JWTToken, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("Only accept POST method")
	}
	tokenStr := getTokenString(r)
	if tokenStr == "" {
		return nil, errors.New("jwt token invalid")
	}
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Method.Alg() != alg {
		return nil, errors.New("wrong alg")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &JWTToken{UserID: claims["user_id"].(string)}, nil
	} else {
		return nil, errors.New("jwt token wrong")
	}
}

func getTokenString(r *http.Request) string {
	ah := r.Header.Get("Authorization")
	if ah == "" || len(ah) < 8 || strings.ToLower(ah[:6]) != "bearer" {
		return ""
	}
	return ah[7:]
}
