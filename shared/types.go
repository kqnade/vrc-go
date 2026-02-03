package shared

import "encoding/json"

// AuthConfig は認証設定です
type AuthConfig struct {
	Username string
	Password string
	TOTPCode string // 2要素認証コード（オプション）
}

// CurrentUser は現在のユーザー情報です
type CurrentUser struct {
	ID                      string   `json:"id"`
	DisplayName             string   `json:"displayName"`
	Username                string   `json:"username"`
	Bio                     string   `json:"bio"`
	Tags                    []string `json:"tags"`
	Status                  string   `json:"status"`
	StatusDescription       string   `json:"statusDescription"`
	CurrentAvatar           string   `json:"currentAvatar"`
	CurrentAvatarThumbnail  string   `json:"currentAvatarImageUrl"`
	RequiresTwoFactorAuth   []string `json:"requiresTwoFactorAuth,omitempty"`
	EmailVerified           bool     `json:"emailVerified"`
	HasBirthday             bool     `json:"hasBirthday"`
	HasEmail                bool     `json:"hasEmail"`
	HasPendingEmail         bool     `json:"hasPendingEmail"`
	ObfuscatedEmail         string   `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail  string   `json:"obfuscatedPendingEmail"`
	SteamID                 string   `json:"steamId"`
	OculusID                string   `json:"oculusId"`
	AccountDeletionDate     *string  `json:"accountDeletionDate,omitempty"`
	AccountDeletionLog      *string  `json:"accountDeletionLog,omitempty"`
	AcceptedTOSVersion      int      `json:"acceptedTOSVersion"`
	AcceptedPrivacyVersion  int      `json:"acceptedPrivacyVersion"`
	SteamDetails            struct{} `json:"steamDetails"`
	OculusDetails           struct{} `json:"oculusDetails"`
	HasLoggedInFromClient   bool     `json:"hasLoggedInFromClient"`
	FriendKey               string   `json:"friendKey"`
	OnlineFriends           []string `json:"onlineFriends"`
	ActiveFriends           []string `json:"activeFriends"`
	OfflineFriends          []string `json:"offlineFriends"`
	FriendGroupNames        []string `json:"friendGroupNames"`
	CurrentAvatarAssetURL   string   `json:"currentAvatarAssetUrl"`
	FallbackAvatar          string   `json:"fallbackAvatar"`
	IsFriend                bool     `json:"isFriend"`
	LastLogin               string   `json:"last_login"`
	LastPlatform            string   `json:"last_platform"`
	AllowAvatarCopying      bool     `json:"allowAvatarCopying"`
	State                   string   `json:"state"`
	DateJoined              string   `json:"date_joined"`
	PastDisplayNames        []struct {
		DisplayName string `json:"displayName"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"pastDisplayNames"`
	TwoFactorAuthEnabled      bool    `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate  *string `json:"twoFactorAuthEnabledDate,omitempty"`
}

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
	Email              *string  `json:"email,omitempty"`
	Birthday           *string  `json:"birthday,omitempty"`
	AcceptedTOSVersion *int     `json:"acceptedTOSVersion,omitempty"`
	Tags               []string `json:"tags,omitempty"`
	Status             *string  `json:"status,omitempty"`
	StatusDescription  *string  `json:"statusDescription,omitempty"`
	Bio                *string  `json:"bio,omitempty"`
	BioLinks           []string `json:"bioLinks,omitempty"`
	UserIcon           *string  `json:"userIcon,omitempty"`
}

