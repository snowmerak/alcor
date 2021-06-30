package wshub

import "github.com/gorilla/websocket"

type WSStationWrapper struct {
	wsstation *WSHub
}

func New() *WSStationWrapper {
	return &WSStationWrapper{
		wsstation: &WSHub{},
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

func (w *WSStationWrapper) Unwarp() *WSHub {
	return w.wsstation
}
