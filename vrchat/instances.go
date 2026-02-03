package vrchat

import (
	"context"
	"fmt"
	"net/url"
)

// Instance はインスタンス情報です
type Instance struct {
	ID               string   `json:"id"`
	Private          *string  `json:"private,omitempty"`
	Friends          *string  `json:"friends,omitempty"`
	Users            []string `json:"users"`
	Name             string   `json:"name"`
	WorldID          string   `json:"worldId"`
	OwnerID          string   `json:"ownerId"`
	Tags             []string `json:"tags"`
	Active           bool     `json:"active"`
	Full             bool     `json:"full"`
	N_Users          int      `json:"n_users"`
	Capacity         int      `json:"capacity"`
	RecommendedCapacity int   `json:"recommendedCapacity"`
	InstanceID       string   `json:"instanceId"`
	Location         string   `json:"location"`
	ShortName        string   `json:"shortName"`
	SecureName       string   `json:"secureName"`
	World            *World   `json:"world,omitempty"`
	Type             string   `json:"type"`
	Region           string   `json:"region"`
	CanRequestInvite bool     `json:"canRequestInvite"`
	Permanent        bool     `json:"permanent"`
	Platforms        struct {
		Android int `json:"android"`
		StandaloneWindows int `json:"standalonewindows"`
	} `json:"platforms"`
}

// GetInstance は指定されたワールドIDとインスタンスIDのインスタンス情報を取得します
func (c *Client) GetInstance(ctx context.Context, worldID, instanceID string) (*Instance, error) {
	var instance Instance
	err := c.doRequest(ctx, "GET", "/instances/"+worldID+":"+instanceID, nil, &instance)
	if err != nil {
		return nil, fmt.Errorf("failed to get instance: %w", err)
	}
	return &instance, nil
}

// GetInstanceByShortName は短縮名でインスタンス情報を取得します
func (c *Client) GetInstanceByShortName(ctx context.Context, shortName string) (*Instance, error) {
	var instance Instance
	params := url.Values{}
	params.Set("shortName", shortName)
	err := c.doRequest(ctx, "GET", "/instances/s/"+shortName, nil, &instance)
	if err != nil {
		return nil, fmt.Errorf("failed to get instance by short name: %w", err)
	}
	return &instance, nil
}

// SendSelfInvite は自分自身にインスタンスへの招待を送信します
func (c *Client) SendSelfInvite(ctx context.Context, worldID, instanceID string) (*Notification, error) {
	var notification Notification
	err := c.doRequest(ctx, "POST", "/instances/"+worldID+":"+instanceID+"/invite", nil, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to send self invite: %w", err)
	}
	return &notification, nil
}

// GetInstanceShortName はインスタンスの短縮名を取得します
func (c *Client) GetInstanceShortName(ctx context.Context, worldID, instanceID string) (string, error) {
	var response struct {
		ShortName  string `json:"shortName"`
		SecureName string `json:"secureName"`
	}
	err := c.doRequest(ctx, "GET", "/instances/"+worldID+":"+instanceID+"/shortName", nil, &response)
	if err != nil {
		return "", fmt.Errorf("failed to get instance short name: %w", err)
	}
	return response.ShortName, nil
}

// CreateInstanceRequest はインスタンス作成リクエストです
type CreateInstanceRequest struct {
	WorldID      string `json:"worldId"`
	Type         string `json:"type"` // "public", "friends+", "friends", "invite+", "invite"
	Region       string `json:"region"`
	InstanceID   string `json:"instanceId,omitempty"`
	OwnerID      string `json:"ownerId,omitempty"`
	GroupAccessType string `json:"groupAccessType,omitempty"`
	QueueEnabled bool   `json:"queueEnabled,omitempty"`
	CanRequestInvite bool `json:"canRequestInvite,omitempty"`
}

// CreateInstance は新しいインスタンスを作成します
func (c *Client) CreateInstance(ctx context.Context, req CreateInstanceRequest) (*Instance, error) {
	var instance Instance
	err := c.doRequest(ctx, "POST", "/instances", req, &instance)
	if err != nil {
		return nil, fmt.Errorf("failed to create instance: %w", err)
	}
	return &instance, nil
}

// CloseInstance はインスタンスを閉じます（インスタンスのオーナーのみ）
func (c *Client) CloseInstance(ctx context.Context, worldID, instanceID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/instances/"+worldID+":"+instanceID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to close instance: %w", err)
	}
	return nil
}
