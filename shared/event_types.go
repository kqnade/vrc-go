package shared

import "encoding/json"

// Event はWebSocketイベントです
type Event struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

// EventHandler はイベントハンドラー関数です
type EventHandler func(event Event)

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
	ID             string                 `json:"id"`
	Type           string                 `json:"type"`
	SenderUserID   string                 `json:"senderUserId"`
	SenderUsername string                 `json:"senderUsername"`
	ReceiverUserID string                 `json:"receiverUserId"`
	Message        string                 `json:"message"`
	Details        map[string]interface{} `json:"details"`
	Seen           bool                   `json:"seen"`
	CreatedAt      string                 `json:"created_at"`
}

// FriendOnlineEvent はフレンドオンラインイベントです
type FriendOnlineEvent struct {
	UserID           string      `json:"userId"`
	User             *LimitedUser `json:"user,omitempty"`
	Location         string      `json:"location"`
	Instance         string      `json:"instance"`
	World            string      `json:"world"`
	CanRequestInvite bool        `json:"canRequestInvite"`
	Platform         string      `json:"platform"`
}

// FriendOfflineEvent はフレンドオフラインイベントです
type FriendOfflineEvent struct {
	UserID string `json:"userId"`
}

// FriendLocationEvent はフレンドロケーション変更イベントです
type FriendLocationEvent struct {
	UserID           string      `json:"userId"`
	User             *LimitedUser `json:"user,omitempty"`
	Location         string      `json:"location"`
	Instance         string      `json:"instance"`
	World            string      `json:"world"`
	CanRequestInvite bool        `json:"canRequestInvite"`
}

// FriendActiveEvent はフレンドアクティブイベントです
type FriendActiveEvent struct {
	UserID string `json:"userId"`
	User   *LimitedUser `json:"user,omitempty"`
}

// FriendAddEvent はフレンド追加イベントです
type FriendAddEvent struct {
	UserID string `json:"userId"`
	User   *LimitedUser `json:"user,omitempty"`
}

// FriendDeleteEvent はフレンド削除イベントです
type FriendDeleteEvent struct {
	UserID string `json:"userId"`
}

// UserUpdateEvent はユーザー更新イベントです
type UserUpdateEvent struct {
	UserID string       `json:"userId"`
	User   *CurrentUser `json:"user,omitempty"`
}

// GroupJoinedEvent はグループ参加イベントです
type GroupJoinedEvent struct {
	GroupID string `json:"groupId"`
	Group   *Group `json:"group,omitempty"`
}

// GroupLeftEvent はグループ脱退イベントです
type GroupLeftEvent struct {
	GroupID string `json:"groupId"`
}

// GroupAnnoucementEvent はグループお知らせイベントです
type GroupAnnoucementEvent struct {
	GroupID      string              `json:"groupId"`
	Announcement *GroupAnnouncement  `json:"announcement,omitempty"`
}

// NotificationV2Event は通知v2イベントです
type NotificationV2Event struct {
	ID              string                 `json:"id"`
	NotificationType string                `json:"notificationType"`
	SenderUserID    string                 `json:"senderUserId"`
	ReceiverUserID  string                 `json:"receiverUserId"`
	Message         string                 `json:"message"`
	Details         map[string]interface{} `json:"details"`
	Read            bool                   `json:"read"`
	CreatedAt       string                 `json:"createdAt"`
}
