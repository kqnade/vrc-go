package vrchat

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// User はVRChatユーザーの情報です
type User struct {
	ID                     string   `json:"id"`
	DisplayName            string   `json:"displayName"`
	Username               string   `json:"username"`
	Bio                    string   `json:"bio"`
	Tags                   []string `json:"tags"`
	Status                 string   `json:"status"`
	StatusDescription      string   `json:"statusDescription"`
	CurrentAvatar          string   `json:"currentAvatar"`
	CurrentAvatarThumbnail string   `json:"currentAvatarImageUrl"`
	CurrentAvatarAssetURL  string   `json:"currentAvatarAssetUrl"`
	FallbackAvatar         string   `json:"fallbackAvatar"`
	ProfilePicOverride     string   `json:"profilePicOverride"`
	IsFriend               bool     `json:"isFriend"`
	FriendKey              string   `json:"friendKey"`
	LastLogin              string   `json:"last_login"`
	LastPlatform           string   `json:"last_platform"`
	AllowAvatarCopying     bool     `json:"allowAvatarCopying"`
	State                  string   `json:"state"`
	DateJoined             string   `json:"date_joined"`
	Location               string   `json:"location"`
	WorldID                string   `json:"worldId"`
	InstanceID             string   `json:"instanceId"`
	DeveloperType          string   `json:"developerType"`
	Note                   string   `json:"note"`
	PastDisplayNames       []struct {
		DisplayName string `json:"displayName"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"pastDisplayNames"`
}

// LimitedUser は制限されたユーザー情報です（検索結果など）
type LimitedUser struct {
	ID                     string   `json:"id"`
	DisplayName            string   `json:"displayName"`
	Username               string   `json:"username"`
	Bio                    string   `json:"bio"`
	Tags                   []string `json:"tags"`
	Status                 string   `json:"status"`
	StatusDescription      string   `json:"statusDescription"`
	CurrentAvatar          string   `json:"currentAvatar"`
	CurrentAvatarThumbnail string   `json:"currentAvatarImageUrl"`
	IsFriend               bool     `json:"isFriend"`
	FriendKey              string   `json:"friendKey"`
	LastLogin              string   `json:"last_login"`
	LastPlatform           string   `json:"last_platform"`
	Location               string   `json:"location"`
	DeveloperType          string   `json:"developerType"`
}

// UpdateUserRequest はユーザー情報更新リクエストです
type UpdateUserRequest struct {
	Email             *string  `json:"email,omitempty"`
	Birthday          *string  `json:"birthday,omitempty"`
	AcceptedTOSVersion *int    `json:"acceptedTOSVersion,omitempty"`
	Tags              []string `json:"tags,omitempty"`
	Status            *string  `json:"status,omitempty"`
	StatusDescription *string  `json:"statusDescription,omitempty"`
	Bio               *string  `json:"bio,omitempty"`
	BioLinks          []string `json:"bioLinks,omitempty"`
	UserIcon          *string  `json:"userIcon,omitempty"`
}

// GetUser は指定されたユーザーIDのユーザー情報を取得します
func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	var user User
	err := c.doRequest(ctx, "GET", "/users/"+userID, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// GetUserByName は指定されたユーザー名のユーザー情報を取得します
func (c *Client) GetUserByName(ctx context.Context, username string) (*User, error) {
	var user User
	err := c.doRequest(ctx, "GET", "/users/"+username+"/name", nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by name: %w", err)
	}
	return &user, nil
}

// SearchUsersOptions はユーザー検索のオプションです
type SearchUsersOptions struct {
	Search         string // 検索クエリ
	DeveloperType  string // フィルター: "none", "trusted", "internal", "moderator"
	N              int    // 取得件数（デフォルト: 60）
	Offset         int    // オフセット
	Fuzzy          bool   // あいまい検索
}

// SearchUsers はユーザーを検索します
func (c *Client) SearchUsers(ctx context.Context, opts SearchUsersOptions) ([]LimitedUser, error) {
	params := url.Values{}
	if opts.Search != "" {
		params.Set("search", opts.Search)
	}
	if opts.DeveloperType != "" {
		params.Set("developerType", opts.DeveloperType)
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Fuzzy {
		params.Set("fuzzy", "true")
	}

	var users []LimitedUser
	path := "/users?" + params.Encode()
	err := c.doRequest(ctx, "GET", path, nil, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to search users: %w", err)
	}
	return users, nil
}

// UpdateUser は現在のユーザー情報を更新します
func (c *Client) UpdateUser(ctx context.Context, userID string, req UpdateUserRequest) (*CurrentUser, error) {
	var user CurrentUser
	err := c.doRequest(ctx, "PUT", "/users/"+userID, req, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &user, nil
}

// UserGroup はユーザーのグループ情報です
type UserGroup struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	ShortCode               string   `json:"shortCode"`
	DiscriminatorString     string   `json:"discriminator"`
	Description             string   `json:"description"`
	IconURL                 string   `json:"iconUrl"`
	BannerURL               string   `json:"bannerUrl"`
	Privacy                 string   `json:"privacy"`
	OwnerId                 string   `json:"ownerId"`
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
	Galleries               []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		MembersOnly bool   `json:"membersOnly"`
		RoleIDsToView []string `json:"roleIdsToView"`
		RoleIDsToSubmit []string `json:"roleIdsToSubmit"`
		RoleIDsToAutoApprove []string `json:"roleIdsToAutoApprove"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
	} `json:"galleries"`
	CreatedAt string `json:"createdAt"`
	OnlineMemberCount int `json:"onlineMemberCount"`
	MembershipStatus string `json:"membershipStatus"`
	MyMember struct {
		ID                 string   `json:"id"`
		GroupID            string   `json:"groupId"`
		UserID             string   `json:"userId"`
		RoleIDs            []string `json:"roleIds"`
		JoinedAt           string   `json:"joinedAt"`
		MembershipStatus   string   `json:"membershipStatus"`
		Visibility         string   `json:"visibility"`
		IsSubscribedToAnnouncements bool `json:"isSubscribedToAnnouncements"`
	} `json:"myMember"`
}

// GetUserGroups は指定されたユーザーが所属するグループのリストを取得します
func (c *Client) GetUserGroups(ctx context.Context, userID string) ([]UserGroup, error) {
	var groups []UserGroup
	err := c.doRequest(ctx, "GET", "/users/"+userID+"/groups", nil, &groups)
	if err != nil {
		return nil, fmt.Errorf("failed to get user groups: %w", err)
	}
	return groups, nil
}

// GetUserRepresentedGroup は指定されたユーザーが代表として表示しているグループを取得します
func (c *Client) GetUserRepresentedGroup(ctx context.Context, userID string) (*UserGroup, error) {
	var group UserGroup
	err := c.doRequest(ctx, "GET", "/users/"+userID+"/groups/represented", nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to get user represented group: %w", err)
	}
	return &group, nil
}
