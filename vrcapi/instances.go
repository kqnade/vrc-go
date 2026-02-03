package vrcapi

import (
	"context"
	"fmt"

	"github.com/kqnade/vrcgo/shared"
)

// GetInstance は指定されたワールドIDとインスタンスIDのインスタンス情報を取得します
func (c *Client) GetInstance(ctx context.Context, worldID, instanceID string) (*shared.Instance, error) {
	var instance shared.Instance
	err := c.doRequest(ctx, "GET", "/instances/"+worldID+":"+instanceID, nil, &instance)
	if err != nil {
		return nil, fmt.Errorf("failed to get instance: %w", err)
	}
	return &instance, nil
}

// GetInstanceByShortName は短縮名でインスタンス情報を取得します
func (c *Client) GetInstanceByShortName(ctx context.Context, shortName string) (*shared.Instance, error) {
	var instance shared.Instance
	err := c.doRequest(ctx, "GET", "/instances/s/"+shortName, nil, &instance)
	if err != nil {
		return nil, fmt.Errorf("failed to get instance by short name: %w", err)
	}
	return &instance, nil
}

// SendSelfInvite は自分自身にインスタンスへの招待を送信します
func (c *Client) SendSelfInvite(ctx context.Context, worldID, instanceID string) (*shared.Notification, error) {
	var notification shared.Notification
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

// CreateInstance は新しいインスタンスを作成します
func (c *Client) CreateInstance(ctx context.Context, req shared.CreateInstanceRequest) (*shared.Instance, error) {
	var instance shared.Instance
	err := c.doRequest(ctx, "POST", "/instances", req, &instance)
	if err != nil {
		return nil, fmt.Errorf("failed to create instance: %w", err)
	}
	return &instance, nil
}

// UpdateInstance はインスタンスを更新します
func (c *Client) UpdateInstance(ctx context.Context, worldID, instanceID string, req shared.UpdateInstanceRequest) (*shared.Instance, error) {
	var instance shared.Instance
	err := c.doRequest(ctx, "PUT", "/instances/"+worldID+":"+instanceID, req, &instance)
	if err != nil {
		return nil, fmt.Errorf("failed to update instance: %w", err)
	}
	return &instance, nil
}

// CloseInstance はインスタンスをクローズします
func (c *Client) CloseInstance(ctx context.Context, worldID, instanceID string) error {
	err := c.doRequest(ctx, "DELETE", "/instances/"+worldID+":"+instanceID, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to close instance: %w", err)
	}
	return nil
}
