package messaging

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ardhihdra/chirpbird/datautils"
	"github.com/ardhihdra/chirpbird/db"
	"github.com/ardhihdra/chirpbird/models"
	"github.com/gorilla/websocket"
)

var eventsService = newEvents()
var messageService = newMessages()
var typingService = newTyping()

var sessions = models.NewSessionsHandler()

func init() {
	go publishListener()
}

func publishListener() {
	db.Redis.Subscribe(func(channel string, data []byte) {
		chunks := strings.Split(channel, ":")
		sessionID := chunks[len(chunks)-1]
		conn, err := datautils.ConnectionBySessionID(sessionID)
		if err != nil {
			return
		}
		conn.SetWriteDeadline(time.Now().Add(datautils.WriteWait))
		conn.WriteMessage(websocket.TextMessage, data)
	})
}

func Start() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			access_token := r.URL.Query().Get("access_token")
			s, err := sessions.GetByAccessToken(access_token)
			if err != nil {
				fmt.Println("Error: Unauthorized, wrong token or expired")
			}
			userConn, err := datautils.CreateUserConnection(s, w, r)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			}
			quit := make(chan int)
			msgChan := make(chan *datautils.RPC)
			actChan := make(chan *datautils.RPC)
			go dispatcher(userConn, msgChan, quit)
			go dispatcher(userConn, actChan, quit)
			rpcReader(userConn, msgChan, actChan, quit)
		},
	)
}

func dispatcher(c *datautils.UserConnection, ch chan *datautils.RPC, quit chan int) {
	defer func() {
		fmt.Println("MESSAGING: Disconnect from client ID", c.UserID)
	}()

	for {
		select {
		case rm, ok := <-ch:
			if !ok {
				fmt.Println("Couldn't receive from client ID", c.UserID)
				c.SendCloseConnection()
				return
			}
			HandleMessaging(c, rm)
		case <-quit:
			fmt.Println("Server is disconnecting ws")
			return
		default:
			// BAD Busy Waiting handler
			time.Sleep(50 * time.Millisecond)
		}

	}
}

func rpcReader(c *datautils.UserConnection, mCh chan *datautils.RPC, aCh chan *datautils.RPC, quit chan int) {
	defer fmt.Println("MESSAGING: rpc reader disconnected from client ID", c.UserID)
	defer c.Connection.Close()

	for {
		mtype, data, err := c.Connection.ReadMessage()
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		if mtype == 0 {
			fmt.Println("Error: Invalid socket data")
		}
		rpc := &datautils.RPC{}
		if err := json.Unmarshal(data, rpc); err != nil {
			fmt.Println("Error: ", err.Error())
		}
		isActChan := rpc.Method == datautils.RPC_TYPING_START || rpc.Method == datautils.RPC_TYPING_END
		if isActChan {
			aCh <- rpc
		} else {
			mCh <- rpc
		}
	}
}

func HandleMessaging(c *datautils.UserConnection, r *datautils.RPC) {
	parseMsg := func(obj interface{}) {
		byteData, _ := json.Marshal(r.Body)
		if err := json.Unmarshal(byteData, obj); err != nil {
			fmt.Println("Error while parse ws message", r.Body, err)
		}
	}

	switch r.Method {
	case datautils.RPC_MESSAGE_GET:
		params := datautils.RpcMessageGet{}
		parseMsg(&params)
		eventsService.Get(c, &params)

	case datautils.RPC_MESSAGE_SEND:
		params := datautils.RpcMessageSend{}
		parseMsg(&params)
		messageService.Send(c, &params)

	case datautils.RPC_MESSAGE_DELIVERED:
		params := datautils.RpcMessageDelivered{}
		parseMsg(&params)
		messageService.Delivered(c, &params)

	case datautils.RPC_MESSAGE_READ:
		params := datautils.RpcMessageRead{}
		parseMsg(&params)
		messageService.Read(c, &params)

	case datautils.RPC_TYPING_START:
		params := datautils.RpcTypingStart{}
		parseMsg(&params)
		typingService.Start(c, &params)

	case datautils.RPC_TYPING_END:
		params := datautils.RpcTypingEnd{}
		parseMsg(&params)
		typingService.End(c, &params)
	}
}

func withGroup(groupID, userID string) *models.Group {
	g, err := models.Groups.ByIDAndUserID(groupID, userID)
	if err != nil {
		return nil
	}
	return g
}