// UserGroup はユーザーのグループ情報です
type UserGroup struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	ShortCode           string   `json:"shortCode"`
	DiscriminatorString string   `json:"discriminator"`
	Description         string   `json:"description"`
	IconURL             string   `json:"iconUrl"`
	BannerURL           string   `json:"bannerUrl"`
	Privacy             string   `json:"privacy"`
	OwnerId             string   `json:"ownerId"`
	Rules               string   `json:"rules"`
	Links               []string `json:"links"`
	Languages           []string `json:"languages"`
	IconID              string   `json:"iconId"`
	BannerID            string   `json:"bannerId"`
	MemberCount         int      `json:"memberCount"`
	MemberCountSyncedAt string   `json:"memberCountSyncedAt"`
	IsVerified          bool     `json:"isVerified"`
	JoinState           string   `json:"joinState"`
	Tags                []string `json:"tags"`
	Galleries           []struct {
		ID                string   `json:"id"`
		Name              string   `json:"name"`
		Description       string   `json:"description"`
		MembersOnly       bool     `json:"membersOnly"`
		RoleIDsToView     []string `json:"roleIdsToView"`
		RoleIDsToSubmit   []string `json:"roleIdsToSubmit"`
		RoleIDsToAutoApprove []string `json:"roleIdsToAutoApprove"`
		CreatedAt         string   `json:"createdAt"`
		UpdatedAt         string   `json:"updatedAt"`
	} `json:"galleries"`
	CreatedAt           string `json:"createdAt"`
	OnlineMemberCount   int    `json:"onlineMemberCount"`
	MembershipStatus    string `json:"membershipStatus"`
	MyMember            struct {
		ID                              string   `json:"id"`
		GroupID                         string   `json:"groupId"`
		UserID                          string   `json:"userId"`
		RoleIDs                         []string `json:"roleIds"`
		JoinedAt                        string   `json:"joinedAt"`
		MembershipStatus                string   `json:"membershipStatus"`
		Visibility                      string   `json:"visibility"`
		IsSubscribedToAnnouncements     bool     `json:"isSubscribedToAnnouncements"`
	} `json:"myMember"`
}

// Avatar はアバター情報です
type Avatar struct {
	ID                    string         `json:"id"`
	Name                  string         `json:"name"`
	Description           string         `json:"description"`
	AuthorID              string         `json:"authorId"`
	AuthorName            string         `json:"authorName"`
	Tags                  []string       `json:"tags"`
	AssetURL              string         `json:"assetUrl"`
	AssetURLObject        string         `json:"assetUrlObject"`
	ImageURL              string         `json:"imageUrl"`
	ThumbnailImageURL     string         `json:"thumbnailImageUrl"`
	ReleaseStatus         string         `json:"releaseStatus"`
	Version               int            `json:"version"`
	Featured              bool           `json:"featured"`
	UnityPackages         []UnityPackage `json:"unityPackages"`
	UnityPackageURL       string         `json:"unityPackageUrl"`
	UnityPackageURLObject string         `json:"unityPackageUrlObject"`
	CreatedAt             string         `json:"created_at"`
	UpdatedAt             string         `json:"updated_at"`
}

// UnityPackage はUnityパッケージ情報です
type UnityPackage struct {
	ID               string `json:"id"`
	AssetURL         string `json:"assetUrl"`
	AssetURLObject   string `json:"assetUrlObject"`
	AssetVersion     int    `json:"assetVersion"`
	CreatedAt        string `json:"created_at"`
	Platform         string `json:"platform"`
	PluginURL        string `json:"pluginUrl"`
	PluginURLObject  string `json:"pluginUrlObject"`
	UnitySortNumber  int64  `json:"unitySortNumber"`
	UnityVersion     string `json:"unityVersion"`
}

