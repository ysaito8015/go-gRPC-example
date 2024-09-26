## Server Streaming の実装

- メソッドを追加して
- サーバサイドで実装して
- gRPCurl で動作確認して
- クライアント側を実装して
- 動作確認する


### メソッドの追加

- `service GreetingService {}` に `rpc HelloServerStream (HelloRequest) returns (stream HelloResponse) {}` を追加する

## コードの自動生成

```shell
$ protoc \
--go_out=../pkg/grpc \                 # hello.pg.go の出力先, req/res 型を定義した部分
--go_opt=paths=source_relative \
--go-grpc_out=../pkg/grpc \            # hello_grpc.pb.go の出力先, サービス部分サーバ側の実装部分
--go-grpc_opt=paths=source_relative \
hello.proto
```

`Makefile` にコンパイルする内容を追加

```makefile
.PHONY: compile-proto # Compile the protobuf files
compile-proto:
	protoc proto/*.proto \
		--go_out=. \
		--go_opt=paths=import \
		--go-grpc_out=. \
		--go-grpc_opt=paths=import \
		--proto_path=.
```

### 自動生成されたコードの確認

- サービス側の定義: `../pkg/grpc/hello_grpc.pb.go`
    - `type GreetingServiceServer interface {}` に `HelloServerStream` が追加されている


```go
type GreetingServiceServer interface {
    // サービスが持つメソッド
    Hello(context.Context, *HelloRequest) (*HelloResponse, error)
    // ストリーミング用メソッド
    HelloServerStream(*HelloRequest, GreetingService_HelloServerStreamServer) error

    mustEmbedUnimplementedGreetingServiceServer()
}
```


追加されたサーバストリーミングのためのインタフェース


```go
type GreetingService_HelloServerStreamServer interface {
    Send(*HelloResponse) error
    grpc.ServerStream
}
```


### サーバサイドのサーバストリームメソッドの追加


- 構造体 `myServer` 型に `HelloServerStream` メソッドを追加する
- 追加するメソッドは、 サーバ側に追加した interface のシグネチャに従う


```go
    // ./pkg/grpc/hello_grpc.pb.go
    // ストリーミング用メソッド
    HelloServerStream(*HelloRequest, GreetingService_HelloServerStreamServer) error
```


## クライアントコードの実装

`./pkg/grpc/hello_grpc.pb.go` にクライアント側のインタフェースが追加されている

```go
type GreetingServiceClient interface {
    // Server Streaming RPC
    HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreetingService_HelloServerStreamClient, error)
}
```


サーバから送られてくる複数個のレスポンスを受け取るために `GreetingService_HelloServerStreamClient` interface を戻り値に使う


```go
type GreetingService_HelloServerStreamClient interface {
    Recv() (*HelloResponse, error)
    grpc.ClientStream
}
```


### クライアントコードの実装


```go
func main() {
    // snip

    // receive stream from server
	stream, err := client.HelloServerStream(context.Background(), req)

    // snip

    for {
        // receive response from stream
        res, err := stream.Recv()
        // get end of stream
        if errors.Is(err, io.EOF) {
            fmt.Println("all the responses are received")
            break
        }
        if err != nil {
            log.Fatalf("cannot receive: %v", err)
        }
        log.Printf("Greeting: %s", res.GetMessage())
    }
}
```
