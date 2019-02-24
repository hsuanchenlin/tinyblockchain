Tinyblockchain
===
## Usage
touch state/state.json
put the following text
```$xslt
{"Accs":{},"Nonce":0}
```

There might be path problem, I use Goland and set the content root to this proj not src folder.
```
go run src/main/chain.go
go run src/main/server.go
go run src/main/client.go 
go run src/main/client0.go 
```
## Design
## Additional
利用mine裡面的miner的mine(goroutine)進行挖礦
並調整難度至區塊產生時間

## Ref
ethereum
go
