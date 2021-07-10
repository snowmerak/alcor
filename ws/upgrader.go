package ws

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

var Upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(ctx *fasthttp.RequestCtx) bool { return true },
}
