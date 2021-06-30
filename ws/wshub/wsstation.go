package wshub

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WSStation struct {
	actions   map[int8]func(conn *websocket.Conn) error
	sessions  Sessions
	mutex     sync.RWMutex
	onConnect *int8
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (w *WSStation) Open(rw http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	signalChan := make(chan Signal, 8)
	w.mutex.Lock()
	w.sessions.Add(signalChan)
	defer func(signalChan chan<- Signal) {
		close(signalChan)
		w.sessions.Delete(signalChan)
	}(signalChan)
	w.mutex.RLock()

	if w.onConnect != nil {
		if err := w.actions[*w.onConnect](conn); err != nil {
			return err
		}
	}

	for signal := range signalChan {
		if err := w.actions[signal.state](conn); err != nil {
			return err
		}
	}
	return nil
}
