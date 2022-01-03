package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/db"
	"github.com/twinj/uuid"
)

type messages struct{}

var Messages = new(messages)

func (messages) ByID(ID string) (*datautils.Message, error) {
	var m datautils.Message
	query := db.MatchCondition(map[string]interface{}{"id": ID})
	values, err := db.FindOne(query, db.IdxMessaging)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return &m, json.Unmarshal([]byte(values[1].String()), &m)
}

func (messages) Create(groupID, userID, data string, ts int64) (*datautils.Message, error) {
	m := &datautils.Message{
		ID:        uuid.NewV4().String(),
		UserID:    userID,
		GroupID:   groupID,
		Body:      datautils.Body{Data: data},
		CreatedAt: ts,
		UpdatedAt: ts,
	}

	marshaled, _ := json.Marshal(m)
	res, err := db.Elastic.Index(
		db.IdxMessaging,                       // Index name
		strings.NewReader(string(marshaled)),  // Document body
		db.Elastic.Index.WithDocumentID(m.ID), // Document ID
		db.Elastic.Index.WithRefresh("true"),  // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		db.PrintErrorResponse(res)
		return nil, err
	}

	return m, nil
}
