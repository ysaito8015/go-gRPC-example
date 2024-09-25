## gRPC で可能な通信方法
1. Unary RPC
2. Server Streaming RPC
3. Client Streaming RPC
4. Bidirectional Streaming RPC


### Unary RPC
- 1つのリクエストと1つのレスポンスを送受信する


### Server Streaming RPC
- 1つのリクエストと複数のレスポンスを送受信する


### Client Streaming RPC
- 複数のリクエストと1つのレスポンスを送受信する


### Bidirectional Streaming RPC
- 複数のリクエストと複数のレスポンスを送受信する
- WebSocket のような通信が可能


### gRPC の通信を支える技術
- HTTP/2
- Frame
    - Type (8bit)
        - フレームの種類を格納
            - DATA
                - req/res の body
            - HEADERS
                - req/res の header
    - Flags (8bit)
        - フレームのフラグを格納
            - END_STREAM
                - フレームの終了を示す
            - END_HEADERS
                - ヘッダーの終了を示す
