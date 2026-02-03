# パッケージ構成とモジュール名のリファクタリング計画（v0.3.0）

## 概要

現状の `vrchat` パッケージは REST API と WebSocket 機能が混在しており、かつモジュール名も `vrc-go` で一貫性が弱い。

これを v0.3.0 で以下のように整理する：

- モジュール名を **`github.com/kqnade/vrcgo`** に変更
- REST API 用パッケージ **`vrcapi`**
- WebSocket 用パッケージ **`vrcws`**
- 共有型/ユーティリティ用パッケージ **`shared`**
- 旧 `vrchat` パッケージは廃止（必要なら薄いラッパとして一時的に残すが、基本は削除）

## 現在の問題点

1. **関心事の混在**
   - `vrchat.Client` 型に以下が同居している：
     - REST API メソッド（`Authenticate()`, `GetUser()`, `SearchWorlds()` など多数）
     - `ConnectWebSocket()` など WebSocket 関連メソッド
   - 全く異なる責務が同じ型／同じパッケージに混在している。

2. **WebSocket クライアントの依存関係が重い**
   - `WebSocketClient` が `vrchat.Client` に依存しているが、
   - 実際に必要なのは「認証クッキーや設定情報」程度であり、依存が太い。

3. **モジュール／パッケージ名の一貫性不足**
   - モジュール名が `github.com/kqnade/vrc-go`
   - パッケージ候補が `api` / `websocket` など汎用名
   - `vrc`（VRChat）であることがモジュール／パッケージレベルで統一されていない。

4. **スケーラビリティとテスト性**
   - REST API が増えると `Client` が肥大化
   - WebSocket イベントやハンドラが増えても同じパッケージ内で追う必要がある
   - テスト時に REST / WebSocket を明確に分離しにくい

## 目標構成

```text
github.com/kqnade/vrcgo/
├── vrcapi/
│   ├── client.go              (HTTPクライアント基盤 / REST API Client 型)
│   ├── auth.go                (認証関連)
│   ├── users.go               (ユーザー操作)
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
├── vrcws/
│   ├── client.go              (WebSocket クライアント)
│   └── handlers.go            (イベントハンドラー定義)
├── shared/
│   ├── errors.go              (APIError 型など共通エラー)
│   ├── types.go               (User, Avatar, World など共有データモデル)
│   ├── cookie.go              (Cookie 永続化)
│   ├── options.go             (NewClient オプション)
│   └── event_types.go         (WebSocket イベント型)
├── examples/                  (サンプルコード)
└── README.md, CLAUDE.md など
```

- **モジュール名**：`github.com/kqnade/vrcgo`
- **REST API**：`github.com/kqnade/vrcgo/vrcapi`
- **WebSocket**：`github.com/kqnade/vrcgo/vrcws`

## 実装手順

### フェーズ 0: モジュール名とリポジトリ名の変更

1. `go.mod` の module 行を変更：

   ```mod
   module github.com/kqnade/vrcgo
   ```

2. GitHub のリポジトリ名を `vrcgo` にリネームする。
3. README 等に記載されている旧パス
   - `github.com/kqnade/vrc-go`
   をすべて
   - `github.com/kqnade/vrcgo`
   に更新する。

> 備考: 現時点（v0.2.0）では利用者がほぼいない前提のため、このモジュール名変更による破壊的変更は許容する。

---

### フェーズ 1: shared パッケージ作成

1. `shared/` ディレクトリを作成。

2. 既存 `vrchat/` から以下を移動 or 抜き出す：

   - `errors.go` → `shared/errors.go`
   - `cookie.go` → `shared/cookie.go`
   - `options.go` → `shared/options.go`

3. WebSocket イベント型を `shared/event_types.go` に集約：

   例：

   - `Event`
   - `EventContent`
   - `NotificationEvent`
   - `FriendOnlineEvent`
   - `FriendOfflineEvent`
   - `FriendLocationEvent`
   - `FriendActiveEvent`
   - `FriendAddEvent`
   - `FriendDeleteEvent`
   - `UserUpdateEvent`
   - など WebSocket イベントに対応する型

4. 共通データモデル（REST / WebSocket 両方から参照されるもの）を `shared/types.go` に集約：

   - `User`
   - `Avatar`
   - `World`
   - `Instance`
   - `Notification`
   - その他共通で使う構造体

5. REST / WebSocket 双方から `shared` をインポートする形にする：

   ```go
   import "github.com/kqnade/vrcgo/shared"
   ```

---

### フェーズ 2: vrcapi パッケージ構築（REST API）

1. `vrcapi/` ディレクトリを作成。

2. 既存の `vrchat/client.go` 相当を `vrcapi/client.go` に移植：

   - `type Client struct { ... }`
   - `func NewClient(opts ...Option) (*Client, error)` など
   - HTTP クライアントの初期化・設定をここに集約。

3. REST API 実装ファイルを `vrcapi/` に移動：

   - `auth.go`
   - `users.go`
   - `friends.go`
   - `avatars.go`
   - `worlds.go`
   - `instances.go`
   - `notifications.go`
   - `favorites.go`
   - `groups.go`
   - `files.go`
   - `player_moderation.go`
   - `system.go`

4. `vrcapi` 内のファイルから `shared` を参照するよう import を更新：

   ```go
   import "github.com/kqnade/vrcgo/shared"
   ```

5. 旧 `vrchat` パッケージ内の REST 関連コードは削除する（互換性レイヤーは原則用意しない）。

---

