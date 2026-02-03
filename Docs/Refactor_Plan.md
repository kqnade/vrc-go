# パッケージ構成のリファクタリング計画

## 概要

現在の `vrchat` パッケージは REST API とWebSocket 機能が混在しており、関心事の分離ができていない。これを`api`と`websocket`に分離する。

## 現在の問題点

1. **関心事の混在**
   - `Client` 型：`Authenticate()`, `GetUser()`, `SearchWorlds()` など REST API 12個分のメソッド
   - `Client` 型：`ConnectWebSocket()` ← WebSocket専用メソッド
   - 全く異なる責務が同じ型に混在している

2. **WebSocketClient の依存関係**
   - `WebSocketClient` は `Client` に依存
   - 実際に使用しているのは認証クッキーの取得（`getAuthCookie()`）のみ
   - 不要な依存関係が生じている

3. **スケーラビリティの問題**
   - REST API が増えると `Client` が肥大化
   - WebSocket ハンドラーが複雑化しても同じパッケージで追跡困難
   - テスト時に両者を分離しにくい

## 提案される構造

```
vrc-go/
├── api/
│   ├── client.go              (HTTPクライアント基盤)
│   ├── auth.go               (認証関連)
│   ├── users.go              (ユーザー操作)
│   ├── friends.go
│   ├── avatars.go
│   ├── worlds.go
│   ├── instances.go
│   ├── notifications.go
│   ├── favorites.go
│   ├── groups.go
│   ├── files.go
│   ├── player_moderation.go
│   └── system.go
├── websocket/
│   ├── client.go             (WebSocketClient型)
│   └── handlers.go           (イベントハンドラー定義)
├── shared/
│   ├── errors.go             (APIError型)
│   ├── types.go              (共有データモデル)
│   ├── cookie.go             (Cookie永続化)
│   ├── options.go            (NewClientオプション)
│   └── event_types.go        (WebSocketイベント型)
├── examples/                  (既存のまま)
└── CLAUDE.md, README.md etc
```

## 実装手順

### フェーズ 1: 共有パッケージ作成

1. `shared/` ディレクトリ作成
2. 既存の `vrchat/` から以下を移動：
   - `errors.go` → `shared/errors.go`
   - `cookie.go` → `shared/cookie.go`
   - `options.go` → `shared/options.go`
3. WebSocket イベント型を新規作成：
   - `shared/event_types.go` に以下を移動：
     - `Event`
     - `EventContent`
     - `NotificationEvent`
     - `FriendOnlineEvent`
     - `FriendOfflineEvent`
     - `FriendLocationEvent`
     - `UserUpdateEvent`
4. データモデル型を集約：
   - `shared/types.go` に共通のデータ構造を定義（User, Avatar, World等）
   - 既存ファイルから参照される型を整理

### フェーズ 2: API パッケージ構築

1. `api/` ディレクトリ作成
2. `vrchat/client.go` → `api/client.go` に移動
   - `Client` 型の定義は変わらない
   - `NewClient()` のシグネチャは変わらない
3. REST API 実装ファイルを移動：
   - `vrchat/*.go` → `api/` へ (websocket.go を除く)
4. インポートパスを更新：
   ```go
   // 例：auth.go内
   import "github.com/kqnade/vrc-go/shared"
   ```

### フェーズ 3: WebSocket パッケージ構築

1. `websocket/` ディレクトリ作成
2. `vrchat/websocket.go` をリファクタリング：
   - `Client` メソッドの削除（`ConnectWebSocket()` を `websocket` パッケージのファクトリ関数に）
   - `WebSocketClient` 型を移動
   - イベントハンドラーメソッドを `websocket/handlers.go` に分割（オプション）
3. クッキー取得ロジックを共有パッケージから利用

**修正内容：**
```go
// 現在：Client のメソッド
func (c *Client) ConnectWebSocket(ctx context.Context) (*WebSocketClient, error)

// 変更後：websocket パッケージのファクトリ関数
// websocket/client.go
func New(ctx context.Context, apiClient *api.Client) (*Client, error)
```

4. インポート：
   ```go
   import (
       "github.com/kqnade/vrc-go/api"
       "github.com/kqnade/vrc-go/shared"
   )
   ```

### フェーズ 4: Examples の更新

1. `examples/*/main.go` を新しいインポートパスに対応：
   ```go
   import (
       "github.com/kqnade/vrc-go/api"
       "github.com/kqnade/vrc-go/websocket"
   )

   func main() {
       apiClient, _ := api.NewClient()
       apiClient.Authenticate(ctx, username, password)

       wsClient, _ := websocket.New(ctx, apiClient)
       wsClient.OnFriendOnline(...)
   }
   ```

### フェーズ 5: go.mod の確認

- モジュールパスは `github.com/kqnade/vrc-go` のまま変わらない
- サブパッケージの相対インポートのみ更新

## 利点

- ✅ **明確な責務分離**：REST API とWebSocket が独立
- ✅ **インポートの柔軟性**：必要なパッケージだけインポート可能
  ```go
  // WebSocket のみ使う場合
  import "github.com/kqnade/vrc-go/websocket"

  // REST API のみ使う場合
  import "github.com/kqnade/vrc-go/api"
  ```
- ✅ **テスト性向上**：各パッケージを独立してテスト可能
- ✅ **保守性向上**：ファイル数が削減、関心ごとが明確
- ✅ **将来のスケーラビリティ**：REST API と WebSocket を独立したライブラリとして分割可能

## 破壊的変更

- **ユーザーのインポート文が変わる**
  ```go
  // 変更前
  import "github.com/kqnade/vrc-go/vrchat"
  c, _ := vrchat.NewClient()

  // 変更後
  import "github.com/kqnade/vrc-go/api"
  c, _ := api.NewClient()
  ```
- **バージョンバンプが必要**（v0.2.0 推奨）
- README.md とドキュメントの更新が必須

## 代替案

### 案 A: 別パッケージなし（現状維持）
- **利点**：破壊的変更がない
- **欠点**：関心事の混在が続く、スケーラビリティに問題

### 案 B: vrchat をサブパッケージ化
```go
// webhook を独立パッケージにするが、API操作は vrchat に残す
import (
    "github.com/kqnade/vrc-go/vrchat"
    "github.com/kqnade/vrc-go/vrchat/websocket"
)
```
- **利点**：部分的な改善、破壊的変更が少ない
- **欠点**：依然として関心事の混在がある

### 案 C: 提案内容（推奨）
```go
import (
    "github.com/kqnade/vrc-go/api"
    "github.com/kqnade/vrc-go/websocket"
)
```
- **利点**：完全な責務分離、明確な設計
- **欠点**：破壊的変更が大きい

## 推奨事項

**案 C（提案内容）を実装する** ことをお勧めします。理由：
1. VRChat API ライブラリはまだ v0.x で、メジャーバージョン未到達
2. 破壊的変更は今後が許容される段階
3. 長期的なメンテナンス性と拡張性を優先すべき
4. ユーザーベースがまだ小さい段階での変更がリスクが低い
