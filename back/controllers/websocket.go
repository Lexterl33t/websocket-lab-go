package controllers

import (
	"fmt"
	"net/http"

	"labgo/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

	fmt.Println(ws)
}
