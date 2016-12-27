// Package websocket implements a WebSocket server by executing
// a command and piping its input and output through the WebSocket
// connection.
package cgi

import (
	"fmt"
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("cgi", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	fmt.Println("Woot !!")
	return nil
}

/*
type MyHandler struct {
	Next httpserver.Handler
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Println("Woot !!")
	return h.Next.ServeHTTP(w, r)
}
*/
