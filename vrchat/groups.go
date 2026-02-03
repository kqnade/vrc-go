package vrchat

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// Group はグループ情報です
type Group struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	ShortCode               string   `json:"shortCode"`
	DiscriminatorString     string   `json:"discriminator"`
	Description             string   `json:"description"`
	IconURL                 string   `json:"iconUrl"`
	BannerURL               string   `json:"bannerUrl"`
	Privacy                 string   `json:"privacy"`
	OwnerID                 string   `json:"ownerId"`
	Rules                   string   `json:"rules"`
	Links                   []string `json:"links"`
	Languages               []string `json:"languages"`
	IconID                  string   `json:"iconId"`
	BannerID                string   `json:"bannerId"`
	MemberCount             int      `json:"memberCount"`
	MemberCountSyncedAt     string   `json:"memberCountSyncedAt"`
	IsVerified              bool     `json:"isVerified"`
	JoinState               string   `json:"joinState"`
	Tags                    []string `json:"tags"`
	CreatedAt               string   `json:"createdAt"`
	OnlineMemberCount       int      `json:"onlineMemberCount"`
	MembershipStatus        string   `json:"membershipStatus"`
}

// GetGroup は指定されたグループIDのグループ情報を取得します
func (c *Client) GetGroup(ctx context.Context, groupID string) (*Group, error) {
	var group Group
	err := c.doRequest(ctx, "GET", "/groups/"+groupID, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to get group: %w", err)
	}
	return &group, nil
}

// SearchGroupsOptions はグループ検索のオプションです
type SearchGroupsOptions struct {
	Query  string // 検索クエリ
	Offset int    // オフセット
	N      int    // 取得件数（デフォルト: 60）
}

// SearchGroups はグループを検索します
func (c *Client) SearchGroups(ctx context.Context, opts SearchGroupsOptions) ([]Group, error) {
	params := url.Values{}
	if opts.Query != "" {
		params.Set("query", opts.Query)
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}

	var groups []Group
	path := "/groups"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &groups)
	if err != nil {
		return nil, fmt.Errorf("failed to search groups: %w", err)
	}
	return groups, nil
}

// CreateGroupRequest はグループ作成リクエストです
type CreateGroupRequest struct {
	Name        string   `json:"name"`
	ShortCode   string   `json:"shortCode"`
	Description string   `json:"description"`
	JoinState   string   `json:"joinState"` // "open", "request", "invite"
	IconID      string   `json:"iconId,omitempty"`
	BannerID    string   `json:"bannerId,omitempty"`
	Privacy     string   `json:"privacy,omitempty"`
	RoleTemplate string  `json:"roleTemplate,omitempty"`
}

// CreateGroup は新しいグループを作成します
func (c *Client) CreateGroup(ctx context.Context, req CreateGroupRequest) (*Group, error) {
	var group Group
	err := c.doRequest(ctx, "POST", "/groups", req, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}
	return &group, nil
}

