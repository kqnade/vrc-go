# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

vrc-goはVRChat API用の非公式Goクライアントライブラリです。Cookie認証、2要素認証、WebSocketによるリアルタイムイベント受信をサポートしています。

**重要な制約:**
- VRChatの公式サポートを受けていないコミュニティプロジェクト
- Photonサーバーへの直接アクセスは禁止
- APIレート制限を遵守する必要がある
- VRChatの利用規約に準拠すること

## Development Commands

```bash
# ビルド
make build              # すべてのパッケージをビルド
go build ./...          # 同上

# テスト
make test               # レースディテクター有効でテスト実行
go test -v -race -cover ./...

# コード品質
make fmt                # gofmtでフォーマット
make vet                # go vetで静的解析
make lint               # golangci-lintで詳細チェック

# 依存関係管理
make tidy               # go.modを整理

# サンプル実行
make run-basic-auth     # 基本認証の例 (VRCHAT_USERNAME, VRCHAT_PASSWORD必要)
make run-cookie-auth    # Cookie認証の例
make run-websocket      # WebSocketイベント受信の例

# クリーンアップ
make clean              # ビルド成果物とcookies.jsonを削除
```

## Architecture

### Core Components

**Client (`vrchat/client.go`):**
- HTTPクライアントのラッパー
- Cookie jarによるセッション管理 (`cookiejar.New` with publicsuffix)
- すべてのAPI操作の基盤となる `doRequest()` と `doRequestWithBasicAuth()` メソッド
- デフォルトのベースURL: `https://api.vrchat.cloud/api/1`

**Authentication Flow (`vrchat/auth.go`):**
1. `Authenticate()` - Basic認証でログイン
2. 2FAが必要な場合: `Verify2FA()` または `VerifyEmailOTP()` で完了
3. 認証後、CookieがHTTPクライアントのjarに保存される
4. Cookie永続化: `SaveCookies()` / `LoadCookies()` でJSON形式で保存/読み込み

**WebSocket Architecture (`vrchat/websocket.go`):**
- `ConnectWebSocket()` で接続確立 (authCookieをクエリパラメータとして送信)
- イベント駆動型: `On()`, `OnNotification()`, `OnFriendOnline()` 等でハンドラー登録
- 自動再接続機能 (指数バックオフ: 5秒から最大60秒)
- ワイルドカード `"*"` ハンドラーですべてのイベントをキャッチ可能
- `Wait()` でメインゴルーチンをブロック

### API Categories

APIは機能別に分割されています:

| File | Responsibility |
|------|----------------|
| `auth.go` | 認証、ログアウト、2FA |
| `users.go` | ユーザー検索、プロフィール取得/更新 |
| `friends.go` | フレンドリスト、リクエスト送信/承認/拒否 |
| `avatars.go` | アバター検索、作成、装着、削除 |
| `worlds.go` | ワールド検索、公開/非公開、メタデータ |
| `instances.go` | インスタンス情報、招待送信、作成/クローズ |
| `notifications.go` | 通知取得、既読、削除、招待応答 |
| `favorites.go` | お気に入り追加/削除、グループ管理 |
| `groups.go` | グループ検索、参加/脱退、メンバー管理、BAN |
| `files.go` | ファイルアップロード、バージョン管理、削除 |
| `player_moderation.go` | ミュート、ブロック、アバター非表示 |
| `system.go` | システム設定、ヘルスチェック、オンラインユーザー数 |

### Error Handling (`vrchat/errors.go`)

`APIError` 型を使用し、HTTPステータスコードとメッセージを保持:
- `IsAuthenticationError()` - 401エラー
- `IsRateLimitError()` - 429エラー
- `IsNotFoundError()` - 404エラー
- すべてのエラーは `*APIError` 型として返される

### Options Pattern (`vrchat/options.go`)

`NewClient()` はfunctional options パターンを使用:
- `WithUserAgent(ua string)` - User-Agentをカスタマイズ
- `WithTimeout(timeout time.Duration)` - HTTPタイムアウト設定
- `WithProxy(proxyURL string)` - プロキシ設定

## Key Implementation Details

### Cookie Persistence

`SaveCookies()` と `LoadCookies()` は `vrchat/cookie.go` で実装:
- Cookieを `[]CookieData` 構造体のJSON配列として保存
- 再認証なしでセッションを維持可能

### WebSocket Events

イベント型は `Event` 構造体の `Type` フィールドで識別:
```go
type Event struct {
    Type    string          `json:"type"`
    Content json.RawMessage `json:"content"`
}
```

主要なイベント型:
- `notification`, `friend-online`, `friend-offline`, `friend-location`
- `friend-add`, `friend-delete`, `user-update`
- `group-joined`, `group-left`, `notification-v2`

### Context Usage

すべてのAPI呼び出しは `context.Context` を第一引数に取り、タイムアウトとキャンセルをサポート。

## Module Information

- Module path: `k4na.de/vrc-go`
- Go version: 1.25.5
- External dependencies:
  - `golang.org/x/net` (publicsuffix, HTTP/2サポート)
  - `github.com/gorilla/websocket` (WebSocket接続)
