package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s12i/maicon-fullstack-test/app/jsonrpc"
)

// transaction struct - 用以儲存要交易的設定
type transaction struct {
	From  string `json:"from"`  // 轉出帳號
	To    string `json:"to"`    // 轉入帳號
	Value int    `json:"value"` // 交易金額
}

/*
GetTrxInfo - 取得交易資訊
Go-Ethereum JSON-RPC API：https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionbyhash
*/
func GetTrxInfo(context *gin.Context) {
	// 準備要查詢的交易編碼
	trxHash := []interface{}{context.Param("trxHash")}

	// 送出交易查詢
	response := jsonrpc.ClientRequest("eth_getTransactionByHash", trxHash)

	// 回傳執行結果
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
	// 讀入要執行的交易設定, 透過 FORM POST 傳送上來的 RAW ( JSON )
	var trx transaction
	context.BindJSON(&trx)

	// 準備要查詢的交易編碼
	trxDetail := map[string]interface{}{"from": trx.From, "to": trx.To, "value": trx.Value}
	params := []map[string]interface{}{trxDetail}

	// 送出交易查詢
	response := jsonrpc.ClientRequest("eth_sendTransaction", params)

	// 回傳執行結果
	context.JSON(http.StatusCreated, gin.H{
		"trx": response.Get("result").MustString(),
	})
}
