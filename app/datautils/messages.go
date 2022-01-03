package datautils

type Body struct {
	Data string `json:"data"`
}

type Message struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	GroupID   string `json:"group_id"`
	Body      Body   `json:"body"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
