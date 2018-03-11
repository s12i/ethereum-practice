package app

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app/jsonrpc"
)

/*
StartMine - 開始挖礦
Go-Ethereum JSON-RPC API：https://github.com/ethereum/go-ethereum/wiki/Management-APIs#miner_start
*/
func StartMine(context *gin.Context) {
	// 判斷傳入的是否為合法的執行緒個數 ( 須為整數 )
	if !govalidator.IsInt(context.Param("threads")) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "threads must be an number",
		})

		return
	}

	// 若傳入的執行緒個數小於 1, 則預設為 1
	threads, _ := strconv.Atoi(context.Param("threads"))
	if threads < 1 {
		threads = 1
	}

	// 準備挖礦請求的參數 - 多少個執行緒
	params := []interface{}{threads}

	// 送出開始挖礦請求
	jsonrpc.ClientRequest("miner_start", params)

	// 回傳執行結果
	context.Data(http.StatusNoContent, gin.MIMEHTML, nil)
}

/*
StopMine - 停止挖礦
Go-Ethereum JSON-RPC API：https://github.com/ethereum/go-ethereum/wiki/Management-APIs#miner_stop
*/
func StopMine(context *gin.Context) {
	// 送出停止挖礦請求
	jsonrpc.ClientRequest("miner_stop", nil)

	// 回傳執行結果
	context.Data(http.StatusNoContent, gin.MIMEHTML, nil)
}
