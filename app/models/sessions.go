package models

import (
	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/helper"
	"github.com/twinj/uuid"
)

type sessionsHandler struct {
	PlatformRE string
}

func NewSessionsHandler() *sessionsHandler {
	return &sessionsHandler{
		PlatformRE: "(web|ios|android|live)+",
	}
}

func (h *sessionsHandler) Create(userID, deviceID, platform string, build int, name string) (*datautils.Session, error) {
	s := &datautils.Session{
		ID:          uuid.NewV4().String(),
		UserID:      userID,
		Type:        typeByPlatform(platform),
		DeviceID:    deviceID,
		Platform:    platform,
		Build:       build,
		Name:        name,
		AccessToken: uuid.NewV4().String(),
		Online:      true,
		CreatedAt:   helper.Timestamp(),
		UpdatedAt:   helper.Timestamp(),
	}

	// if s.DeviceID == "" {
	// 	return nil, errors.New("device_id is required")
	// }
	// if matched, _ := regexp.MatchString(h.PlatformRE, platform); !matched {
	// 	return nil, errors.New("platform invalid")
	// }
	// if s.Build == 0 {
	// 	return nil, errors.New("build invalid")
	// }
	if err := s.CreateSession(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *sessionsHandler) GetByAccessToken(access_token string) (datautils.Session, error) {
	ses, err := datautils.GetSessionByAccessToken(access_token)
	if err != nil {
		return ses, err
	}
	return ses, nil

}

func typeByPlatform(platform string) string {
	switch platform {
	case "ios":
		return "mobile"
	case "android":
		return "mobile"
	default:
		return "web"
	}
}
