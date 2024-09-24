## RPC とは、gRPC とは
- RPC: Remote Procedure Call


```go
package main

import ("fmt")

func main() {
    res := hello("world")
    fmt.Println(res)
}

func hello(name string) string {
    return fmt.Sprintf("hello, %s!", name)
}
```

- `main()` から `hello()` を呼び出している
    - Procedure (`hello()`) Call
- これは local でおこなわれている
- Remote Procedure Call は、呼び出し元と呼び出し先が別の位置にある


## gRPC が用いる技術
- HTTP/2
- Protocol Buffers


### HTTP/2
- 公式ドキュメント
    - https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
- Call する関数: リクエストパス
    - POST /helloworld/hello
- Call する関数の引数: リクエストボディ
    - body: "world"
- Call する関数の戻り値: レスポンスボディ
    - header: 200 OK, body: "hello, world!"


### Protocol Buffers
- Request body, Response body の内容をシリアライズするための仕組み
