package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app/jsonrpc"
)

type transaction struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}

/*
GetTrxInfo - 取得交易資訊
Go-Ethereum JSON-RPC API：https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionbyhash
*/
func GetTrxInfo(context *gin.Context) {
	trxHash := []interface{}{context.Param("trxHash")}

	response := jsonrpc.ClientRequest("eth_getTransactionByHash", trxHash)

	context.JSON(http.StatusOK, gin.H{
		"blockHash":   response.Get("result").Get("blockHash").MustString(),
		"blockNumber": response.Get("result").Get("blockNumber").MustString(),
		"from":        response.Get("result").Get("from").MustString(),
		"gas":         response.Get("result").Get("gas").MustString(),
		"gasPrice":    response.Get("result").Get("gasPrice").MustString(),
		"hash":        response.Get("result").Get("hash").MustString(),
		"nonce":       response.Get("result").Get("nonce").MustString(),
		"to":          response.Get("result").Get("to").MustString(),
		"value":       response.Get("result").Get("value").MustString(),
	})
}

/*
SendTrx - 送出交易
Go-Ethereum JSON-RPC API：https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction
*/
func SendTrx(context *gin.Context) {
	var trx transaction
	context.BindJSON(&trx)

	trxDetail := map[string]interface{}{"from": trx.From, "to": trx.To, "value": trx.Value}
	params := []map[string]interface{}{trxDetail}
	response := jsonrpc.ClientRequest("eth_sendTransaction", params)

	context.JSON(http.StatusCreated, gin.H{
		"trx": response.Get("result").MustString(),
	})
}
