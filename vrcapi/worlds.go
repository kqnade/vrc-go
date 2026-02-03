package vrcapi

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kqnade/vrcgo/shared"
)

// GetWorld は指定されたワールドIDのワールド情報を取得します
func (c *Client) GetWorld(ctx context.Context, worldID string) (*shared.World, error) {
	var world shared.World
	err := c.doRequest(ctx, "GET", "/worlds/"+worldID, nil, &world)
	if err != nil {
		return nil, fmt.Errorf("failed to get world: %w", err)
	}
	return &world, nil
}

// SearchWorlds はワールドを検索します
func (c *Client) SearchWorlds(ctx context.Context, opts shared.SearchWorldsOptions) ([]shared.LimitedWorld, error) {
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

	var worlds []shared.LimitedWorld
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
