## proto ファイル
- 関数の呼び出しの定義ファイル


## proto ファイルの書き方
```go
// proto のバージョン
syntax = "proto3";

// proto ファイルから生成される Go のパッケージ名
option go_package = "pkg/grpc";

// declear package name
// 他の proto ファイルで定義された型を使って記述できる
// myapp.HelloRequest
package myapp;

// define service
// arrange methods
service GreetingService {
    // define method
    // (Procedure)
    rpc Hello(HelloRequest)) returns (HelloResponse);
}

// define type
// Procedure name + Request or Response
message HelloRequest {
    string name = 1;
}

message HelloResponse {
    string message = 1;
}
```


## proto ファイル内で定義できる型
### 参考情報
- proto3 の型情報と言語の変換表 https://protobuf.dev/programming-guides/proto3/#scalar
- https://nextpublishing.jp/book/11746.html
- Well-Known types https://developers.google.com/protocol-buffers/docs/reference/google.protobuf


```go
// use google.protobuf.Timestamp
import "google/protobuf/timestamp.proto";

message MyMessage {
    string message = 1;
    google.protobuf.Timestamp create_time = 2;
}


