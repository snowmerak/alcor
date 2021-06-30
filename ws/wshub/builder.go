package wshub

import "github.com/gorilla/websocket"

type WSStationWrapper struct {
	wsstation *WSStation
}

func New() *WSStationWrapper {
	return &WSStationWrapper{
		wsstation: &WSStation{},
	}
}

func (w *WSStationWrapper) AddAction(state int8, action func(conn *websocket.Conn) error) *WSStationWrapper {
	w.wsstation.mutex.Lock()
	defer w.wsstation.mutex.Unlock()
	w.wsstation.actions[state] = action
	return w
}

func (w *WSStationWrapper) SetOnConnect(state int8) *WSStationWrapper {
	w.wsstation.mutex.Lock()
	defer w.wsstation.mutex.Unlock()
	w.wsstation.onConnect = &state
	return w
}

func (w *WSStationWrapper) Unwarp() *WSStation {
	return w.wsstation
}
