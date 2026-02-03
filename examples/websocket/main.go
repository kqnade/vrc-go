package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kqnade/vrcgo/shared"
	"github.com/kqnade/vrcgo/vrcapi"
	"github.com/kqnade/vrcgo/vrcws"
)

func main() {
	// 環境変数から認証情報を取得
	username := os.Getenv("VRCHAT_USERNAME")
	password := os.Getenv("VRCHAT_PASSWORD")
	totpCode := os.Getenv("VRCHAT_TOTP_CODE")

	if username == "" || password == "" {
		log.Fatal("VRCHAT_USERNAME and VRCHAT_PASSWORD environment variables are required")
	}

	// クライアント作成
	client, err := vrcapi.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 認証
	fmt.Println("🔐 Authenticating...")
	err = client.Authenticate(context.Background(), shared.AuthConfig{
		Username: username,
		Password: password,
		TOTPCode: totpCode,
	})
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}
	fmt.Println("✓ Authentication successful!")

	// ユーザー情報取得
	user, err := client.GetCurrentUser(context.Background())
	if err != nil {
		log.Fatalf("Failed to get user info: %v", err)
	}
	fmt.Printf("✓ Logged in as: %s (%s)\n\n", user.DisplayName, user.Username)

	// WebSocket接続
	fmt.Println("🔌 Connecting to WebSocket...")
	ws, err := vrcws.New(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to connect websocket: %v", err)
	}
	defer ws.Close()
	fmt.Println("✓ WebSocket connected!")
	fmt.Println("\n📡 Listening for events... (Press Ctrl+C to exit)")

	// すべてのイベントをログ
	ws.On("*", func(event shared.Event) {
		fmt.Printf("📨 Event [%s]: %s\n", event.Type, string(event.Content))
	})

	// 通知イベント
	ws.OnNotification(func(notification shared.NotificationEvent) {
		fmt.Printf("🔔 Notification: %s from %s\n", notification.Type, notification.SenderUsername)
		data, _ := json.MarshalIndent(notification, "  ", "  ")
		fmt.Printf("  %s\n\n", data)
	})

	// フレンドオンライン
	ws.OnFriendOnline(func(friend shared.FriendOnlineEvent) {
		userName := friend.UserID
		if friend.User != nil {
			userName = friend.User.DisplayName
		}
		fmt.Printf("✅ Friend Online: %s @ %s\n\n", userName, friend.Location)
	})

	// フレンドオフライン
	ws.OnFriendOffline(func(friend shared.FriendOfflineEvent) {
		fmt.Printf("❌ Friend Offline: %s\n\n", friend.UserID)
	})

	// フレンドロケーション変更
	ws.OnFriendLocation(func(friend shared.FriendLocationEvent) {
		userName := friend.UserID
		if friend.User != nil {
			userName = friend.User.DisplayName
		}
		fmt.Printf("📍 Friend Location: %s moved to %s\n\n", userName, friend.Location)
	})

	// フレンド追加
	ws.OnFriendAdd(func(event shared.FriendAddEvent) {
		userName := event.UserID
		if event.User != nil {
			userName = event.User.DisplayName
		}
		fmt.Printf("➕ Friend Added: %s (%s)\n\n", userName, event.UserID)
	})

	// フレンド削除
	ws.OnFriendDelete(func(event shared.FriendDeleteEvent) {
		fmt.Printf("➖ Friend Deleted: %s\n\n", event.UserID)
	})

	// ユーザー更新
	ws.OnUserUpdate(func(user shared.UserUpdateEvent) {
		fmt.Printf("👤 User Update: %s\n\n", user.UserID)
	})

	// シグナル待機
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\n\n👋 Disconnecting...")
}
