package vrcapi

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kqnade/vrcgo/shared"
)

// GetGroup は指定されたグループIDのグループ情報を取得します
func (c *Client) GetGroup(ctx context.Context, groupID string) (*shared.Group, error) {
	var group shared.Group
	err := c.doRequest(ctx, "GET", "/groups/"+groupID, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to get group: %w", err)
	}
	return &group, nil
}

// SearchGroups はグループを検索します
func (c *Client) SearchGroups(ctx context.Context, query string, n, offset int) ([]shared.Group, error) {
	var groups []shared.Group
	params := url.Values{}
	params.Set("query", query)
	params.Set("n", strconv.Itoa(n))
	params.Set("offset", strconv.Itoa(offset))
	path := "/groups?" + params.Encode()
	err := c.doRequest(ctx, "GET", path, nil, &groups)
	if err != nil {
		return nil, fmt.Errorf("failed to search groups: %w", err)
	}
	return groups, nil
}

// JoinGroup はグループに参加します
func (c *Client) JoinGroup(ctx context.Context, groupID string) error {
	err := c.doRequest(ctx, "POST", "/groups/"+groupID+"/join", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to join group: %w", err)
	}
	return nil
}

// LeaveGroup はグループから脱退します
func (c *Client) LeaveGroup(ctx context.Context, groupID string) error {
	err := c.doRequest(ctx, "DELETE", "/groups/"+groupID+"/leave", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to leave group: %w", err)
	}
	return nil
}

// GetGroupMembers はグループのメンバーリストを取得します
func (c *Client) GetGroupMembers(ctx context.Context, groupID string, n, offset int) ([]shared.GroupMember, error) {
	var members []shared.GroupMember
	path := fmt.Sprintf("/groups/%s/members?n=%d&offset=%d", groupID, n, offset)
	err := c.doRequest(ctx, "GET", path, nil, &members)
	if err != nil {
		return nil, fmt.Errorf("failed to get group members: %w", err)
	}
	return members, nil
}

// BanGroupMember はグループメンバーをBANします
func (c *Client) BanGroupMember(ctx context.Context, groupID, userID string) error {
	err := c.doRequest(ctx, "POST", "/groups/"+groupID+"/bans", map[string]string{"userId": userID}, nil)
	if err != nil {
		return fmt.Errorf("failed to ban group member: %w", err)
	}
	return nil
}

// UnbanGroupMember はグループメンバーのBANを解除します
func (c *Client) UnbanGroupMember(ctx context.Context, groupID, userID string) error {
	err := c.doRequest(ctx, "DELETE", "/groups/"+groupID+"/bans/"+userID, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to unban group member: %w", err)
	}
	return nil
}

// GetGroupAnnouncements はグループのお知らせリストを取得します
func (c *Client) GetGroupAnnouncements(ctx context.Context, groupID string) ([]shared.GroupAnnouncement, error) {
	var announcements []shared.GroupAnnouncement
	err := c.doRequest(ctx, "GET", "/groups/"+groupID+"/announcements", nil, &announcements)
	if err != nil {
		return nil, fmt.Errorf("failed to get group announcements: %w", err)
	}
	return announcements, nil
}
