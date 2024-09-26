## 双方向ストリーミング RPC
- proto ファイルに追加
- サーバー側の実装
- クライアント側の実装


## proto ファイルに追加
```proto
service GreetingService {
    // snip

    // bidirectional streaming RPC
    // it has stream keyword in both request and response
    rpc HelloBiStream(stream HelloRequest) returns (stream HelloResponse);

}
```


```shell
$ make compile-proto
```


## サーバー側の実装
自動生成されたコード


```go
type GreetingServiceServer interface {
    // snip

    // bidirectional streaming RPC
    HelloBiStreams(GreetingService_HelloBiStreamsServer) error
    //snip
}
```


データのやり取りを行う interface 部分


```go
type GreetingService_HelloBiStreamsServer interface {
    Send(*HelloResponse) error
    Recv() (*HelloRequest, error)
    grpc.ServerStream
}
```


サーバサイドの動作確認


```shell
$ grpcurl -plaintext -d \
'{"name": "world"}{"name": "foo"}{"name": "bar"}{"name": "baz"}' \
localhost:8080 \
myapp.GreetingService.HelloBiStreams
```


## クライアント側の実装


