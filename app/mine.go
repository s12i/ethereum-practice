package app

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app/jsonrpc"
)

func StartMine(context *gin.Context) {
	if !govalidator.IsInt(context.Param("threads")) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "threads must be an number",
		})

		return
	}

	threads, _ := strconv.Atoi(context.Param("threads"))
	if threads < 1 {
		threads = 1
	}

	params := []interface{}{threads}
	jsonrpc.ClientRequest("miner_start", params)

	context.Data(http.StatusNoContent, gin.MIMEHTML, nil)
}

func StopMine(context *gin.Context) {
	jsonrpc.ClientRequest("miner_stop", nil)
	context.Data(http.StatusNoContent, gin.MIMEHTML, nil)
}
