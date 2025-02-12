package models

import (
	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/repository"
)

type MessageModel interface {
	ByID(ID string) (*datautils.Message, error)
	Create(groupID, userID, data string, ts int64) (*datautils.Message, error)
}

type messageModel struct{}

var (
	messageRepo repository.MessageRepository
)

func NewMessageModel(repos repository.MessageRepository) MessageModel {
	messageRepo = repos
	return &messageModel{}
}

func (m *messageModel) ByID(ID string) (*datautils.Message, error) {
	return messageRepo.ByID(ID)
}

func (m *messageModel) Create(groupID, userID, data string, ts int64) (*datautils.Message, error) {
	return messageRepo.Create(groupID, userID, data, ts)
}
