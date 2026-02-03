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

### 認証 (Authentication)

- `Authenticate(ctx, config)` - ユーザー名/パスワードでログイン（2FA対応）
- `GetCurrentUser(ctx)` - 現在のユーザー情報を取得
- `Logout(ctx)` - ログアウト

### ユーザー (Users)

- `GetUser(ctx, userId)` - ユーザー情報を取得
- `GetUserByName(ctx, username)` - ユーザー名からユーザー情報を取得
- `SearchUsers(ctx, opts)` - ユーザーを検索
- `UpdateUser(ctx, userId, req)` - ユーザー情報を更新
- `GetUserGroups(ctx, userId)` - ユーザーのグループリストを取得
- `GetUserRepresentedGroup(ctx, userId)` - 代表グループを取得

### フレンド (Friends)

- `GetFriends(ctx, opts)` - フレンドリストを取得
- `GetFriendStatus(ctx, userId)` - フレンドステータスを取得
- `SendFriendRequest(ctx, userId)` - フレンドリクエストを送信
- `DeleteFriend(ctx, userId)` - フレンドを削除
- `AcceptFriendRequest(ctx, notificationId)` - フレンドリクエストを承認
- `DeclineFriendRequest(ctx, notificationId)` - フレンドリクエストを拒否
- `GetOnlineFriends(ctx)` - オンラインフレンドを取得
- `GetOfflineFriends(ctx)` - オフラインフレンドを取得

### アバター (Avatars)

- `GetAvatar(ctx, avatarId)` - アバター情報を取得
- `SearchAvatars(ctx, opts)` - アバターを検索
- `SelectAvatar(ctx, avatarId)` - アバターを装着
- `CreateAvatar(ctx, req)` - アバターを作成
- `UpdateAvatar(ctx, avatarId, req)` - アバター情報を更新
- `DeleteAvatar(ctx, avatarId)` - アバターを削除
- `GetFavoriteAvatars(ctx)` - お気に入りアバターを取得

### ワールド (Worlds)

- `GetWorld(ctx, worldId)` - ワールド情報を取得
- `SearchWorlds(ctx, opts)` - ワールドを検索
- `GetActiveWorlds(ctx)` - アクティブなワールドを取得
- `GetRecentWorlds(ctx)` - 最近訪問したワールドを取得
- `GetFavoriteWorlds(ctx)` - お気に入りワールドを取得
- `CreateWorld(ctx, req)` - ワールドを作成
- `UpdateWorld(ctx, worldId, req)` - ワールド情報を更新
- `DeleteWorld(ctx, worldId)` - ワールドを削除
- `GetWorldMetadata(ctx, worldId)` - ワールドメタデータを取得
- `PublishWorld(ctx, worldId)` - ワールドを公開
- `UnpublishWorld(ctx, worldId)` - ワールドを非公開化

### インスタンス (Instances)

- `GetInstance(ctx, worldId, instanceId)` - インスタンス情報を取得
- `GetInstanceByShortName(ctx, shortName)` - 短縮名でインスタンスを取得
- `SendSelfInvite(ctx, worldId, instanceId)` - 自分自身に招待を送信
- `CreateInstance(ctx, req)` - インスタンスを作成
- `CloseInstance(ctx, worldId, instanceId)` - インスタンスを閉じる

### 通知 (Notifications)

- `GetNotifications(ctx, opts)` - 通知リストを取得
- `MarkNotificationAsRead(ctx, notificationId)` - 通知を既読にする
- `DeleteNotification(ctx, notificationId)` - 通知を削除
- `ClearAllNotifications(ctx)` - すべての通知をクリア
- `SendNotification(ctx, req)` - 通知を送信
- `SendInvite(ctx, userId, req)` - インスタンス招待を送信
- `RespondToInvite(ctx, notificationId, req)` - 招待に応答
- `RequestInvite(ctx, userId, instanceLocation)` - 招待をリクエスト

### お気に入り (Favorites)

- `AddFavorite(ctx, req)` - お気に入りを追加
- `RemoveFavorite(ctx, favoriteId)` - お気に入りを削除
- `GetFavorites(ctx, opts)` - お気に入りリストを取得
- `GetFavoriteGroups(ctx, type)` - お気に入りグループを取得
- `UpdateFavoriteGroup(ctx, type, name, userId, req)` - お気に入りグループを更新
- `ClearFavoriteGroup(ctx, type, name, userId)` - お気に入りグループをクリア

### グループ (Groups)

- `GetGroup(ctx, groupId)` - グループ情報を取得
- `SearchGroups(ctx, opts)` - グループを検索
- `CreateGroup(ctx, req)` - グループを作成
- `UpdateGroup(ctx, groupId, req)` - グループ情報を更新
- `DeleteGroup(ctx, groupId)` - グループを削除
- `JoinGroup(ctx, groupId)` - グループに参加
- `LeaveGroup(ctx, groupId)` - グループから脱退
- `GetGroupMembers(ctx, groupId, n, offset)` - グループメンバーを取得
- `BanGroupMember(ctx, groupId, userId)` - メンバーをBAN
- `UnbanGroupMember(ctx, groupId, userId)` - メンバーのBANを解除
- `CreateGroupAnnouncement(ctx, groupId, req)` - グループアナウンスを作成
- `DeleteGroupAnnouncement(ctx, groupId, announcementId)` - アナウンスを削除

### ファイル (Files)

- `GetFile(ctx, fileId)` - ファイル情報を取得
- `CreateFile(ctx, req)` - ファイルを作成
- `DeleteFile(ctx, fileId)` - ファイルを削除
- `DownloadFile(ctx, fileId, version)` - ファイルダウンロードURLを取得
- `CreateFileVersion(ctx, fileId, req)` - ファイルバージョンを作成
- `DeleteFileVersion(ctx, fileId, version)` - ファイルバージョンを削除

### プレイヤーモデレーション (Player Moderation)

- `ModerateUser(ctx, req)` - ユーザーをモデレート
- `GetPlayerModerations(ctx)` - モデレーションリストを取得
- `MuteUser(ctx, userId)` - ユーザーをミュート
- `UnmuteUser(ctx, userId)` - ミュートを解除
- `BlockUser(ctx, userId)` - ユーザーをブロック
- `UnblockUser(ctx, userId)` - ブロックを解除
- `HideUserAvatar(ctx, userId)` - アバターを非表示
- `ShowUserAvatar(ctx, userId)` - アバターを表示

### システム (System)

- `GetConfig(ctx)` - システム設定を取得
- `GetSystemTime(ctx)` - サーバー時刻を取得
- `GetInfoPushes(ctx)` - 情報プッシュを取得
- `GetCurrentOnlineUsers(ctx)` - オンラインユーザー数を取得
- `GetHealth(ctx)` - APIヘルスチェック

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
    } else if vrchat.IsNotFoundError(err) {
        log.Println("Resource not found")
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
