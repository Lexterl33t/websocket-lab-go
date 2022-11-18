package models

import (
	"github.com/gorilla/websocket"
)

const (
	UPGRADER_CONN = iota
)

var Pool *Pool_t

type Pool_t struct {
	ID         uint
	Clients    map[*Client_t]bool
	Register   chan *Client_t
	Unregister chan *Client_t
	Commands   chan Command_t
}

type Command_t struct {
	ID  uint
	CMD interface{}
}

type Client_t struct {
	ID       uint
	Username string
	Pool     *Pool_t
	Conn     *websocket.Conn
}

type Message_t struct {
	SendBy string
	To     string
	Msg    string
}

type UserJoin_t struct {
	Msg string
}

type UserExit_t struct {
	Msg string
}
