package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/db"
	"github.com/twinj/uuid"
)

type Message struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	GroupID   string `json:"group_id"`
	Body      Body   `json:"body"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type Body struct {
	Data string `json:"data"`
}

type messages struct{}

var Messages = new(messages)

func (messages) ByID(ID string) (*Message, error) {
	var m Message
	query := db.MatchCondition(map[string]interface{}{"id": ID})
	values, err := db.FindOne(query, db.IdxMessages)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return &m, json.Unmarshal([]byte(values[1].String()), &m)
}

func (messages) Create(groupID, userID, data string, ts int64) (*Message, error) {
	m := &Message{
		ID:        uuid.NewV4().String(),
		UserID:    userID,
		GroupID:   groupID,
		Body:      Body{Data: data},
		CreatedAt: ts,
		UpdatedAt: ts,
	}

	marshaled, _ := json.Marshal(m)
	res, err := db.Elastic.Index(
		db.IdxMessages,                        // Index name
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
