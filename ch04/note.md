## proto ファイルとコード生成

## 必要なこと
- 呼び出し元
    - クライアント
    - POST `HelloRequest`
- 呼び出し先
    - サーバ
    - POST `HelloResponse`
- Request, Response をどうやって作るかのビジネスロジック
- Protocol Buffers で serialize and deserialize する
- Protocol Buffers Compiler
    - https://github.com/protocolbuffers/protobuf
    - installation https://grpc.io/docs/protoc-installation/
    - compile from source https://qiita.com/moritaoy/items/30c695fa39943831e17b

```shell
$ sudo apt install autoconf automake libtool curl make g++ unzip -y
$ git clone https://github.com/google/protobuf.git
$ cd protobuf
$ git submodule update --init --recursive
$ cmake -Dprotobuf_BUILD_SHARED_LIBS=ON .
$ cmake --build . --parallel 8
$ sudo cmake --install .
```

```shell
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v28.2/protoc-28.2-linux-x86_64.zip .
$ unzip protoc-28.2-linux-x86_64.zip -d protoc
$ sudo mv protoc/bin/protoc /usr/local/bin/
$ sudo mv protoc/include /usr/local/
```


## protoc-gen-go-grpc, protoc-gen-go
- https://zenn.dev/mikankitten/articles/0437fa6fb7de82


```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@lates
```

## generate code

```shell
$ cd ./api
$ protoc \
  --go_out=../pkg/grpc \                # hello.pb.go の出力先
  --go_opt=paths=source_relative \      # hello.pb.go 生成時のオプション
  --go-grpc_out=../pkg/grpc \           # hello_grpc.pb.go の出力先
  --go-grpc_opt=paths=source_relative \ # hello_grpc.pb.go 生成時のオプション
  hello.proto
```
