package vrchat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// WebSocketURL はVRChat WebSocketのエンドポイントです
	WebSocketURL = "wss://pipeline.vrchat.cloud/"

	// 再接続の設定
	reconnectDelay     = 5 * time.Second
	maxReconnectDelay  = 60 * time.Second
	reconnectDelayMult = 2
)

// WebSocketClient はWebSocket接続を管理します
type WebSocketClient struct {
	client       *Client
	conn         *websocket.Conn
	handlers     map[string][]EventHandler
	handlersMux  sync.RWMutex
	done         chan struct{}
	reconnect    bool
	authToken    string
	ctx          context.Context
	cancel       context.CancelFunc
}

// EventHandler はイベントハンドラー関数です
type EventHandler func(event Event)

// Event はWebSocketイベントです
type Event struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

// EventContent はイベントコンテンツの共通フィールドです
type EventContent struct {
	UserID      string `json:"userId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	World       string `json:"world,omitempty"`
	Location    string `json:"location,omitempty"`
	Instance    string `json:"instance,omitempty"`
	GroupID     string `json:"groupId,omitempty"`
}

// NotificationEvent は通知イベントです
type NotificationEvent struct {
	ID           string                 `json:"id"`
	Type         string                 `json:"type"`
	SenderUserID string                 `json:"senderUserId"`
	SenderUsername string               `json:"senderUsername"`
	ReceiverUserID string               `json:"receiverUserId"`
	Message      string                 `json:"message"`
	Details      map[string]interface{} `json:"details"`
	Seen         bool                   `json:"seen"`
	CreatedAt    string                 `json:"created_at"`
}

// FriendOnlineEvent はフレンドオンラインイベントです
type FriendOnlineEvent struct {
	UserID      string `json:"userId"`
	User        *LimitedUser `json:"user,omitempty"`
	Location    string `json:"location"`
	Instance    string `json:"instance"`
	World       string `json:"world"`
	CanRequestInvite bool `json:"canRequestInvite"`
	Platform    string `json:"platform"`
}

// FriendOfflineEvent はフレンドオフラインイベントです
type FriendOfflineEvent struct {
	UserID string `json:"userId"`
}

// FriendLocationEvent はフレンドロケーション変更イベントです
type FriendLocationEvent struct {
	UserID      string `json:"userId"`
	User        *LimitedUser `json:"user,omitempty"`
	Location    string `json:"location"`
	Instance    string `json:"instance"`
	World       string `json:"world"`
	CanRequestInvite bool `json:"canRequestInvite"`
}

// UserUpdateEvent はユーザー更新イベントです
type UserUpdateEvent struct {
	UserID      string `json:"userId"`
	User        *CurrentUser `json:"user,omitempty"`
}

// ConnectWebSocket はWebSocket接続を確立します
func (c *Client) ConnectWebSocket(ctx context.Context) (*WebSocketClient, error) {
	// authcookieを取得
	authToken, err := c.getAuthCookie()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth cookie: %w", err)
	}

	wsClient := &WebSocketClient{
		client:    c,
		handlers:  make(map[string][]EventHandler),
		done:      make(chan struct{}),
		reconnect: true,
		authToken: authToken,
	}
	wsClient.ctx, wsClient.cancel = context.WithCancel(ctx)

	if err := wsClient.connect(); err != nil {
		return nil, err
	}

	// イベントループを開始
	go wsClient.readLoop()

	return wsClient, nil
}

// getAuthCookie はCookieJarから認証クッキーを取得します
func (c *Client) getAuthCookie() (string, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	cookies := c.httpClient.Jar.Cookies(u)
	for _, cookie := range cookies {
		if cookie.Name == "auth" || cookie.Name == "authcookie" {
			return cookie.Value, nil
		}
	}

	return "", fmt.Errorf("auth cookie not found")
}

// connect はWebSocket接続を確立します
func (ws *WebSocketClient) connect() error {
	u, err := url.Parse(WebSocketURL)
	if err != nil {
		return fmt.Errorf("failed to parse websocket URL: %w", err)
	}

	q := u.Query()
	q.Set("authToken", ws.authToken)
	u.RawQuery = q.Encode()

	header := http.Header{}
	header.Set("User-Agent", ws.client.userAgent)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return fmt.Errorf("failed to connect websocket: %w", err)
	}

	ws.conn = conn
	return nil
}

// readLoop はメッセージ受信ループです
func (ws *WebSocketClient) readLoop() {
	defer close(ws.done)

	currentDelay := reconnectDelay

	for {
		select {
		case <-ws.ctx.Done():
			return
		default:
		}

		if ws.conn == nil {
			if !ws.reconnect {
				return
			}

			// 再接続を試みる
			time.Sleep(currentDelay)
			if err := ws.connect(); err != nil {
				// 指数バックオフ
				currentDelay *= reconnectDelayMult
				if currentDelay > maxReconnectDelay {
					currentDelay = maxReconnectDelay
				}
				continue
			}
			currentDelay = reconnectDelay // 接続成功時にリセット
		}

		var event Event
		err := ws.conn.ReadJSON(&event)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// 予期しない切断
			}
			ws.conn.Close()
			ws.conn = nil
			continue
		}

		// イベントハンドラーを実行
		ws.handleEvent(event)
	}
}

// handleEvent はイベントを処理します
func (ws *WebSocketClient) handleEvent(event Event) {
	ws.handlersMux.RLock()
	handlers := ws.handlers[event.Type]
	ws.handlersMux.RUnlock()

	for _, handler := range handlers {
		go handler(event)
	}

	// 全イベントハンドラー
	ws.handlersMux.RLock()
	allHandlers := ws.handlers["*"]
	ws.handlersMux.RUnlock()

	for _, handler := range allHandlers {
		go handler(event)
	}
}

// On はイベントハンドラーを登録します
func (ws *WebSocketClient) On(eventType string, handler EventHandler) {
	ws.handlersMux.Lock()
	defer ws.handlersMux.Unlock()

	ws.handlers[eventType] = append(ws.handlers[eventType], handler)
}

// OnNotification は通知イベントハンドラーを登録します
func (ws *WebSocketClient) OnNotification(handler func(notification NotificationEvent)) {
	ws.On("notification", func(event Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var notification NotificationEvent
		if err := json.Unmarshal([]byte(content), &notification); err != nil {
			return
		}

		handler(notification)
	})
}

// OnFriendOnline はフレンドオンラインイベントハンドラーを登録します
func (ws *WebSocketClient) OnFriendOnline(handler func(friend FriendOnlineEvent)) {
	ws.On("friend-online", func(event Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friend FriendOnlineEvent
		if err := json.Unmarshal([]byte(content), &friend); err != nil {
			return
		}

		handler(friend)
	})
}

// OnFriendOffline はフレンドオフラインイベントハンドラーを登録します
func (ws *WebSocketClient) OnFriendOffline(handler func(friend FriendOfflineEvent)) {
	ws.On("friend-offline", func(event Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friend FriendOfflineEvent
		if err := json.Unmarshal([]byte(content), &friend); err != nil {
			return
		}

		handler(friend)
	})
}

// OnFriendLocation はフレンドロケーション変更イベントハンドラーを登録します
func (ws *WebSocketClient) OnFriendLocation(handler func(friend FriendLocationEvent)) {
	ws.On("friend-location", func(event Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friend FriendLocationEvent
		if err := json.Unmarshal([]byte(content), &friend); err != nil {
			return
		}

		handler(friend)
	})
}

// OnFriendActive はフレンドアクティブイベントハンドラーを登録します
func (ws *WebSocketClient) OnFriendActive(handler func(content EventContent)) {
	ws.On("friend-active", func(event Event) {
		var contentStr string
		if err := json.Unmarshal(event.Content, &contentStr); err != nil {
			return
		}

		var content EventContent
		if err := json.Unmarshal([]byte(contentStr), &content); err != nil {
			return
		}

		handler(content)
	})
}

// OnFriendAdd はフレンド追加イベントハンドラーを登録します
func (ws *WebSocketClient) OnFriendAdd(handler func(content EventContent)) {
	ws.On("friend-add", func(event Event) {
		var contentStr string
		if err := json.Unmarshal(event.Content, &contentStr); err != nil {
			return
		}

		var content EventContent
		if err := json.Unmarshal([]byte(contentStr), &content); err != nil {
			return
		}

		handler(content)
	})
}

// OnFriendDelete はフレンド削除イベントハンドラーを登録します
func (ws *WebSocketClient) OnFriendDelete(handler func(content EventContent)) {
	ws.On("friend-delete", func(event Event) {
		var contentStr string
		if err := json.Unmarshal(event.Content, &contentStr); err != nil {
			return
		}

		var content EventContent
		if err := json.Unmarshal([]byte(contentStr), &content); err != nil {
			return
		}

		handler(content)
	})
}

// OnUserUpdate はユーザー更新イベントハンドラーを登録します
func (ws *WebSocketClient) OnUserUpdate(handler func(user UserUpdateEvent)) {
	ws.On("user-update", func(event Event) {
		var contentStr string
		if err := json.Unmarshal(event.Content, &contentStr); err != nil {
			return
		}

		var user UserUpdateEvent
		if err := json.Unmarshal([]byte(contentStr), &user); err != nil {
			return
		}

		handler(user)
	})
}

// Close はWebSocket接続を閉じます
func (ws *WebSocketClient) Close() error {
	ws.reconnect = false
	ws.cancel()

	if ws.conn != nil {
		err := ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			return err
		}

		<-ws.done
		return ws.conn.Close()
	}

	<-ws.done
	return nil
}

// Wait はWebSocket接続が終了するまで待機します
func (ws *WebSocketClient) Wait() {
	<-ws.done
}