// World はワールド情報です
type World struct {
	ID                   string         `json:"id"`
	Name                 string         `json:"name"`
	Description          string         `json:"description"`
	AuthorID             string         `json:"authorId"`
	AuthorName           string         `json:"authorName"`
	TotalLikes           int            `json:"likes"`
	TotalVisits          int            `json:"visits"`
	Capacity             int            `json:"capacity"`
	RecommendedCapacity  int            `json:"recommendedCapacity"`
	ImageURL             string         `json:"imageUrl"`
	ThumbnailImageURL    string         `json:"thumbnailImageUrl"`
	ReleaseStatus        string         `json:"releaseStatus"`
	Organization         string         `json:"organization"`
	Tags                 []string       `json:"tags"`
	Favorites            int            `json:"favorites"`
	CreatedAt            string         `json:"created_at"`
	UpdatedAt            string         `json:"updated_at"`
	PublicationDate      string         `json:"publicationDate"`
	LabsPublicationDate  string         `json:"labsPublicationDate"`
	Instances            [][]interface{} `json:"instances"`
	PublicOccupants      int            `json:"publicOccupants"`
	PrivateOccupants     int            `json:"privateOccupants"`
	Occupants            int            `json:"occupants"`
	UnityPackages        []UnityPackage `json:"unityPackages"`
	Namespace            string         `json:"namespace"`
	Version              int            `json:"version"`
	PreviewYoutubeID     *string        `json:"previewYoutubeId"`
	UdonProducts         []string       `json:"udonProducts"`
	Heat                 int            `json:"heat"`
}

// LimitedWorld は制限されたワールド情報です
type LimitedWorld struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	AuthorID            string         `json:"authorId"`
	AuthorName          string         `json:"authorName"`
	TotalLikes          int            `json:"likes"`
	Capacity            int            `json:"capacity"`
	RecommendedCapacity int            `json:"recommendedCapacity"`
	ImageURL            string         `json:"imageUrl"`
	ThumbnailImageURL   string         `json:"thumbnailImageUrl"`
	ReleaseStatus       string         `json:"releaseStatus"`
	Organization        string         `json:"organization"`
	Tags                []string       `json:"tags"`
	Favorites           int            `json:"favorites"`
	CreatedAt           string         `json:"created_at"`
	UpdatedAt           string         `json:"updated_at"`
	PublicationDate     string         `json:"publicationDate"`
	LabsPublicationDate string         `json:"labsPublicationDate"`
	Heat                int            `json:"heat"`
	PublicOccupants     int            `json:"publicOccupants"`
	PrivateOccupants    int            `json:"privateOccupants"`
	Occupants           int            `json:"occupants"`
	UnityPackages       []UnityPackage `json:"unityPackages"`
}

// Instance はインスタンス情報です
type Instance struct {
	ID                  string   `json:"id"`
	Private             *string  `json:"private,omitempty"`
	Friends             *string  `json:"friends,omitempty"`
	Users               []string `json:"users"`
	Name                string   `json:"name"`
	WorldID             string   `json:"worldId"`
	OwnerID             string   `json:"ownerId"`
	Tags                []string `json:"tags"`
	Active              bool     `json:"active"`
	Full                bool     `json:"full"`
	N_Users             int      `json:"n_users"`
	Capacity            int      `json:"capacity"`
	RecommendedCapacity int      `json:"recommendedCapacity"`
	InstanceID          string   `json:"instanceId"`
	Location            string   `json:"location"`
	ShortName           string   `json:"shortName"`
	SecureName          string   `json:"secureName"`
	World               *World   `json:"world,omitempty"`
	Type                string   `json:"type"`
	Region              string   `json:"region"`
	CanRequestInvite    bool     `json:"canRequestInvite"`
	Permanent           bool     `json:"permanent"`
	Platforms           struct {
		Android           int `json:"android"`
		StandaloneWindows int `json:"standalonewindows"`
	} `json:"platforms"`
}

// FriendStatus はフレンドステータスです
type FriendStatus struct {
	IsFriend        bool `json:"isFriend"`
	IncomingRequest bool `json:"incomingRequest"`
	OutgoingRequest bool `json:"outgoingRequest"`
}

