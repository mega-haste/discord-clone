package SocketIO

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ID_t uint32
type OnConnectHandler func(ID ID_t) error
type OnDisconnectHandler func()
type OnErrorHandler func(err error) bool
type EventsCall func(ID ID_t, msg SocketMessageEvent)

type SocketMessageEvent struct {
	EventKey string   `json:"eventKey"`
	Message  string   `json:"message"`
	To       string   `json:"to"`
	Except   []string `json:"except"`
}

type ClientsMap map[ID_t]Client

type SocketIO struct {
	clients      ClientsMap
	onConnect    OnConnectHandler
	onDisconnect OnDisconnectHandler
	onError      OnErrorHandler
	events       map[string]EventsCall
}

func New() SocketIO {
	return SocketIO{
		clients:      make(ClientsMap),
		onConnect:    func(ID ID_t) error { return nil },
		onDisconnect: func() {},
		onError:      func(err error) bool { return true },
		events:       make(map[string]EventsCall),
	}
}

func (s *SocketIO) OnConnect(callback OnConnectHandler) {
	s.onConnect = callback
}

func (s *SocketIO) OnDisconnect(callback OnDisconnectHandler) {
	s.onDisconnect = callback
}

func (s *SocketIO) OnError(callback OnErrorHandler) {
	s.onError = callback
}

func (s *SocketIO) On(eventId string, callback EventsCall) {
	s.events[eventId] = callback
}

func (s *SocketIO) Emit(eventId string, message string) {
	for _, v := range s.clients {
		v.Conn.WriteJSON(SocketMessageEvent{
			EventKey: eventId,
			Message:  message,
		})
	}
}

func (s *SocketIO) To(ID ID_t) SocketEmiter {
	res := NewSocketEmiter(s)
	res.To(ID)
	return res
}

func (s *SocketIO) Except(ID ID_t) SocketEmiter {
	res := NewSocketEmiter(s)
	res.Except(ID)
	return res
}

func (s *SocketIO) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.onError(err)
		return
	}
	var ID ID_t
	{
		currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
		uniqeID := uuid.New().ID()
		ID = ID_t(uint32(currentTimestamp) + uniqeID)
	}

	s.clients[ID] = NewClient(conn)
	defer s.onDisconnect()
	defer delete(s.clients, ID)
	defer conn.Close()

	s.To(ID).Emit("connected", fmt.Sprint(ID))
	s.onConnect(ID)

	echo(conn, s, ID)
}

func echo(conn *websocket.Conn, s *SocketIO, ID ID_t) {
	for {
		socketData := &SocketMessageEvent{}
		err := conn.ReadJSON(socketData)
		if err != nil {
			break
		}

		call, exits := s.events[socketData.EventKey]

		if exits {
			call(ID, *socketData)
		} else {
			log.Printf("%s doesn't exist\n", socketData.EventKey)
		}
	}
}
