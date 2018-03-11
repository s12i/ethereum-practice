package main

import (
	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	// 取得節點資訊
	router.GET("/node", app.GetNodeInfo)

	// 透過區塊編號取得區塊資訊
	router.GET("/block/:blockNumber", app.GetBlockInfo)

	// 透過交易編號取得交易資訊
	router.GET("/trx/:trxHash", app.GetTrxInfo)

	// 送出交易
	router.POST("/trx", app.SendTrx)

	return router
}

func main() {
	initRouter().Run()
}
