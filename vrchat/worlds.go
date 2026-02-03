package vrchat

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// World はワールド情報です
type World struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	AuthorID             string   `json:"authorId"`
	AuthorName           string   `json:"authorName"`
	TotalLikes           int      `json:"likes"`
	TotalVisits          int      `json:"visits"`
	Capacity             int      `json:"capacity"`
	RecommendedCapacity  int      `json:"recommendedCapacity"`
	ImageURL             string   `json:"imageUrl"`
	ThumbnailImageURL    string   `json:"thumbnailImageUrl"`
	ReleaseStatus        string   `json:"releaseStatus"`
	Organization         string   `json:"organization"`
	Tags                 []string `json:"tags"`
	Favorites            int      `json:"favorites"`
	CreatedAt            string   `json:"created_at"`
	UpdatedAt            string   `json:"updated_at"`
	PublicationDate      string   `json:"publicationDate"`
	LabsPublicationDate  string   `json:"labsPublicationDate"`
	Instances            [][]interface{} `json:"instances"`
	PublicOccupants      int      `json:"publicOccupants"`
	PrivateOccupants     int      `json:"privateOccupants"`
	Occupants            int      `json:"occupants"`
	UnityPackages        []UnityPackage `json:"unityPackages"`
	Namespace            string   `json:"namespace"`
	Version              int      `json:"version"`
	PreviewYoutubeID     *string  `json:"previewYoutubeId"`
	UdonProducts         []string `json:"udonProducts"`
	Heat                 int      `json:"heat"`
}

// LimitedWorld は制限されたワールド情報です
type LimitedWorld struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	AuthorID            string   `json:"authorId"`
	AuthorName          string   `json:"authorName"`
	TotalLikes          int      `json:"likes"`
	Capacity            int      `json:"capacity"`
	RecommendedCapacity int      `json:"recommendedCapacity"`
	ImageURL            string   `json:"imageUrl"`
	ThumbnailImageURL   string   `json:"thumbnailImageUrl"`
	ReleaseStatus       string   `json:"releaseStatus"`
	Organization        string   `json:"organization"`
	Tags                []string `json:"tags"`
	Favorites           int      `json:"favorites"`
	CreatedAt           string   `json:"created_at"`
	UpdatedAt           string   `json:"updated_at"`
	PublicationDate     string   `json:"publicationDate"`
	LabsPublicationDate string   `json:"labsPublicationDate"`
	Heat                int      `json:"heat"`
	PublicOccupants     int      `json:"publicOccupants"`
	PrivateOccupants    int      `json:"privateOccupants"`
	Occupants           int      `json:"occupants"`
	UnityPackages       []UnityPackage `json:"unityPackages"`
}

// GetWorld は指定されたワールドIDのワールド情報を取得します
func (c *Client) GetWorld(ctx context.Context, worldID string) (*World, error) {
	var world World
	err := c.doRequest(ctx, "GET", "/worlds/"+worldID, nil, &world)
	if err != nil {
		return nil, fmt.Errorf("failed to get world: %w", err)
	}
	return &world, nil
}

// SearchWorldsOptions はワールド検索のオプションです
type SearchWorldsOptions struct {
	Featured      bool   // 注目のワールドのみ
	Sort          string // ソート項目: "popularity", "heat", "trust", "shuffle", "random", "favorites", "reportScore", "reportCount", "publicationDate", "labsPublicationDate", "created", "_created_at", "updated", "_updated_at", "order", "relevance", "magic", "name"
	N             int    // 取得件数（デフォルト: 60）
	Order         string // ソート順: "ascending", "descending"
	Offset        int    // オフセット
	Search        string // 検索クエリ
	Tag           string // タグでフィルター
	UserID        string // 特定のユーザーのワールド
	ReleaseStatus string // リリースステータス: "public", "private", "hidden", "all"
	MaxUnityVersion string // 最大Unityバージョン
	MinUnityVersion string // 最小Unityバージョン
	Platform      string // プラットフォーム: "android", "standalonewindows"
}

// SearchWorlds はワールドを検索します
func (c *Client) SearchWorlds(ctx context.Context, opts SearchWorldsOptions) ([]LimitedWorld, error) {
	params := url.Values{}
	if opts.Featured {
		params.Set("featured", "true")
	}
	if opts.Sort != "" {
		params.Set("sort", opts.Sort)
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}
	if opts.Order != "" {
		params.Set("order", opts.Order)
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Search != "" {
		params.Set("search", opts.Search)
	}
	if opts.Tag != "" {
		params.Set("tag", opts.Tag)
	}
	if opts.UserID != "" {
		params.Set("userId", opts.UserID)
	}
	if opts.ReleaseStatus != "" {
		params.Set("releaseStatus", opts.ReleaseStatus)
	}
	if opts.MaxUnityVersion != "" {
		params.Set("maxUnityVersion", opts.MaxUnityVersion)
	}
	if opts.MinUnityVersion != "" {
		params.Set("minUnityVersion", opts.MinUnityVersion)
	}
	if opts.Platform != "" {
		params.Set("platform", opts.Platform)
	}

	var worlds []LimitedWorld
	path := "/worlds"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &worlds)
	if err != nil {
		return nil, fmt.Errorf("failed to search worlds: %w", err)
	}
	return worlds, nil
}

