# MaiCoin Fullstack Engineer Assignment

專案的目的在實作 [MaiCoin Fullstack Engineer Assignment](https://www.evernote.com/l/ADpe8SRPT4tM36027Y1MiOUWN5t8M2vzPfQ) 裡所提及的需求。

## 安裝

專案是透過 [Gin Web Framework](https://github.com/gin-gonic/gin) 開發，並配合 [Dep](https://github.com/golang/dep) 確保使用的相關套件能夠正確的安裝。假設欲執行專案的環境裡已經有安裝了 [Dep](https://github.com/golang/dep)，請進入專案目錄後，執行命令：

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
(0) 0xc78270f5ff9121689a4e6205cd549efe51756fde
(1) 0x0a7c4e52a314798ee5a3c8ec127818786430d78e
(2) 0x9414176541f9fc66c7066442d6163d4980a609b5
(3) 0xae49abf36d168804692ae7d64b0915089ae44277
(4) 0xecd4757dc4ac59ad9b038f750ff36c2f7d118658
(5) 0x0dbe9800a523235e18847ff6e3cac596543f0a02
(6) 0xa029d164e0beae7ae19827731edf672d13251bf0
(7) 0x8835d0c1efb48148ba9af1ce6e0533beded0e815
(8) 0x0d321482c50e9973556bde322a81758df036fb5a
(9) 0x5d7b24450faa783785cf48f2e46ab5dd4a2badf9

Private Keys
==================
(0) 9da80ff9510ee055dbabefb454d91cd423d3aec9fce9fc8b58dc512086c46a1f
(1) 10ea47f4ae33aef6eb4749f3be603ffac04b339b5aaa63d693b41223b5699120
(2) c72fd2f8815c865e6738318451adfda454566c75212b8fd88be96bc53807d25a
(3) 5490250c510ccb39462011e78eb38ef7e507109cbd3cc81750718b4114166322
(4) 34074844e410184062be5719e46ce3602e81d48a9430f020ea74e35ba5c47e7a
(5) 7b26877f4c89b2795aba67608b25c0572c03817ec3c3a2af266b91754b2e8e18
(6) 060e98155ef966aac2f38b6a07046c999eb2021500d85ce14fc0e2556bc6dd3d
(7) 0d9c89ef8dcbe6e61a84098c3e3ff8ab418965677f32c38d7a4bcf801b3b1520
(8) fed85f5ce15c25e4acda4ca280deaab2a1a5607454b3961fbc3bca9d03c908a2
(9) 7bca7b8d2920c5cc465fed1da53634063ea846a8bb35e762bc82617180ec9dc4

HD Wallet
==================
Mnemonic:      wife faculty never hard canoe truth bulk swift balance outdoor script sell
Base HD Path:  m/44'/60'/0'/0/{account_index}

Listening on localhost:8545
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

執行後，將會在專案下產生一個私有區塊資料用的目錄 `privatechain`，接著執行：

```
geth --datadir privatechain \
     --networkid 1116 \
     --rpc --rpcapi admin,db,eth,net,web3,personal,miner \
     --cache=1024 --rpcport 8545 --rpcaddr 127.0.0.1 --rpccorsdomain "*" \
     --etherbase=65fa89786c08c2019351ff5f27d92578bc6130a7
```

上述命令中加入了 `admin` 以及 `miner` 模組，提供對應的 `RPC` 呼叫。`geth` 與 `testrpc` 將只能擇一啟動，避免相關 Port ( 8545 ) 使用。

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




