package api

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/shamaton/msgpack"
)

type APIClient struct {
	conn *websocket.Conn
}

func (client *APIClient) Hold() {
	for {
		client.receive()
	}
}

func (client *APIClient) receive() {
	mt, message, err := client.conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	opmsg := OpMessage{}
	if err = msgpack.Decode(message, &opmsg); err != nil {
		log.Println("msgpack:", err)
		return
	}

	log.Println(opmsg.Op)

	// TODO: use database
	switch opmsg.Op {
	case OpHello:
		break
	case OpReqAuth:
		client.SendAuthComplete(opmsg.Cb, &User{1, "Mempler", 6666, UserStatusActive})
		break
	case OpReqServerList:
		client.SendServerList(opmsg.Cb, []*Server{&Server{ID: 0}})
		break
	case OpReqUser:
		client.SendUser(opmsg.Cb, &User{1, "Mempler", 6666, UserStatusActive})
		break
	case OpReqServer:
		client.SendServer(opmsg.Cb, &Server{
			ID:   0,
			Name: "Sample Server",

			Categories: []int64{0},
			Users:      []int64{0},

			Logo: 0,
		})
		break
	case OpReqCategory:
		client.SendCategory(opmsg.Cb, &Category{
			ID:       0,
			Name:     "Information",
			Channels: []int64{0},
		})
	case OpReqChannel:
		client.SendChannel(opmsg.Cb, &Channel{
			ID:           0,
			Name:         "announcements",
			MessageCount: 0,
			MessageHistory: []Message{Message{
				ID:        0,
				Message:   "Hello World!",
				Timestamp: 0,
				User:      0,
			}},
			MessagesRead:        0,
			Topic:               "No Topic",
			VoiceConnectedUsers: []int64{0},
		})
	default:
		break
	}

	log.Printf("recv: %s", message)
	err = client.conn.WriteMessage(mt, message)
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func (client *APIClient) Close() {
	client.conn.Close()
}

func (client *APIClient) Send(op int, cb int, data interface{}) {
	d, err := msgpack.Encode(data)
	if err != nil {
		log.Println("send:", err)
		return
	}

	opMsg, err := msgpack.Encode(OpMessage{
		Op:   op,
		Cb:   cb,
		Data: d,
	})
	if err != nil {
		log.Println("send:", err)
		return
	}

	client.conn.WriteMessage(websocket.BinaryMessage, opMsg)
}

func CreateAPIClient(conn *websocket.Conn) *APIClient {
	return &APIClient{conn}
}
