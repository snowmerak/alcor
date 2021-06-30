package wsstream

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSStream struct {
	actions map[int8]func(conn *websocket.Conn, typ int, message []byte) (int8, error)
	mutex   sync.RWMutex
}

func New() WSStream {
	return WSStream{
		actions: map[int8]func(conn *websocket.Conn, typ int, message []byte) (int8, error){},
		mutex:   sync.RWMutex{},
	}
}

func (w *WSStream) AddAction(state int8, action func(conn *websocket.Conn, typ int, message []byte) (int8, error)) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.actions[state] = action
}

func (w *WSStream) Connect(url string) error {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	state := int8(0)
	for {
		typ, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		if fn, ok := w.actions[state]; ok {
			next, err := fn(conn, typ, message)
			if err != nil {
				return err
			}
			state = next
		}
	}
}
