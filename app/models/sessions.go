package models

import (
	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/repository"
)

type SessionModel interface {
	Create(userID, deviceID, platform string, build int, name string) (*datautils.Session, error)
	GetByAccessToken(access_token string) (datautils.Session, error)
}
type sessionModel struct {
	PlatformRE string
}

var (
	sessionRepo repository.SessionRepository
)

func NewSessionsModel(repos repository.SessionRepository) SessionModel {
	sessionRepo = repos
	return &sessionModel{
		PlatformRE: "(web|ios|android|live)+",
	}
}

func (h *sessionModel) Create(userID, deviceID, platform string, build int, name string) (*datautils.Session, error) {
	return sessionRepo.Create(userID, deviceID, platform, build, name)
}

func (s *sessionModel) GetByAccessToken(access_token string) (datautils.Session, error) {
	return sessionRepo.GetByAccessToken((access_token))
}
