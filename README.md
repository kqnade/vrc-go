# vrc-go

VRChat APIの非公式Goクライアントライブラリ

## Features

- OpenAPI 3.0仕様ベースの自動生成コード
- 型安全なAPIクライアント
- Cookieベースの認証管理
- 2要素認証対応（予定）
- レート制限対応

## Installation

```bash
go get k4na.de/vrc-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "k4na.de/vrc-go/vrchat"
)

func main() {
    client, err := vrchat.NewClient()
    if err != nil {
        log.Fatal(err)
    }

    // 認証
    err = client.Authenticate(context.Background(), vrchat.AuthConfig{
        Username: "your-username",
        Password: "your-password",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Authentication successful!")
}
```

## Documentation

- [API Documentation](https://pkg.go.dev/k4na.de/vrc-go)
- [Examples](./examples)
- [VRChat API Specification](https://github.com/vrchatapi/specification)

## Important Notes

⚠️ これはコミュニティ主導のプロジェクトであり、VRChat公式サポートはありません。

- Photonサーバーへの直接アクセスは禁止されています
- APIのレート制限を遵守してください
- VRChatの利用規約に従ってください

## Development

### コード生成

```bash
# API仕様取得
make fetch-spec

# コード生成
make generate
```

### テスト

```bash
make test
```

## License

Apache-2.0
