## クライアント側からのストリーミング
- proto ファイルに定義する
- サーバサイドを実装する
- クライアントサイドを実装する

## proto ファイルに追加する

```proto
service GreetingService {
    // snip
    // client streaing RPC
    rpc HelloClientStream(stream HelloRequest) returns (HelloResponse); 
}
```


```shell
$ cd ./src && make complie-proto
```


## サーバサイドの実装

自動生成されたコード


```go
type GreetingServiceServer struct {
    // snip

    // client streaming RPC
    HelloClientStream(GreetingService_HelloClientStreamServer) error
}
```


サーバ側でクライアントストリーミングを受け取る interface


```go
type GreetingService_HelloClientStreamServer interface {
    Send(*HelloResponse) error
    Recv() (*HelloRequest, error)
    grpc.ServerStream
}
```


サーバサイドの実装

リクエストに含まれている構造体 `HelloRequest` 型の受け取り方


```go
// 引数で直接参照していない
func (s *myServer) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
    // snip

    // get HelloRequest struct in for loop
	for {
        // snip
        // get HelloRequest struct
        req, err := stream.Recv()

        // end of stream
        if err == io.EOF {
            // snip

            // return response
			return stream.SendAndClose(&hellopb.HelloResponse{
        }
    }
}
```


サーバサイドの動作確認


```shell
$ grpcurl --plaintext -d \
'{"name": "world"}{"name": "foo"}{"name": "bar"}{"name": "baz"}' \
localhost:8080 \
myapp.GreetingService.HelloClientStream
```


## クライアントサイドの実装

自動生成されたコード


`./pkg/grpc/hello_grpc.pb.go`


```go
type GreetingServiceClient interface {
    // snip
    HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetingService_HelloClientStreamClient, error)
}
```


クライアント側の client streaming RPC のための interface 


```go
type GreetingService_HelloClientStreamClient interface {
    Send(*HelloRequest) error
    CloseAndRecv() (*HelloResponse, error)
    grpc.ClientStream
}
```


実装



