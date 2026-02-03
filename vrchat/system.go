package vrchat

import (
	"context"
	"fmt"
	"time"
)

// Config はシステム設定情報です
type Config struct {
	Address                         string            `json:"address"`
	Announcements                   []Announcement    `json:"announcements"`
	APIKey                          string            `json:"apiKey"`
	AppName                         string            `json:"appName"`
	BuildVersionTag                 string            `json:"buildVersionTag"`
	ClientAPIKey                    string            `json:"clientApiKey"`
	ClientBPSCeiling                int               `json:"clientBPSCeiling"`
	ClientDisconnectTimeout         int               `json:"clientDisconnectTimeout"`
	ClientReservedPlayerBPS         int               `json:"clientReservedPlayerBPS"`
	ClientSentCountAllowance        int               `json:"clientSentCountAllowance"`
	ContactEmail                    string            `json:"contactEmail"`
	CopyrightEmail                  string            `json:"copyrightEmail"`
	CurrentTOSVersion               int               `json:"currentTOSVersion"`
	DefaultAvatar                   string            `json:"defaultAvatar"`
	DeploymentGroup                 string            `json:"deploymentGroup"`
	DevAppVersionStandalone         string            `json:"devAppVersionStandalone"`
	DevDownloadLinkWindows          string            `json:"devDownloadLinkWindows"`
	DevSdkUrl                       string            `json:"devSdkUrl"`
	DevSdkVersion                   string            `json:"devSdkVersion"`
	DevServerVersionStandalone      string            `json:"devServerVersionStandalone"`
	DisCountdown                    time.Time         `json:"dis-countdown"`
	DisableAvatarCopying            bool              `json:"disableAvatarCopying"`
	DisableAvatarGating             bool              `json:"disableAvatarGating"`
	DisableCommunityLabs            bool              `json:"disableCommunityLabs"`
	DisableCommunityLabsPromotion   bool              `json:"disableCommunityLabsPromotion"`
	DisableEmail                    bool              `json:"disableEmail"`
	DisableEventStream              bool              `json:"disableEventStream"`
	DisableFeedbackGating           bool              `json:"disableFeedbackGating"`
	DisableFrontendBuilds           bool              `json:"disableFrontendBuilds"`
	DisableHello                    bool              `json:"disableHello"`
	DisableOculusSubs               bool              `json:"disableOculusSubs"`
	DisableRegistration             bool              `json:"disableRegistration"`
	DisableSteamNetworking          bool              `json:"disableSteamNetworking"`
	DisableTwoFactorAuth            bool              `json:"disableTwoFactorAuth"`
	DisableUdon                     bool              `json:"disableUdon"`
	DisableUpgradeAccount           bool              `json:"disableUpgradeAccount"`
	DownloadLinkWindows             string            `json:"downloadLinkWindows"`
	DownloadUrls                    DownloadUrls      `json:"downloadUrls"`
	DynamicWorldRows                []DynamicWorldRow `json:"dynamicWorldRows"`
	Events                          Events            `json:"events"`
	GearDemoRoomID                  string            `json:"gearDemoRoomId"`
	HomeWorldID                     string            `json:"homeWorldId"`
	HomepageRedirectTarget          string            `json:"homepageRedirectTarget"`
	HubWorldID                      string            `json:"hubWorldId"`
	JobsEmail                       string            `json:"jobsEmail"`
	MessageOfTheDay                 string            `json:"messageOfTheDay"`
	ModerationEmail                 string            `json:"moderationEmail"`
	ModerationQueryPeriod           int               `json:"moderationQueryPeriod"`
	NotAllowedToSelectAvatarInPrivateWorldMessage string `json:"notAllowedToSelectAvatarInPrivateWorldMessage"`
	Plugin                          string            `json:"plugin"`
	ReleaseAppVersionStandalone     string            `json:"releaseAppVersionStandalone"`
	ReleaseSdkUrl                   string            `json:"releaseSdkUrl"`
	ReleaseSdkVersion               string            `json:"releaseSdkVersion"`
	ReleaseServerVersionStandalone  string            `json:"releaseServerVersionStandalone"`
	SdkDeveloperFaqUrl              string            `json:"sdkDeveloperFaqUrl"`
	SdkDiscordUrl                   string            `json:"sdkDiscordUrl"`
	SdkNotAllowedToPublishMessage   string            `json:"sdkNotAllowedToPublishMessage"`
	SdkUnityVersion                 string            `json:"sdkUnityVersion"`
	ServerName                      string            `json:"serverName"`
	SupportEmail                    string            `json:"supportEmail"`
	TimeOutWorldID                  string            `json:"timeOutWorldId"`
	TutorialWorldID                 string            `json:"tutorialWorldId"`
	UpdateRateMsMaximum             int               `json:"updateRateMsMaximum"`
	UpdateRateMsMinimum             int               `json:"updateRateMsMinimum"`
	UpdateRateMsNormal              int               `json:"updateRateMsNormal"`
	UpdateRateMsUdonManual          int               `json:"updateRateMsUdonManual"`
	UploadAnalysisPercent           int               `json:"uploadAnalysisPercent"`
	UrlList                         []string          `json:"urlList"`
	UseReliableUdpHost              bool              `json:"useReliableUdpHost"`
	UserUpdatePeriod                int               `json:"userUpdatePeriod"`
	UserVerificationDelay           int               `json:"userVerificationDelay"`
	UserVerificationRetry           int               `json:"userVerificationRetry"`
	UserVerificationTimeout         int               `json:"userVerificationTimeout"`
	ViveWindowsUrl                  string            `json:"viveWindowsUrl"`
	WhiteListedAssetUrls            []string          `json:"whiteListedAssetUrls"`
	WorldUpdatePeriod               int               `json:"worldUpdatePeriod"`
}

