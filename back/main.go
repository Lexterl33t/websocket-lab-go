package main

import (
	"fmt"

	"labgo/controllers"
	"labgo/env"
	"labgo/models"

	"labgo/server"

	"github.com/gin-gonic/gin"
)

func init() {
	models.Pool = controllers.InitPool()

	go controllers.RunPool(models.Pool)

}

func main() {
	http := gin.Default()

	server.RoutesWS(http)
	http.Run(fmt.Sprintf(":%v", env.PORT))
}
