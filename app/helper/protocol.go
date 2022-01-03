package helper

type EventType int

const (
	EVENT_MESSAGE           EventType = 20
	EVENT_MESSAGE_SENT      EventType = 21
	EVENT_MESSAGE_DELIVERED EventType = 22
	EVENT_MESSAGE_READ      EventType = 23
	//EVENT_MESSAGE_UPDATED   EventType = 24
	//EVENT_MESSAGE_DELETED   EventType = 25
	EVENT_TYPING_START EventType = 40
	EVENT_TYPING_END   EventType = 41
	EVENT_GROUP        EventType = 70
	//EVENT_GROUP_UPDATED     EventType = 71
	EVENT_GROUP_JOINED EventType = 72
	EVENT_GROUP_LEFT   EventType = 73
)

type Event struct {
	Type      EventType   `json:"type"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Body      interface{} `json:"body,omitempty"`
}

type EventMessage struct {
	MessageID string `json:"message_id,omitempty"`
	Data      string `json:"data,omitempty"`
}

type EventMessageSent struct {
	MessageID string `json:"message_id,omitempty"`
}

type EventMessageDelivered struct {
	MessageID string `json:"message_id,omitempty"`
}

type EventTypingStart struct {
	GroupID string `json:"group_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}

type EventTypingEnd struct {
	GroupID string `json:"group_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}

type EventGroup struct {
	GroupID string   `json:"group_id,omitempty"`
	Name    string   `json:"name,omitempty"`
	UserIDs []string `json:"user_ids,omitempty"`
}

type EventGroupJoined struct {
	GroupID string `json:"group_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}

type EventGroupLeft struct {
	GroupID string `json:"group_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}
