package repository

import "github.com/ardhihdra/chirpbird/app/datautils"

type SessionRepository interface {
	Create(userID, deviceID, platform string, build int, name string) (*datautils.Session, error)
	GetByAccessToken(access_token string) (datautils.Session, error)
}
