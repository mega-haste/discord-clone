package socket

import (
	"log"
	"strconv"

	"upchat.com/server/lib/SocketIO"
	"upchat.com/server/model"
)

var SocketServer = SocketIO.New()

func init() {
	SocketServer.OnError(func(err error) bool {
		log.Println("error:", err)
		return false
	})
	SocketServer.OnConnect(func(ID SocketIO.ID_t) error {
		log.Printf("Connected with the id %d\n", ID)
		return nil
	})

	SocketServer.OnDisconnect(func() {
		log.Println("A client diconnected")
	})

	SocketServer.On("chat-send", func(ID SocketIO.ID_t, msg SocketIO.SocketMessageEvent) {
		emt := SocketServer.Except(ID)
		message := model.MessageModel{
			Content: msg.Message,
		}
		if msg.To != "" {
			to_ID, err := strconv.ParseUint(msg.To, 10, 32)
			if err != nil {
				log.Fatalln(err)
				return
			}
			emt.To(SocketIO.ID_t(to_ID)).Emit("receive-message", msg.Message)
			return
		}
		res := model.DB.Select("content").Create(&message)
		if res.Error != nil {
			log.Println("Error: ", res)
		}
		emt.Emit("receive-message", msg.Message)
	})
}
