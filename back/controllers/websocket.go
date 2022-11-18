package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"labgo/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ClientExtended_t struct {
	Client_t *models.Client_t
}

/*
	var whitelist_origin = map[string]bool{
		"https://google.fr":      true,
		"https://facebook.com":   true,
		"http://127.0.0.1:5500/": true,
	}
*/
func IdentifyOrigin(http_req *http.Request) bool {
	/*
		origin := http_req.Header.Get("Origin")
		fmt.Println(origin)
		if whitelist_origin[origin] {
			return true
		} else {
			return false
		}*/

	return true
}

func GeneratePoolID() uint32 {
	return rand.Uint32()
}

func ClientJoin(conn *websocket.Conn, pool *models.Pool_t) *ClientExtended_t {
	client := &models.Client_t{
		Username: fmt.Sprintf("Bryton-%v", GeneratePoolID()),
		Conn:     conn,
		Pool:     pool,
	}

	pool.Register <- client

	clientExtended := ClientExtended_t{
		Client_t: client,
	}

	return &clientExtended
}

func (c *ClientExtended_t) ClientExit() {
	models.Pool.Unregister <- c.Client_t
	c.SendListConnectedUsers()
}

func (c *ClientExtended_t) ListenerClientMessage() {
	defer c.ClientExit()

	for {
		var command models.Command_t
		err := c.Client_t.Conn.ReadJSON(&command)
		if err != nil {
			log.Println(err)
			return
		}

		c.Client_t.Pool.Commands <- command
	}
}

func (client_d *ClientExtended_t) SendListConnectedUsers() {

	users := make([]string, len(client_d.Client_t.Pool.Clients))
	for client, _ := range client_d.Client_t.Pool.Clients {
		if client != client_d.Client_t {
			users = append(users, client.Username)
		}
	}

	for client, _ := range client_d.Client_t.Pool.Clients {
		if client != client_d.Client_t {
			client.Conn.WriteJSON(models.Command_t{
				ID:  0x3,
				CMD: users,
			})
		}
	}
}

func (client_d *ClientExtended_t) NewUserJoinMessage(message string) {
	for client, _ := range client_d.Client_t.Pool.Clients {
		if client != client_d.Client_t {
			client.Conn.WriteJSON(models.Command_t{
				ID: 0x1,
				CMD: models.UserJoin_t{
					Msg: message,
				},
			})
		}
	}
}

func (client_d *ClientExtended_t) UserExitServerMessage(message string) {
	for client, _ := range client_d.Client_t.Pool.Clients {
		if client != client_d.Client_t {
			client.Conn.WriteJSON(models.Command_t{
				ID: 0x2,
				CMD: models.UserExit_t{
					Msg: message,
				},
			})
		}
	}
}

func RunPool(p *models.Pool_t) {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true

			client_extended := ClientExtended_t{Client_t: client}

			client_extended.NewUserJoinMessage(
				fmt.Sprintf("%v à rejoins le serveur", client.Username),
			)

		case client := <-p.Unregister:
			delete(p.Clients, client)
			client_extended := ClientExtended_t{Client_t: client}

			client_extended.UserExitServerMessage(
				fmt.Sprintf("%v à quitté le serveur", client.Username),
			)
		}
	}
}

func InitPool() *models.Pool_t {
	return &models.Pool_t{
		ID:         uint(GeneratePoolID()),
		Clients:    make(map[*models.Client_t]bool),
		Register:   make(chan *models.Client_t),
		Unregister: make(chan *models.Client_t),
		Commands:   make(chan models.Command_t),
	}
}

func InitWS(gin_ctx *gin.Context) {
	upgrader.CheckOrigin = IdentifyOrigin
	ws, err := upgrader.Upgrade(gin_ctx.Writer, gin_ctx.Request, nil)
	if err != nil {
		gin_ctx.JSON(403, models.Error{
			Status:  models.UPGRADER_CONN,
			Message: "Error upgrade websocket connection",
		})
		return
	}

	client := ClientJoin(ws, models.Pool)

	client.ListenerClientMessage()
}
