package wshub

import "github.com/gorilla/websocket"

type WSHubWrapper struct {
	wsstation *WSHub
}

func New() *WSHubWrapper {
	return &WSHubWrapper{
		wsstation: &WSHub{},
	}
}

func (w *WSHubWrapper) AddAction(state int8, action func(conn *websocket.Conn) error) *WSHubWrapper {
	w.wsstation.mutex.Lock()
	defer w.wsstation.mutex.Unlock()
	w.wsstation.actions[state] = action
	return w
}

func (w *WSHubWrapper) SetOnConnect(state int8) *WSHubWrapper {
	w.wsstation.mutex.Lock()
	defer w.wsstation.mutex.Unlock()
	w.wsstation.onConnect = &state
	return w
}

func (w *WSHubWrapper) SetErrorHandler(handler func(error)) *WSHubWrapper {
	w.wsstation.errorHandler = handler
	return w
}

func (w *WSHubWrapper) Unwarp() *WSHub {
	return w.wsstation
}