// GetActiveWorlds はアクティブなワールドを取得します
func (c *Client) GetActiveWorlds(ctx context.Context) ([]LimitedWorld, error) {
	return c.SearchWorlds(ctx, SearchWorldsOptions{
		Sort: "heat",
		N:    100,
	})
}

// GetRecentWorlds は最近訪問したワールドを取得します
func (c *Client) GetRecentWorlds(ctx context.Context) ([]LimitedWorld, error) {
	var worlds []LimitedWorld
	err := c.doRequest(ctx, "GET", "/worlds/recent", nil, &worlds)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent worlds: %w", err)
	}
	return worlds, nil
}

// GetFavoriteWorlds はお気に入りのワールドを取得します
func (c *Client) GetFavoriteWorlds(ctx context.Context) ([]LimitedWorld, error) {
	return c.SearchWorlds(ctx, SearchWorldsOptions{
		Tag: "favorite",
	})
}

// CreateWorldRequest はワールド作成リクエストです
type CreateWorldRequest struct {
	AssetURL        string `json:"assetUrl"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Tags            []string `json:"tags"`
	ImageURL        string `json:"imageUrl"`
	ReleaseStatus   string `json:"releaseStatus"`
	Capacity        int    `json:"capacity"`
	RecommendedCapacity int `json:"recommendedCapacity"`
	UnityPackageURL string `json:"unityPackageUrl"`
}

// CreateWorld は新しいワールドを作成します
func (c *Client) CreateWorld(ctx context.Context, req CreateWorldRequest) (*World, error) {
	var world World
	err := c.doRequest(ctx, "POST", "/worlds", req, &world)
	if err != nil {
		return nil, fmt.Errorf("failed to create world: %w", err)
	}
	return &world, nil
}

// UpdateWorldRequest はワールド更新リクエストです
type UpdateWorldRequest struct {
	AssetURL            *string  `json:"assetUrl,omitempty"`
	Name                *string  `json:"name,omitempty"`
	Description         *string  `json:"description,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	ImageURL            *string  `json:"imageUrl,omitempty"`
	ReleaseStatus       *string  `json:"releaseStatus,omitempty"`
	Capacity            *int     `json:"capacity,omitempty"`
	RecommendedCapacity *int     `json:"recommendedCapacity,omitempty"`
	UnityPackageURL     *string  `json:"unityPackageUrl,omitempty"`
}

// UpdateWorld はワールド情報を更新します
func (c *Client) UpdateWorld(ctx context.Context, worldID string, req UpdateWorldRequest) (*World, error) {
	var world World
	err := c.doRequest(ctx, "PUT", "/worlds/"+worldID, req, &world)
	if err != nil {
		return nil, fmt.Errorf("failed to update world: %w", err)
	}
	return &world, nil
}

// DeleteWorld はワールドを削除します
func (c *Client) DeleteWorld(ctx context.Context, worldID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/worlds/"+worldID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete world: %w", err)
	}
	return nil
}

// WorldMetadata はワールドのメタデータです
type WorldMetadata struct {
	ID       string `json:"id"`
	Metadata struct {
		Capacity int `json:"capacity"`
		RecommendedCapacity int `json:"recommendedCapacity"`
	} `json:"metadata"`
}

// GetWorldMetadata はワールドのメタデータを取得します
func (c *Client) GetWorldMetadata(ctx context.Context, worldID string) (*WorldMetadata, error) {
	var metadata WorldMetadata
	err := c.doRequest(ctx, "GET", "/worlds/"+worldID+"/metadata", nil, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to get world metadata: %w", err)
	}
	return &metadata, nil
}

// PublishWorld はワールドを公開します
func (c *Client) PublishWorld(ctx context.Context, worldID string) (*World, error) {
	var world World
	err := c.doRequest(ctx, "PUT", "/worlds/"+worldID+"/publish", nil, &world)
	if err != nil {
		return nil, fmt.Errorf("failed to publish world: %w", err)
	}
	return &world, nil
}

// UnpublishWorld はワールドを非公開にします
func (c *Client) UnpublishWorld(ctx context.Context, worldID string) (*World, error) {
	var world World
	err := c.doRequest(ctx, "DELETE", "/worlds/"+worldID+"/publish", nil, &world)
	if err != nil {
		return nil, fmt.Errorf("failed to unpublish world: %w", err)
	}
	return &world, nil
}
