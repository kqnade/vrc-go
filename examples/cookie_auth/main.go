package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kqnade/vrc-go/vrchat"
)

func main() {
	// クライアント作成
	client, err := vrchat.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 保存済みCookieから認証
	err = client.LoadCookies("cookies.json")
	if err != nil {
		log.Fatalf("Failed to load cookies: %v\nPlease run basic_auth example first to save cookies.", err)
	}

	fmt.Println("✓ Cookies loaded")

	// ユーザー情報取得（Cookie認証の確認）
	user, err := client.GetCurrentUser(context.Background())
	if err != nil {
		log.Fatalf("Failed to get user info: %v\nCookies may have expired.", err)
	}

	fmt.Printf("Logged in as: %s (%s)\n", user.DisplayName, user.Username)
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Friends online: %d\n", len(user.OnlineFriends))
	fmt.Printf("Friends active: %d\n", len(user.ActiveFriends))
	fmt.Printf("Friends offline: %d\n", len(user.OfflineFriends))
}