// Notification は通知情報です
type Notification struct {
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

// TwoFactorAuthRequest は2FA検証リクエストです
type TwoFactorAuthRequest struct {
	Code string `json:"code"`
}

// TwoFactorAuthResponse は2FA検証レスポンスです
type TwoFactorAuthResponse struct {
	Verified bool `json:"verified"`
}

// Favorite はお気に入り情報です
type Favorite struct {
	ID         string   `json:"id"`
	Type       string   `json:"type"` // "world", "avatar", "friend"
	FavoriteID string   `json:"favoriteId"`
	Tags       []string `json:"tags"`
}

// FavoriteGroup はお気に入りグループです
type FavoriteGroup struct {
	ID         string   `json:"id"`
	Type       string   `json:"type"`
	OwnerID    string   `json:"ownerId"`
	Name       string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Visibility string   `json:"visibility"`
	Tags       []string `json:"tags"`
}

// Group はグループ情報です
type Group struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	ShortCode        string   `json:"shortCode"`
	Description      string   `json:"description"`
	IconURL          string   `json:"iconUrl"`
	BannerURL        string   `json:"bannerUrl"`
	Privacy          string   `json:"privacy"`
	OwnerID          string   `json:"ownerId"`
	Rules            string   `json:"rules"`
	Links            []string `json:"links"`
	Languages        []string `json:"languages"`
	Tags             []string `json:"tags"`
	MemberCount      int      `json:"memberCount"`
	OnlineMemberCount int     `json:"onlineMemberCount"`
	IsVerified       bool     `json:"isVerified"`
	JoinState        string   `json:"joinState"`
	MembershipStatus string   `json:"membershipStatus"`
	CreatedAt        string   `json:"createdAt"`
}

// GroupMember はグループメンバー情報です
type GroupMember struct {
	ID                 string      `json:"id"`
	GroupID            string      `json:"groupId"`
	UserID             string      `json:"userId"`
	RoleIDs            []string    `json:"roleIds"`
	User               *LimitedUser `json:"user,omitempty"`
	IsRepresenting     bool        `json:"isRepresenting"`
	MembershipStatus   string      `json:"membershipStatus"`
	Visibility         string      `json:"visibility"`
	JoinedAt           string      `json:"joinedAt"`
	CreatedAt          string      `json:"createdAt"`
	BannedAt           *string     `json:"bannedAt,omitempty"`
}

// GroupAnnouncement はグループのお知らせです
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

// File はファイル情報です
type File struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	MimeType string        `json:"mimeType"`
	Extension string       `json:"extension"`
	Tags     []string      `json:"tags"`
	Versions []FileVersion `json:"versions"`
}

// FileVersion はファイルのバージョン情報です
type FileVersion struct {
	Version   int                  `json:"version"`
	Status    string               `json:"status"`
	CreatedAt string               `json:"created_at"`
	File      *FileDescriptor      `json:"file,omitempty"`
	Signature *FileVersionSignature `json:"signature,omitempty"`
	Delta     *FileVersionDelta    `json:"delta,omitempty"`
}

// FileDescriptor はファイル記述子です
type FileDescriptor struct {
	URL        string `json:"url"`
	URLObject  string `json:"urlObject"`
	MD5        string `json:"md5"`
	SizeInMB   int    `json:"sizeInMb"`
	Status     string `json:"status"`
}

// FileVersionSignature はファイルバージョンの署名です
type FileVersionSignature struct {
	Signature string `json:"signature"`
	Algorithm string `json:"algorithm"`
}

// FileVersionDelta はファイルバージョンのデルタです
type FileVersionDelta struct {
	Server   string `json:"server"`
	Signature string `json:"signature"`
}

// PlayerModeration はプレイヤーモデレーション情報です
type PlayerModeration struct {
	ID                 string `json:"id"`
	Type               string `json:"type"` // "mute", "unmute", "block", "unblock", "hideAvatar", "showAvatar"
	SourceUserID       string `json:"sourceUserId"`
	SourceDisplayName  string `json:"sourceDisplayName"`
	TargetUserID       string `json:"targetUserId"`
	TargetDisplayName  string `json:"targetDisplayName"`
	Created            string `json:"created"`
}

