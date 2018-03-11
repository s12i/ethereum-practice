# MaiCoin Fullstack Engineer Assignment

專案的目的在實作 [MaiCoin Fullstack Engineer Assignment](https://www.evernote.com/l/ADpe8SRPT4tM36027Y1MiOUWN5t8M2vzPfQ) 裡所提及的需求。

## 安裝

專案是透過 [Gin Web Framework](https://github.com/gin-gonic/gin) 開發，並配合 [Dep](https://github.com/golang/dep) 確保使用的相關套件能夠正確的安裝。假設欲執行專案的環境裡已經有安裝了 [Dep](https://github.com/golang/dep)，請進入專案目錄後，執行命令：

```
$ dep ensure
```

完成安裝後，可以在專案下發現多出一個 `vendor` 目錄，裡面放置了專案所需的相關第三方套件：

```
~/go/src/github.com/s12i/maicon-fullstack-test master*
❯ ll
total 32
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
~/go/src/github.com/s12i/maicon-fullstack-test master*
❯ go run server.go
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

如訊息所顯示，專案提供了 6 個 API 供存取區塊鏈，各 API 的使用方式將在下方敘述。