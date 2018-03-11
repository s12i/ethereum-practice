package jsonrpc

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/rpc/v2/json2"
)

const (
	rpcURL  string = "localhost"
	rpcPort int    = 8545
)

func ClientRequest(method string, params interface{}) *simplejson.Json {
	data, _ := json2.EncodeClientRequest(method, params)
	request, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:%d", rpcURL, rpcPort), bytes.NewBuffer(data))

	request.Header.Set("Content-Type", "application/json")
	reply, err := http.DefaultClient.Do(request)
	defer reply.Body.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reply.Body)
	response, err := simplejson.NewJson(buffer.Bytes())
	if err != nil {
		fmt.Println("parse JSON fail")
	}

	return response
}
