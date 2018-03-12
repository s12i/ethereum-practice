# MaiCoin Fullstack Engineer Assignment

專案的目的在實作 [MaiCoin Fullstack Engineer Assignment](https://www.evernote.com/l/ADpe8SRPT4tM36027Y1MiOUWN5t8M2vzPfQ) 裡所提及的需求。

## 安裝

專案是透過 [Gin Web Framework](https://github.com/gin-gonic/gin) 開發，並配合 [Dep](https://github.com/golang/dep) 確保使用的相關套件能夠正確的安裝。假設欲執行專案的環境裡已經有安裝了 [Dep](https://github.com/golang/dep)，請進入專案目錄後 ( $HOME/go/src/github.com/s12i/maicon-fullstack-test )，執行命令：

```
$ dep ensure
```

完成安裝後，可以在專案下發現多出一個 `vendor` 目錄，裡面放置了專案所需的相關第三方套件：

```
-rw-r--r--  1 s12i  staff   2.2K  3 12 01:21 Gopkg.lock
-rw-r--r--  1 s12i  staff   738B  3 11 12:04 Gopkg.toml
-rw-r--r--@ 1 s12i  staff   546B  3 12 01:20 README.md
drwxr-xr-x  8 s12i  staff   256B  3 11 23:50 app
-rw-r--r--  1 s12i  staff   747B  3 12 00:55 server.go
drwxr-xr-x  5 s12i  staff   160B  3 12 01:21 vendor
```

爾後，執行下述命令啟動專案：

```
$ go run server.go
```

或者，直接執行預先編譯產生的執行檔 ( Mac 環境 )：

```
$ ./server
```

若執行成功，可以看見如下的訊息：

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /node                     --> github.com/s12i/maicon-fullstack-test/app.GetNodeInfo (4 handlers)
[GIN-debug] GET    /block/:blockNumber       --> github.com/s12i/maicon-fullstack-test/app.GetBlockInfo (4 handlers)
[GIN-debug] GET    /trx/:trxHash             --> github.com/s12i/maicon-fullstack-test/app.GetTrxInfo (4 handlers)
[GIN-debug] POST   /trx                      --> github.com/s12i/maicon-fullstack-test/app.SendTrx (4 handlers)
[GIN-debug] PUT    /miner/:threads           --> github.com/s12i/maicon-fullstack-test/app.StartMine (4 handlers)
[GIN-debug] DELETE /miner                    --> github.com/s12i/maicon-fullstack-test/app.StopMine (4 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

如訊息所顯示，專案存取位置為 `http://localhost:8080`，並提供了 6 個 API 供存取區塊鏈，各 API 的使用方式將在下方敘述：

* GET /node
* GET /block/:blockNumber
* GET /trx/:trxHash
* POST /trx
* PUT /miner/:threads
* DELETE /miner

另外，為了測試與示範，專案開發過程使用 `testrpc` 來建立乙太坊測試環境，可以透過下述命令安裝所需的工具 ( 需要先安裝 NodeJS )：

```
npm install -g ethereumjs-testrpc
```

安裝完成後，啟動 `testrpc`：

```
$ testrpc
EthereumJS TestRPC v6.0.3 (ganache-core: 2.0.2)

Available Accounts
==================
(0) 0xf9a022e734d573f92943ffd34d203dc16d6dcbcd
(1) 0x6989ff0102d99c804b8b6459cc720b433e5c46c1
(2) 0xd079d212d9c5cddf0e180064b625c76dd213e9ce
(3) 0xaff2ba4ca4e02de13a8f4b71a1ddccd59bb792f9
(4) 0x403047f4f6b05c4334af1e302a2703261e43abce
(5) 0x9bf45892cbef1cf5c2f3fc6ebbae0b6dd3cec0d9
(6) 0xd837e2540a1381c0cae1a8f8ad95db5f11cdb28d
(7) 0x9c3a85c52efa119b9326c1b83ef748f994c17d2a
(8) 0xb3d80df2975100b23135932c782fedaa64328e42
(9) 0x6b3d042e3ac7653b0e6e926572621335182e80e2

Private Keys
==================
(0) c37003d3f0d9ca886b8ebe01e400cf63927c85c772e579b2f780293ad6fd77f9
(1) b3f98d6b254caddb1cbbc1b8791a309910555a075abdb1ca4f888612d26f6eec
(2) 32196ead0f80fa8a8805730f01564ddba969a77d14130929a60551a9a0a3c3df
(3) 55d5f75a5c72155007a606b836cc0e9bfd031a5b1d2d1a34e4255e64b8f5237d
(4) 22fc6511d734e60236f5c3071e52e2ab181c0800b32a6935b526d7d8ed6ac045
(5) e3482a6778338c2fa8a65b30b9261d3b4bd2e2284830650401e4f0ba58f8c05d
(6) 48a70620e59747a8c9e9bd26cbf22c4df63d508c2c05354e215b2ee711b9ac7f
(7) b04ad2976fa5c85e26af1e87f26a3ad0863ba8054aea37fcef3a3d7e365d1cf0
(8) 7dacb0f7b983d07d664802208d469874291474fe53862ff9000832f5ea1c3190
(9) 302c62ffbaee893cf474558586917dbfe5e2b605b675f2ca3729ae71d42a119b

HD Wallet
==================
Mnemonic:      oxygen scrap post cost virus absurd scare shuffle bar blush vote brief
Base HD Path:  m/44'/60'/0'/0/{account_index}

Listening on localhost:8545
eth_sendTransaction
```
可以看見 `testrpc` 已預先建立了 10 組帳號 ( 並包含預設數量的乙太幣 )，稍後說明 API 的使用時，將會以 `testrpc` 所建立的測試帳號做說明。但 `testrpc` 並未包含 `admin` 模組：

```
$ geth attach http://localhost:8545
Welcome to the Geth JavaScript console!

instance: EthereumJS TestRPC/v2.0.2/ethereum-js
coinbase: 0xc78270f5ff9121689a4e6205cd549efe51756fde
at block: 0 (Mon, 12 Mar 2018 01:45:20 CST)
 modules: eth:1.0 evm:1.0 net:1.0 personal:1.0 rpc:1.0 web3:1.0

$ rpc.modules
{
  eth: "1.0",
  evm: "1.0",
  net: "1.0",
  personal: "1.0",
  rpc: "1.0",
  web3: "1.0"
}
```

為了可以讓 API 順利取回節點資訊，會需要存取 `admin` 模組，因此需要實際安裝 [Go-ethereum (geth)](https://github.com/ethereum/go-ethereum/)，假設環境已經安裝了 [Go-ethereum (geth)](https://github.com/ethereum/go-ethereum/)，請執行下述命令啟動：

```
$ geth --identity maicoin-test init genesis.json --datadir privatechain
```
接著產生一位使用者：

```
$ geth account new
INFO [03-12|02:06:05] Maximum peer count                       ETH=25 LES=0 total=25
Your new account is locked with a password. Please give a password. Do not forget this password.
Passphrase:
```

填入密碼後 ( Passphrase )，將回回傳該使用者的編號 ( 編號將與範例說明不同 )：

```
Address: {443b2f62f649929d54ef1c7eaa11948e9901058d}
```

我們將使用該帳號來驗證開始挖礦以及停止挖礦的 API。要啟動 `geth` 可以執行如下命令，將會在專案下產生一個私有區塊資料用的目錄 `privatechain`，接著執行：

```
geth --datadir privatechain \
     --networkid 1116 \
     --rpc --rpcapi admin,db,eth,net,web3,personal,miner \
     --cache=1024 --rpcport 8545 --rpcaddr 127.0.0.1 --rpccorsdomain "*" \
     --etherbase=443b2f62f649929d54ef1c7eaa11948e9901058d
```

請將上述產生的帳號加入 `etherbase`。上述命令中加入了 `admin` 以及 `miner` 模組，是為了提供對應的 `RPC` 呼叫 ( `admin` 以及 `miner` )：

```
$ geth attach http://localhost:8545
Welcome to the Geth JavaScript console!

instance: Geth/v1.8.2-stable/darwin-amd64/go1.10
coinbase: 0x65fa89786c08c2019351ff5f27d92578bc6130a7
at block: 0 (Thu, 01 Jan 1970 08:00:00 CST)
 datadir: /Users/s12i/go/src/github.com/s12i/maicon-fullstack-test/privatechain
 modules: admin:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 web3:1.0
 
$ rpc.modules
{
  admin: "1.0",
  eth: "1.0",
  miner: "1.0",
  net: "1.0",
  personal: "1.0",
  rpc: "1.0",
  web3: "1.0"
} 
```

`geth` 與 `testrpc` 將只能擇一啟動，避免相關 Port ( 8545 ) 使用。

## API

#### GET /node

取得區塊鏈節點資訊 ( 需要啟動 `geth` )

#### 範例

```
// 送出請求
curl -X GET http://localhost:8080/node 

// 回傳結果 ( HTTP 200 )
{
    "enode": "enode://78d20aea232b5a10877f9e2569e1ed72fc4243bff11fecea99f5e63529ba1e0722e16d48e764557a04df2f4f55dc782ae7c77da60d61b400f2bff7b90f0547e2@192.168.0.10:30303",
    "name": "Geth/v1.8.2-stable/darwin-amd64/go1.10"
}
```

#### GET /block/:blockNumber

取得區塊資訊 ( 啟動 `testrpc` 或 `geth` )

#### 參數

區塊編號，須為一個整數

#### 範例

```
// 送出請求
curl -X GET http://localhost:8080/block/0

// 回傳結果 ( HTTP 200 )
{
    "difficulty": "0x0",
    "gasLimit": "0x6691b7",
    "gasUsed": "0x0",
    "hash": "0x9eca50fa8b5b61151e43e13758f1caef6a05e9e65939e65fe57a7ee97113180c",
    "miner": "0x0000000000000000000000000000000000000000",
    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "totalDifficulty": "0x0"
}
```

若送出不為整數的區塊編號：

```
// 送出請求
curl -X GET http://localhost:8080/block/foo


// 回傳結果 ( HTTP 422 )
{
    "message": "blockNumber must be an number"
}
```

#### POST /trx

送出交易 ( 啟動 `testrpc` )

#### 參數

為 `form raw data`：

```
{
	"from": "0xf9a022e734d573f92943ffd34d203dc16d6dcbcd",
	"to": "0x6989ff0102d99c804b8b6459cc720b433e5c46c1",
	"value": 100
}
```

#### 範例

```
// 送出請求
curl -X POST \
  http://localhost:8080/trx \
  -H 'Content-Type: application/json' \
  -d '{
	"from": "0xf9a022e734d573f92943ffd34d203dc16d6dcbcd",
	"to": "0x6989ff0102d99c804b8b6459cc720b433e5c46c1",
	"value": 100
}'


// 回傳結果 ( HTTP 201 )
{
    "trx": "0xb64d57709cc92445f652b32ddfbf54b6d8a78be9d5a19ea2d7ed6c0d2763075b"
}
```

可以在 `testrpc` console 中看見交易訊息：

```
Transaction: 0xb64d57709cc92445f652b32ddfbf54b6d8a78be9d5a19ea2d7ed6c0d2763075b
Gas usage: 21000
Block Number: 1
Block Time: Mon Mar 12 2018 02:14:10 GMT+0800 (CST)
```

#### GET /trx/:trxHash

取回交易資訊 ( 啟動 `testrpc` )

#### 參數

為一串代表交易的編號，以上述產生交易的結果為例，編號將會為 `0xb64d57709cc92445f652b32ddfbf54b6d8a78be9d5a19ea2d7ed6c0d2763075b `

#### 範例

```
// 送出請求
curl -X GET \
  http://localhost:8080/trx/0xb64d57709cc92445f652b32ddfbf54b6d8a78be9d5a19ea2d7ed6c0d2763075b

// 回傳結果 ( HTTP 200 )
{
    "blockHash": "0xd4b0a1490f308b21123ce31e20a0155f5a435c8a2c9b526ea0a85efae1cbef44",
    "blockNumber": "0x01",
    "from": "0xf9a022e734d573f92943ffd34d203dc16d6dcbcd",
    "gas": "0x015f90",
    "gasPrice": "0x01",
    "hash": "0xb64d57709cc92445f652b32ddfbf54b6d8a78be9d5a19ea2d7ed6c0d2763075b",
    "nonce": "0x0",
    "to": "0x6989ff0102d99c804b8b6459cc720b433e5c46c1",
    "value": "0x64"
}
```

#### PUT /miner/:threads

開始挖礦

#### 參數

執行緒 ( threads ) 數目，須為一個整數，若小於 `1`，將預設為 `1`

#### 範例

```
// 送出請求
curl -X PUT http://localhost:8080/miner/3

// 回傳結果 ( HTTP 204 )

// geth console
INFO [03-12|02:23:36] Updated mining threads                   threads=3
INFO [03-12|02:23:36] Transaction pool price threshold updated price=18000000000
INFO [03-12|02:23:36] Starting mining operation
INFO [03-12|02:23:36] Commit new mining work                   number=1 txs=0 uncles=0 elapsed=394.881µs
INFO [03-12|02:23:44] Successfully sealed new block            number=1 hash=07949a…0dd5ae
```

```
// 送出請求，且執行緒數目為 0
curl -X PUT http://localhost:8080/miner/0

// 回傳結果 ( HTTP 204 )

// geth console
INFO [03-12|02:25:08] Updated mining threads                   threads=1
INFO [03-12|02:25:08] Transaction pool price threshold updated price=18000000000
INFO [03-12|02:25:08] Starting mining operation
INFO [03-12|02:25:08] Commit new mining work                   number=55 txs=0 uncles=0 elapsed=429.195µs
INFO [03-12|02:25:09] Successfully sealed new block            number=55 hash=0d55ce…1bfe6f
INFO [03-12|02:25:09] 🔗 block reached canonical chain         number=50 hash=446621…4debc2
INFO [03-12|02:25:09] 🔨 mined potential block                 number=55 hash=0d55ce…1bfe6f
```

#### DELETE /miner

停止挖礦

#### 範例

```
// 送出請求
curl -X DELETE http://localhost:8080/miner

// 回傳結果 ( HTTP 204 )
```






