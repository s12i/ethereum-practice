package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app/jsonrpc"
)

func GetNodeInfo(context *gin.Context) {
	response := jsonrpc.ClientRequest("admin_nodeInfo", nil)

	context.JSON(http.StatusOK, gin.H{
		"enode": response.Get("result").Get("enode").MustString(),
		"name":  response.Get("result").Get("name").MustString(),
	})
}