// UpdateGroupRequest はグループ更新リクエストです
type UpdateGroupRequest struct {
	Name          *string  `json:"name,omitempty"`
	ShortCode     *string  `json:"shortCode,omitempty"`
	Description   *string  `json:"description,omitempty"`
	JoinState     *string  `json:"joinState,omitempty"`
	IconID        *string  `json:"iconId,omitempty"`
	BannerID      *string  `json:"bannerId,omitempty"`
	Languages     []string `json:"languages,omitempty"`
	Links         []string `json:"links,omitempty"`
	Rules         *string  `json:"rules,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

// UpdateGroup はグループ情報を更新します
func (c *Client) UpdateGroup(ctx context.Context, groupID string, req UpdateGroupRequest) (*Group, error) {
	var group Group
	err := c.doRequest(ctx, "PUT", "/groups/"+groupID, req, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to update group: %w", err)
	}
	return &group, nil
}

// DeleteGroup はグループを削除します
func (c *Client) DeleteGroup(ctx context.Context, groupID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/groups/"+groupID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}
	return nil
}

// JoinGroup はグループに参加します
func (c *Client) JoinGroup(ctx context.Context, groupID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "POST", "/groups/"+groupID+"/join", nil, &response)
	if err != nil {
		return fmt.Errorf("failed to join group: %w", err)
	}
	return nil
}

// LeaveGroup はグループから脱退します
func (c *Client) LeaveGroup(ctx context.Context, groupID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "POST", "/groups/"+groupID+"/leave", nil, &response)
	if err != nil {
		return fmt.Errorf("failed to leave group: %w", err)
	}
	return nil
}

// GroupMember はグループメンバー情報です
type GroupMember struct {
	ID                          string   `json:"id"`
	GroupID                     string   `json:"groupId"`
	UserID                      string   `json:"userId"`
	IsRepresenting              bool     `json:"isRepresenting"`
	User                        *LimitedUser `json:"user,omitempty"`
	RoleIDs                     []string `json:"roleIds"`
	MRoleIDs                    []string `json:"mRoleIds"`
	JoinedAt                    string   `json:"joinedAt"`
	MembershipStatus            string   `json:"membershipStatus"`
	Visibility                  string   `json:"visibility"`
	IsSubscribedToAnnouncements bool     `json:"isSubscribedToAnnouncements"`
	CreatedAt                   *string  `json:"createdAt,omitempty"`
	BannedAt                    *string  `json:"bannedAt,omitempty"`
	ManagerNotes                *string  `json:"managerNotes,omitempty"`
}

// GetGroupMembers はグループメンバーのリストを取得します
func (c *Client) GetGroupMembers(ctx context.Context, groupID string, n, offset int) ([]GroupMember, error) {
	params := url.Values{}
	if n > 0 {
		params.Set("n", strconv.Itoa(n))
	} else {
		params.Set("n", "60")
	}
	if offset > 0 {
		params.Set("offset", strconv.Itoa(offset))
	}

	var members []GroupMember
	path := "/groups/" + groupID + "/members"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &members)
	if err != nil {
		return nil, fmt.Errorf("failed to get group members: %w", err)
	}
	return members, nil
}

// BanGroupMemberRequest はグループメンバーBanリクエストです
type BanGroupMemberRequest struct {
	UserID string `json:"userId"`
}

// BanGroupMember はグループメンバーをBanします
func (c *Client) BanGroupMember(ctx context.Context, groupID, userID string) error {
	req := BanGroupMemberRequest{UserID: userID}
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "POST", "/groups/"+groupID+"/bans", req, &response)
	if err != nil {
		return fmt.Errorf("failed to ban group member: %w", err)
	}
	return nil
}

// UnbanGroupMember はグループメンバーのBanを解除します
func (c *Client) UnbanGroupMember(ctx context.Context, groupID, userID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/groups/"+groupID+"/bans/"+userID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to unban group member: %w", err)
	}
	return nil
}

// GroupAnnouncement はグループアナウンス情報です
type GroupAnnouncement struct {
	ID        string `json:"id"`
	GroupID   string `json:"groupId"`
	AuthorID  string `json:"authorId"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	ImageID   string `json:"imageId"`
	ImageURL  string `json:"imageUrl"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// CreateGroupAnnouncementRequest はグループアナウンス作成リクエストです
type CreateGroupAnnouncementRequest struct {
	Title   string `json:"title"`
	Text    string `json:"text"`
	ImageID string `json:"imageId,omitempty"`
	SendNotification bool `json:"sendNotification,omitempty"`
}

// CreateGroupAnnouncement はグループアナウンスを作成します
func (c *Client) CreateGroupAnnouncement(ctx context.Context, groupID string, req CreateGroupAnnouncementRequest) (*GroupAnnouncement, error) {
	var announcement GroupAnnouncement
	err := c.doRequest(ctx, "POST", "/groups/"+groupID+"/announcements", req, &announcement)
	if err != nil {
		return nil, fmt.Errorf("failed to create group announcement: %w", err)
	}
	return &announcement, nil
}

// DeleteGroupAnnouncement はグループアナウンスを削除します
func (c *Client) DeleteGroupAnnouncement(ctx context.Context, groupID, announcementID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/groups/"+groupID+"/announcements/"+announcementID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete group announcement: %w", err)
	}
	return nil
}
