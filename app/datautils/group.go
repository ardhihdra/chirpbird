package datautils

type Group struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	UserID    string   `json:"user_id"`
	UserIDs   []string `json:"user_ids"`
	Deleted   []string `json:"deleted"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}
