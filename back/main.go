package main

import (
	"fmt"

	"labgo/env"

	"labgo/server"

	"github.com/gin-gonic/gin"
)

func main() {
	http := gin.Default()

	server.RoutesWS(http)
	http.Run(fmt.Sprintf(":%v", env.PORT))
}
