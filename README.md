# vrc-go

VRChat APIの非公式Goクライアントライブラリ

## Features

- シンプルで使いやすいAPI
- Cookieベースのセッション管理
- 2要素認証対応
- 型安全なエラーハンドリング
- 設定可能なHTTPクライアント（プロキシ、タイムアウト、User-Agent）

## Installation

```bash
go get k4na.de/vrc-go
```

## Quick Start

### 基本認証

```go
package main

import (
    "context"
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

    // ユーザー情報取得
    user, err := client.GetCurrentUser(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Logged in as: %s", user.DisplayName)
}
```

### Cookie認証

```go
client, _ := vrchat.NewClient()

// Cookieを読み込み
if err := client.LoadCookies("cookies.json"); err != nil {
    log.Fatal(err)
}

// ユーザー情報取得（Cookie認証）
user, err := client.GetCurrentUser(context.Background())
if err != nil {
    log.Fatal(err)
}
```

### クライアントオプション

```go
client, err := vrchat.NewClient(
    vrchat.WithUserAgent("my-app/1.0"),
    vrchat.WithTimeout(30 * time.Second),
    vrchat.WithProxy("http://proxy.example.com:8080"),
)
```

## Examples

サンプルコードは `examples/` ディレクトリにあります：

```bash
# 基本認証の例
export VRCHAT_USERNAME="your-username"
export VRCHAT_PASSWORD="your-password"
make run-basic-auth

# Cookie認証の例
make run-cookie-auth
```

## API

### 認証

- `Authenticate(ctx, config)` - ユーザー名/パスワードでログイン
- `GetCurrentUser(ctx)` - 現在のユーザー情報を取得
- `Logout(ctx)` - ログアウト

### Cookie管理

- `SaveCookies(path)` - Cookieをファイルに保存
- `LoadCookies(path)` - Cookieをファイルから読み込み

### エラーハンドリング

```go
err := client.Authenticate(ctx, config)
if err != nil {
    if vrchat.IsAuthenticationError(err) {
        log.Println("Invalid credentials")
    } else if vrchat.IsRateLimitError(err) {
        log.Println("Rate limited")
    } else {
        log.Printf("Error: %v", err)
    }
}
```

## Development

### ビルド

```bash
make build
```

### テスト

```bash
make test
```

### コードフォーマット

```bash
make fmt
```

### 静的解析

```bash
make vet
```

## Important Notes

⚠️ **注意事項**

- これはコミュニティ主導のプロジェクトであり、VRChat公式サポートはありません
- Photonサーバーへの直接アクセスは禁止されています
- APIのレート制限を遵守してください
- VRChatの利用規約に従ってください

## References

- [VRChat API Documentation](https://vrchatapi.github.io/)
- [VRChat API Specification](https://github.com/vrchatapi/specification)

## License

Apache-2.0
