package models

import "github.com/ardhihdra/chirpbird/helper"

type Event struct {
	ID        string           `bson:"_id"`
	Type      helper.EventType `bson:"type"`
	ObjectID  string           `bson:"object_id"`
	UserIDs   []string         `bson:"user_ids"`
	Timestamp int64            `bson:"timestamp"`
}

type events struct{}

var Events = new(events)
