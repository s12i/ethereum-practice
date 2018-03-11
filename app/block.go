package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app/jsonrpc"
)

type ethGetBlockByNumber struct {
	string
	bool
}

/*
GetBlockInfo - 透過區塊編號取得區塊資訊
Go-Ethereum JSON-RPC API：https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbynumber
*/
func GetBlockInfo(context *gin.Context) {
	if !govalidator.IsInt(context.Param("blockNumber")) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "blockNumber must be an number",
		})

		return
	}

	blockNumber, _ := strconv.Atoi(context.Param("blockNumber"))
	params := []interface{}{fmt.Sprintf("0x%x", blockNumber), true}
	response := jsonrpc.ClientRequest("eth_getBlockByNumber", params)

	context.JSON(http.StatusOK, gin.H{
		"difficulty":      response.Get("result").Get("difficulty").MustString(),
		"gasLimit":        response.Get("result").Get("gasLimit").MustString(),
		"gasUsed":         response.Get("result").Get("gasUsed").MustString(),
		"hash":            response.Get("result").Get("hash").MustString(),
		"miner":           response.Get("result").Get("miner").MustString(),
		"parentHash":      response.Get("result").Get("parentHash").MustString(),
		"totalDifficulty": response.Get("result").Get("totalDifficulty").MustString(),
	})
}
