package vrcws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/kqnade/vrcgo/shared"
	"github.com/kqnade/vrcgo/vrcapi"
)

const (
	// WebSocketURL はVRChat WebSocketのエンドポイントです
	WebSocketURL = "wss://pipeline.vrchat.cloud/"

	// 再接続の設定
	reconnectDelay     = 5 * time.Second
	maxReconnectDelay  = 60 * time.Second
	reconnectDelayMult = 2
)

// Client はWebSocket接続を管理します
type Client struct {
	conn         *websocket.Conn
	handlers     map[string][]shared.EventHandler
	handlersMux  sync.RWMutex
	done         chan struct{}
	reconnect    bool
	authToken    string
	userAgent    string
	ctx          context.Context
	cancel       context.CancelFunc
}

// New は新しいWebSocketクライアントを作成します
func New(ctx context.Context, apiClient *vrcapi.Client) (*Client, error) {
	// authcookieを取得
	authToken, err := apiClient.GetAuthCookie()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth cookie: %w", err)
	}

	wsClient := &Client{
		handlers:  make(map[string][]shared.EventHandler),
		done:      make(chan struct{}),
		reconnect: true,
		authToken: authToken,
		userAgent: "vrc-go/0.1.0", // デフォルト
	}
	wsClient.ctx, wsClient.cancel = context.WithCancel(ctx)

	if err := wsClient.connect(); err != nil {
		return nil, err
	}

	// イベントループを開始
	go wsClient.readLoop()

	return wsClient, nil
}

// connect はWebSocket接続を確立します
func (ws *Client) connect() error {
	u, err := url.Parse(WebSocketURL)
	if err != nil {
		return fmt.Errorf("failed to parse websocket URL: %w", err)
	}

	q := u.Query()
	q.Set("authToken", ws.authToken)
	u.RawQuery = q.Encode()

	header := http.Header{}
	header.Set("User-Agent", ws.userAgent)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return fmt.Errorf("failed to connect websocket: %w", err)
	}

	ws.conn = conn
	return nil
}

// readLoop はメッセージ受信ループです
func (ws *Client) readLoop() {
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

		var event shared.Event
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
func (ws *Client) handleEvent(event shared.Event) {
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
func (ws *Client) On(eventType string, handler shared.EventHandler) {
	ws.handlersMux.Lock()
	defer ws.handlersMux.Unlock()

	ws.handlers[eventType] = append(ws.handlers[eventType], handler)
}

// OnNotification は通知イベントハンドラーを登録します
func (ws *Client) OnNotification(handler func(notification shared.NotificationEvent)) {
	ws.On("notification", func(event shared.Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var notification shared.NotificationEvent
		if err := json.Unmarshal([]byte(content), &notification); err != nil {
			return
		}

		handler(notification)
	})
}

// OnFriendOnline はフレンドオンラインイベントハンドラーを登録します
func (ws *Client) OnFriendOnline(handler func(friend shared.FriendOnlineEvent)) {
	ws.On("friend-online", func(event shared.Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friend shared.FriendOnlineEvent
		if err := json.Unmarshal([]byte(content), &friend); err != nil {
			return
		}

		handler(friend)
	})
}

// OnFriendOffline はフレンドオフラインイベントハンドラーを登録します
func (ws *Client) OnFriendOffline(handler func(friend shared.FriendOfflineEvent)) {
	ws.On("friend-offline", func(event shared.Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friend shared.FriendOfflineEvent
		if err := json.Unmarshal([]byte(content), &friend); err != nil {
			return
		}

		handler(friend)
	})
}

// OnFriendLocation はフレンドロケーション変更イベントハンドラーを登録します
func (ws *Client) OnFriendLocation(handler func(friend shared.FriendLocationEvent)) {
	ws.On("friend-location", func(event shared.Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friend shared.FriendLocationEvent
		if err := json.Unmarshal([]byte(content), &friend); err != nil {
			return
		}

		handler(friend)
	})
}

// OnFriendActive はフレンドアクティブイベントハンドラーを登録します
func (ws *Client) OnFriendActive(handler func(event shared.FriendActiveEvent)) {
	ws.On("friend-active", func(event shared.Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friendEvent shared.FriendActiveEvent
		if err := json.Unmarshal([]byte(content), &friendEvent); err != nil {
			return
		}

		handler(friendEvent)
	})
}

// OnFriendAdd はフレンド追加イベントハンドラーを登録します
func (ws *Client) OnFriendAdd(handler func(event shared.FriendAddEvent)) {
	ws.On("friend-add", func(event shared.Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friendEvent shared.FriendAddEvent
		if err := json.Unmarshal([]byte(content), &friendEvent); err != nil {
			return
		}

		handler(friendEvent)
	})
}

// OnFriendDelete はフレンド削除イベントハンドラーを登録します
func (ws *Client) OnFriendDelete(handler func(event shared.FriendDeleteEvent)) {
	ws.On("friend-delete", func(event shared.Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var friendEvent shared.FriendDeleteEvent
		if err := json.Unmarshal([]byte(content), &friendEvent); err != nil {
			return
		}

		handler(friendEvent)
	})
}

// OnUserUpdate はユーザー更新イベントハンドラーを登録します
func (ws *Client) OnUserUpdate(handler func(user shared.UserUpdateEvent)) {
	ws.On("user-update", func(event shared.Event) {
		var content string
		if err := json.Unmarshal(event.Content, &content); err != nil {
			return
		}

		var user shared.UserUpdateEvent
		if err := json.Unmarshal([]byte(content), &user); err != nil {
			return
		}

		handler(user)
	})
}

// Close はWebSocket接続を閉じます
func (ws *Client) Close() error {
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
func (ws *Client) Wait() {
	<-ws.done
}
