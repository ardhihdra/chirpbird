package datautils

/** RPC enum */
type RpcMethod int

const (
	RPC_MESSAGE_GET       RpcMethod = 20
	RPC_MESSAGE_SEND      RpcMethod = 40
	RPC_MESSAGE_DELIVERED RpcMethod = 41
	RPC_MESSAGE_READ      RpcMethod = 42
	//RPC_MESSAGE_UPDATED   RpcMethod = 43
	//RPC_MESSAGE_DELETED   RpcMethod = 44
	RPC_TYPING_START RpcMethod = 60
	RPC_TYPING_END   RpcMethod = 61
)

/**
* Remote Protocol Call, just some way to
* differentiate socket datas and how to deal with it
 */
type RPC struct {
	Method RpcMethod   `json:"method"`
	Body   interface{} `json:"body,omitempty"`
}

type RpcMessageGet struct {
	Timestamp int64 `json:"timestamp,omitempty"`
}

type RpcMessageSend struct {
	GroupID string `json:"group_id,omitempty"`
	Data    string `json:"data,omitempty"`
}

type RpcMessageDelivered struct {
	MessageID string `json:"message_id,omitempty"`
}

type RpcMessageRead struct {
	MessageID string `json:"message_id,omitempty"`
}

type RpcTypingStart struct {
	GroupID string `json:"group_id,omitempty"`
}

type RpcTypingEnd struct {
	GroupID string `json:"group_id,omitempty"`
}
