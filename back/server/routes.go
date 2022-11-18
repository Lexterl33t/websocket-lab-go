package server

import (
	"fmt"

	"labgo/controllers"

	"github.com/gin-gonic/gin"
)

func RoutesHTTP() {
	fmt.Println("Hello, world")
}

func RoutesWS(http *gin.Engine) {
	http.GET("/ws", controllers.InitWS)
}
