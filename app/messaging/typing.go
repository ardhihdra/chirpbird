package messaging

import (
	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/helper"
)

type typing struct{}

func newTyping() *typing {
	return &typing{}
}

func (t *typing) Start(c *datautils.UserConnection, p *datautils.RpcTypingStart) {
	group := withGroup(p.GroupID, c.UserID)
	e := datautils.NewEvent(datautils.EVENT_TYPING_START, helper.Timestamp())
	e.Body = datautils.EventTypingStart{
		GroupID: group.ID,
		UserID:  c.UserID,
	}
	e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
}

func (t *typing) End(c *datautils.UserConnection, p *datautils.RpcTypingEnd) {
	group := withGroup(p.GroupID, c.UserID)
	e := datautils.NewEvent(datautils.EVENT_TYPING_END, helper.Timestamp())
	e.Body = datautils.EventTypingEnd{
		GroupID: group.ID,
		UserID:  c.UserID,
	}
	e.SendToUsersWithoutMe(c.SessionID, group.UserIDs)
}
