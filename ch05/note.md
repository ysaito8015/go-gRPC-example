## gRPC サーバを作るよ
- `./cmd/server` につくってく
- grpc パッケージの `Server{}` を作る `NewServer()` 
    - https://pkg.go.dev/google.golang.org/grpc#NewServer
- service を登録して、サービスを動かす


## 動作確認, gRPCurl でリクエストを送る
- https://github.com/fullstorydev/grpcurl


```shell
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

```shell
$ cd ./cmd/server
$ go run main.go
```


- サービスの一覧を取得


```shell
$ grpcurl -plaintext localhost:8080 list
```


- サービスのメソッド一覧を取得


```shell
$ grpcurl -plaintext localhost:8080 list myapp.GreetingService
```


- サービスのメソッドを呼び出す


```shell
$ grpcurl -plaintext -d '{"name": "world"}' localhost:8080 myapp.GreetingService.Hello
