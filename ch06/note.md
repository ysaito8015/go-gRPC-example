## gRPC クライアントの実装
- gRPC クライアントの実装に必要なのは
    1. `grpc.Dial()` 関数を使ってサーバーに接続する
    2. 渡す引数は、サーバーのアドレスとポート番号、どのようにコネクションを確立するかを指定するオプション
        - `grp.WithTransportCredentials()`
        - `grpc.WithBlock()` コネクションが確立されるまで待機する
    3. `pkg/grpc/hello_grpc.pb.go` で自動生成されているクライアントのコンストラクタをよびだす
        - `func NewGreetingServiceClient() GreetingServiceClient {}`
    4. クライアントからサーバの `Hello` メソッドにリクエストを送るためのメソッド `Hello()` をよびだす
        - `type GreeingServiceClient interface { Hello() (*HelloResponse, error) }`
