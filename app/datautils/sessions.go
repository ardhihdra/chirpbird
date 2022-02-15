package datautils

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/app/db"
)

type Session struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Type        string `json:"type"`
	DeviceID    string `json:"device_id"`
	Platform    string `json:"platform"`
	Build       int    `json:"build"`
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
	Online      bool   `json:"online"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	OnlineAt    int64  `json:"online_at"`
	OfflineAt   int64  `json:"offline_at"`
}

func (s *Session) CreateSession() error {
	sesMarshal, _ := json.Marshal(s)
	res, err := db.Elastic.Index(
		db.IdxSessions,                        // Index name
		strings.NewReader(string(sesMarshal)), // Document body
		db.Elastic.Index.WithDocumentID(s.ID), // Document ID
		db.Elastic.Index.WithRefresh("true"),  // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		db.PrintErrorResponse(res)
		return err
	}

	return nil
}

// func (s *Session) MessagingURL() string {
// 	return fmt.Sprintf("%s/%s", "ws://127.0.0.1:3000", s.AccessToken)
// }

func (s *Session) DeleteByID() {
	db.Elastic.Delete(db.IdxSessions, s.ID)
}

func GetSessionByAccessToken(access_token string) (Session, error) {
	query := db.MatchCondition(map[string]interface{}{
		"access_token": access_token,
	})
	var ses Session
	values, err := db.FindOne(query, db.IdxSessions)
	if err != nil {
		return ses, err
	}
	json.Unmarshal([]byte(values[1].String()), &ses)
	return ses, nil
}

func GetSessionByUserID(userID string) ([]*Session, error) {
	query := db.MatchCondition(map[string]interface{}{
		"user_id": userID,
	})
	var sessions []*Session
	values, err := db.FindAll(query, db.IdxSessions)
	arrVal := values[1].Array()
	var ses Session
	for i := range arrVal {
		json.Unmarshal([]byte(arrVal[i].Get("_source").String()), &ses)
		sessions = append(sessions, &ses)
	}
	if err != nil {
		return sessions, err
	}
	return sessions, nil
}

// func trial(values []gjson.Result, target []*interface{}) {
// 	arrVal := values[1].Array()
// 	var usr interface{}
// 	for i := range arrVal {
// 		json.Unmarshal([]byte(arrVal[i].Get("_source").String()), &usr)
// 		target = append(target, &usr)
// 	}
// }
