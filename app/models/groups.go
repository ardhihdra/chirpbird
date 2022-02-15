package models

import (
	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/repository"
)

type GroupModel interface {
	GetByID(ID string) (*datautils.Group, error)
	ByIDAndUserID(ID, userID string) (*datautils.Group, error)
	ByUserIDs(userID string) (*[]datautils.Group, error)
	ByName(name string, exactmatch bool) (*[]datautils.Group, error)
	Create(name, userID string, userIDs []string) (*datautils.Group, error)
}

type groupModel struct{}

var (
	groupRepo repository.GroupRepository
)

func NewGroupsModel(repos repository.GroupRepository) GroupModel {
	groupRepo = repos
	return &groupModel{}
}

func (g *groupModel) GetByID(ID string) (*datautils.Group, error) {
	return groupRepo.GetByID(ID)
}

func (g *groupModel) ByIDAndUserID(ID, userID string) (*datautils.Group, error) {
	return groupRepo.ByIDAndUserID(ID, userID)
}

func (g *groupModel) ByUserIDs(userID string) (*[]datautils.Group, error) {
	return groupRepo.ByUserIDs(userID)
}

func (g *groupModel) ByName(name string, exactmatch bool) (*[]datautils.Group, error) {
	return groupRepo.ByName(name, exactmatch)
}

func (g *groupModel) Create(name, userID string, userIDs []string) (*datautils.Group, error) {
	return groupRepo.Create(name, userID, userIDs)
}