### フェーズ 3: vrcws パッケージ構築（WebSocket）

1. `vrcws/` ディレクトリを作成。

2. 既存の WebSocket 関連コード（`vrchat/websocket.go` 相当）を `vrcws/client.go` / `vrcws/handlers.go` に分割：

   - `type Client struct { ... }`（WebSocket 専用クライアント）
   - `OnNotification`, `OnFriendOnline` などのイベントハンドラ登録用メソッドを `handlers.go` に整理してもよい。

3. `vrcapi.Client` との依存関係をシンプルにする：

   ```go
   // vrcws/client.go
   import (
       "context"

       "github.com/kqnade/vrcgo/vrcapi"
       "github.com/kqnade/vrcgo/shared"
   )

   func New(ctx context.Context, apiClient *vrcapi.Client) (*Client, error) {
       // apiClient から必要な Cookie / 設定を取得して WebSocket 接続
   }
   ```

4. WebSocket イベント型は `shared` パッケージの定義を使う：

   ```go
   func (c *Client) OnFriendOnline(handler func(shared.FriendOnlineEvent)) {
       // ...
   }
   ```

5. 旧 `vrchat.Client` のメソッド `ConnectWebSocket()` は削除し、`vrcws.New()` に一本化する。

---

### フェーズ 4: Examples の更新

1. `examples/*/main.go` を新しいインポートパスに変更：

   ```go
   package main

   import (
       "context"
       "log"
       "os"

       "github.com/kqnade/vrcgo/vrcapi"
       "github.com/kqnade/vrcgo/vrcws"
       "github.com/kqnade/vrcgo/shared"
   )

   func main() {
       ctx := context.Background()

       apiClient, err := vrcapi.NewClient()
       if err != nil {
           log.Fatal(err)
       }

       if err := apiClient.Authenticate(ctx, vrcapi.AuthConfig{
           Username: os.Getenv("VRCHAT_USERNAME"),
           Password: os.Getenv("VRCHAT_PASSWORD"),
       }); err != nil {
           log.Fatal(err)
       }

       wsClient, err := vrcws.New(ctx, apiClient)
       if err != nil {
           log.Fatal(err)
       }

       wsClient.OnFriendOnline(func(ev shared.FriendOnlineEvent) {
           log.Printf("%s is now online at %s", ev.UserID, ev.Location)
       })

       wsClient.Wait()
   }
   ```

2. Makefile のターゲット（`run-basic-auth`, `run-cookie-auth`, `run-websocket` など）が、`examples/` 内の新しいコードを正しく実行するように更新する。

---

### フェーズ 5: README / ドキュメント更新

1. インストール方法：

   ```bash
   go get github.com/kqnade/vrcgo
   ```

2. 基本認証のサンプル：

   ```go
   import (
       "context"
       "log"

       "github.com/kqnade/vrcgo/vrcapi"
   )

   func main() {
       ctx := context.Background()

       client, err := vrcapi.NewClient(
           vrcapi.WithUserAgent("my-app/1.0"),
       )
       if err != nil {
           log.Fatal(err)
       }

       if err := client.Authenticate(ctx, vrcapi.AuthConfig{
           Username: "your-username",
           Password: "your-password",
       }); err != nil {
           log.Fatal(err)
       }
   }
   ```

3. WebSocket のサンプル：

   ```go
   import (
       "context"
       "log"

       "github.com/kqnade/vrcgo/vrcapi"
       "github.com/kqnade/vrcgo/vrcws"
       "github.com/kqnade/vrcgo/shared"
   )

   func main() {
       ctx := context.Background()

       apiClient, err := vrcapi.NewClient(
           vrcapi.WithUserAgent("my-app/1.0"),
       )
       if err != nil {
           log.Fatal(err)
       }

       if err := apiClient.Authenticate(ctx, vrcapi.AuthConfig{...}); err != nil {
           log.Fatal(err)
       }

       wsClient, err := vrcws.New(ctx, apiClient)
       if err != nil {
           log.Fatal(err)
       }

       wsClient.OnFriendOnline(func(ev shared.FriendOnlineEvent) {
           log.Printf("%s is now online at %s", ev.UserID, ev.Location)
       })

       wsClient.Wait()
   }
   ```

4. 旧 `vrchat` や `github.com/kqnade/vrc-go` に言及している箇所は全て更新する。

---

## 利点

- ✅ **明確な責務分離**
  - REST API → `vrcapi`
  - WebSocket → `vrcws`
  - 共通ロジック → `shared`

- ✅ **名前とモジュールの一貫性**
  - モジュール名とパッケージ名が `vrc` 系で揃う（`vrcgo` / `vrcapi` / `vrcws`）

- ✅ **テストしやすさ**
  - REST / WebSocket をそれぞれ独立にテスト可能

- ✅ **今のタイミングだからこそできる大整理**
  - v0.2.0 時点で利用者がほぼいないため、大きな破壊的変更を安全に実施できる

## 破壊的変更

- モジュールパスの変更：

  ```go
  // 変更前
  import "github.com/kqnade/vrc-go/vrchat"

  // 変更後
  import "github.com/kqnade/vrcgo/vrcapi"
  ```

- パッケージ構成の変更：

  ```go
  // 変更前
  c, _ := vrchat.NewClient()
  ws, _ := c.ConnectWebSocket(ctx)

  // 変更後
  apiClient, _ := vrcapi.NewClient()
  wsClient, _ := vrcws.New(ctx, apiClient)
  ```

- v0.3.0 でこの変更を行う。
