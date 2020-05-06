package routine

import (
	"log"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

var client *socketio_client.Client
var err error

func NewSocketIO(uri string) *socketio_client.Client {
	client, err = socketio_client.NewClient(uri, nil)

	if err != nil {
		log.Printf("Socket IO error:%v\n", err)
	}

	return client
}