// Announcement はアナウンス情報です
type Announcement struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

// DownloadUrls はダウンロードURL情報です
type DownloadUrls struct {
	Sdk2         string `json:"sdk2"`
	Sdk3Avatars  string `json:"sdk3-avatars"`
	Sdk3Worlds   string `json:"sdk3-worlds"`
	VCC          string `json:"vcc"`
	Bootstrap    string `json:"bootstrap"`
}

// DynamicWorldRow は動的ワールド行情報です
type DynamicWorldRow struct {
	Index       int    `json:"index"`
	Name        string `json:"name"`
	Platform    string `json:"platform"`
	SortHeading string `json:"sortHeading"`
	SortOrder   string `json:"sortOrder"`
	SortOwnership string `json:"sortOwnership"`
	Tag         string `json:"tag"`
	Type        string `json:"type"`
}

// Events はイベント情報です
type Events struct {
	DistanceClose  int `json:"distanceClose"`
	DistanceFactor int `json:"distanceFactor"`
	DistanceFar    int `json:"distanceFar"`
	GroupDistance  int `json:"groupDistance"`
	MaximumBunchSize int `json:"maximumBunchSize"`
	NotVisibleFactor int `json:"notVisibleFactor"`
	PlayerOrderBucketSize int `json:"playerOrderBucketSize"`
	PlayerOrderFactor int `json:"playerOrderFactor"`
	SlowUpdateFactorThreshold int `json:"slowUpdateFactorThreshold"`
	ViewSegmentLength int `json:"viewSegmentLength"`
}

// GetConfig はシステム設定情報を取得します
func (c *Client) GetConfig(ctx context.Context) (*Config, error) {
	var config Config
	err := c.doRequest(ctx, "GET", "/config", nil, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	return &config, nil
}

// GetSystemTime はシステム時刻を取得します
func (c *Client) GetSystemTime(ctx context.Context) (time.Time, error) {
	var response struct {
		ServerTime string `json:"serverTime"`
	}
	err := c.doRequest(ctx, "GET", "/time", nil, &response)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get system time: %w", err)
	}

	t, err := time.Parse(time.RFC3339, response.ServerTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse server time: %w", err)
	}
	return t, nil
}

// InfoPush は情報プッシュです
type InfoPush struct {
	ID               string `json:"id"`
	IsEnabled        bool   `json:"isEnabled"`
	ReleaseStatus    string `json:"releaseStatus"`
	Priority         int    `json:"priority"`
	Tags             []string `json:"tags"`
	Data             InfoPushData `json:"data"`
	Hash             string `json:"hash"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
}

// InfoPushData は情報プッシュデータです
type InfoPushData struct {
	Article          string `json:"article"`
	ContentList      []string `json:"contentList"`
	Description      string `json:"description"`
	ImageUrl         string `json:"imageUrl"`
	Name             string `json:"name"`
	OnPressed        InfoPushOnPressed `json:"onPressed"`
	Template         string `json:"template"`
	Version          string `json:"version"`
}

// InfoPushOnPressed は情報プッシュの押下時動作です
type InfoPushOnPressed struct {
	Action string `json:"action"`
	Data   string `json:"data"`
	Market string `json:"market"`
}

// GetInfoPushes は情報プッシュのリストを取得します
func (c *Client) GetInfoPushes(ctx context.Context) ([]InfoPush, error) {
	var infoPushes []InfoPush
	err := c.doRequest(ctx, "GET", "/infoPush", nil, &infoPushes)
	if err != nil {
		return nil, fmt.Errorf("failed to get info pushes: %w", err)
	}
	return infoPushes, nil
}

// GetCurrentOnlineUsers は現在のオンラインユーザー数を取得します
func (c *Client) GetCurrentOnlineUsers(ctx context.Context) (int, error) {
	var response struct {
		OnlineUsers int `json:"onlineUsers"`
	}
	err := c.doRequest(ctx, "GET", "/visits", nil, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to get current online users: %w", err)
	}
	return response.OnlineUsers, nil
}

// GetHealth はAPIヘルスチェックを実行します
func (c *Client) GetHealth(ctx context.Context) (bool, error) {
	var response struct {
		OK bool `json:"ok"`
	}
	err := c.doRequest(ctx, "GET", "/health", nil, &response)
	if err != nil {
		return false, fmt.Errorf("failed to check health: %w", err)
	}
	return response.OK, nil
}
