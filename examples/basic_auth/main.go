package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kqnade/vrc-go/vrchat"
)

func main() {
	// 環境変数から認証情報を取得
	username := os.Getenv("VRCHAT_USERNAME")
	password := os.Getenv("VRCHAT_PASSWORD")
	totpCode := os.Getenv("VRCHAT_TOTP_CODE") // オプション

	if username == "" || password == "" {
		log.Fatal("VRCHAT_USERNAME and VRCHAT_PASSWORD environment variables are required")
	}

	// クライアント作成
	client, err := vrchat.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// 認証
	err = client.Authenticate(context.Background(), vrchat.AuthConfig{
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

	fmt.Printf("Logged in as: %s (%s)\n", user.DisplayName, user.Username)
	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Status: %s\n", user.Status)
	if user.StatusDescription != "" {
		fmt.Printf("Status Description: %s\n", user.StatusDescription)
	}

	// Cookieを保存
	if err := client.SaveCookies("cookies.json"); err != nil {
		log.Printf("Warning: Failed to save cookies: %v", err)
	} else {
		fmt.Println("✓ Cookies saved to cookies.json")
	}
}