// Announcement はお知らせです
type Announcement struct {
	ID        string `json:"id"`
	ReleaseStatus string `json:"releaseStatus"`
	Priority  int    `json:"priority"`
	Tags      []string `json:"tags"`
	Data      json.RawMessage `json:"data"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// DownloadUrls はダウンロードURLです
type DownloadUrls struct {
	SDK        string `json:"sdk"`
	VRChat     string `json:"vrchat"`
	Bootstrap  string `json:"bootstrap"`
}

// DynamicWorldRow は動的ワールド行です
type DynamicWorldRow struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Worlds []string `json:"worlds"`
}

// Config はシステム設定です
type Config struct {
	// システム設定に関する80以上のフィールド
	// ここでは基本的なフィールドのみ記載
	Address              string             `json:"address"`
	APIKey               string             `json:"apiKey"`
	AppName              string             `json:"appName"`
	BuildVersionTag      string             `json:"buildVersionTag"`
	ClientAPIKey         string             `json:"clientApiKey"`
	CurrentTOSVersion    int                `json:"currentTOSVersion"`
	DefaultAvatar        string             `json:"defaultAvatar"`
	DeploymentGroup      string             `json:"deploymentGroup"`
	Announcements        []Announcement     `json:"announcements"`
	DownloadUrls         DownloadUrls       `json:"downloadUrls"`
	DynamicWorldRows     []DynamicWorldRow  `json:"dynamicWorldRows"`
}

// InfoPush は情報プッシュです
type InfoPush struct {
	ID            string       `json:"id"`
	ReleaseStatus string       `json:"releaseStatus"`
	Priority      int          `json:"priority"`
	Tags          []string     `json:"tags"`
	Data          InfoPushData `json:"data"`
	StartDate     string       `json:"startDate"`
	EndDate       string       `json:"endDate"`
}

// InfoPushData は情報プッシュのデータです
type InfoPushData struct {
	Content string `json:"content"`
}

// SearchUsersOptions はユーザー検索のオプションです
type SearchUsersOptions struct {
	Search        string
	DeveloperType string
	N             int
	Offset        int
	Fuzzy         bool
}

// SearchAvatarsOptions はアバター検索のオプションです
type SearchAvatarsOptions struct {
	Featured        bool
	Tag             string
	UserID          string
	N               int
	Offset          int
	Order           string
	Sort            string
	ReleaseStatus   string
	MaxUnityVersion string
	MinUnityVersion string
	Platform        string
}

// SearchWorldsOptions はワールド検索のオプションです
type SearchWorldsOptions struct {
	Featured        bool
	Sort            string
	N               int
	Order           string
	Offset          int
	Search          string
	Tag             string
	UserID          string
	ReleaseStatus   string
	MaxUnityVersion string
	MinUnityVersion string
	Platform        string
}

// GetFriendsOptions はフレンドリスト取得のオプションです
type GetFriendsOptions struct {
	Offset  int
	N       int
	Offline bool
}

// GetNotificationsOptions は通知取得のオプションです
type GetNotificationsOptions struct {
	Type   string
	Sent   bool
	Hidden bool
	After  string
	N      int
	Offset int
}

// CreateInstanceRequest はインスタンス作成リクエストです
type CreateInstanceRequest struct {
	WorldID          string `json:"worldId"`
	Type             string `json:"type"`
	Region           string `json:"region"`
	InstanceID       string `json:"instanceId,omitempty"`
	OwnerID          string `json:"ownerId,omitempty"`
	GroupAccessType  string `json:"groupAccessType,omitempty"`
	QueueEnabled     bool   `json:"queueEnabled,omitempty"`
	CanRequestInvite bool   `json:"canRequestInvite,omitempty"`
}

// UpdateInstanceRequest はインスタンス更新リクエストです
type UpdateInstanceRequest struct {
	Hardclose        *bool   `json:"hardclose,omitempty"`
	CanRequestInvite *bool   `json:"canRequestInvite,omitempty"`
	Region           *string `json:"region,omitempty"`
}

// SendNotificationRequest は通知送信リクエストです
type SendNotificationRequest struct {
	Type       string                 `json:"type"`
	UserID     string                 `json:"userId"`
	Message    string                 `json:"message,omitempty"`
	Details    map[string]interface{} `json:"details,omitempty"`
}
