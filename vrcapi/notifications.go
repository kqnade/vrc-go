package vrcapi

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kqnade/vrcgo/shared"
)

// GetNotifications は通知のリストを取得します
func (c *Client) GetNotifications(ctx context.Context, opts shared.GetNotificationsOptions) ([]shared.Notification, error) {
	params := url.Values{}
	if opts.Type != "" {
		params.Set("type", opts.Type)
	}
	if opts.Sent {
		params.Set("sent", "true")
	}
	if opts.Hidden {
		params.Set("hidden", "true")
	}
	if opts.After != "" {
		params.Set("after", opts.After)
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}

	var notifications []shared.Notification
	path := "/auth/user/notifications"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &notifications)
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications: %w", err)
	}
	return notifications, nil
}

// MarkNotificationAsRead は通知を既読にマークします
func (c *Client) MarkNotificationAsRead(ctx context.Context, notificationID string) (*shared.Notification, error) {
	var notification shared.Notification
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/"+notificationID+"/see", nil, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to mark notification as read: %w", err)
	}
	return &notification, nil
}

// DeleteNotification は通知を削除します
func (c *Client) DeleteNotification(ctx context.Context, notificationID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/"+notificationID+"/hide", nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}
	return nil
}

// ClearAllNotifications はすべての通知をクリアします
func (c *Client) ClearAllNotifications(ctx context.Context) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/clear", nil, &response)
	if err != nil {
		return fmt.Errorf("failed to clear all notifications: %w", err)
	}
	return nil
}

// SendNotification は通知を送信します
func (c *Client) SendNotification(ctx context.Context, req shared.SendNotificationRequest) (*shared.Notification, error) {
	var notification shared.Notification
	err := c.doRequest(ctx, "POST", "/auth/user/notifications", req, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %w", err)
	}
	return &notification, nil
}

// RespondToNotification は通知に応答します
func (c *Client) RespondToNotification(ctx context.Context, notificationID string, response string) (*shared.Notification, error) {
	var notification shared.Notification
	req := struct {
		Response string `json:"response"`
	}{Response: response}
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/"+notificationID+"/respond", req, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to respond to notification: %w", err)
	}
	return &notification, nil
}
