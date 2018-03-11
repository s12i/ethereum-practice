package main

import (
	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/node", app.GetNodeInfo)
	router.GET("/block/:blockNumber", app.GetBlockInfo)

	return router
}

func main() {
	initRouter().Run()
}
