package main

import (
	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/node", app.GetNodeInfo)

	return router
}

func main() {
	initRouter().Run()
}
