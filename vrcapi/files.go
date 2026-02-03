package vrcapi

import (
	"context"
	"fmt"

	"github.com/kqnade/vrcgo/shared"
)

// GetFile は指定されたファイルIDのファイル情報を取得します
func (c *Client) GetFile(ctx context.Context, fileID string) (*shared.File, error) {
	var file shared.File
	err := c.doRequest(ctx, "GET", "/file/"+fileID, nil, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}
	return &file, nil
}

// CreateFile はファイルを作成します
func (c *Client) CreateFile(ctx context.Context, req interface{}) (*shared.File, error) {
	var file shared.File
	err := c.doRequest(ctx, "POST", "/file", req, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	return &file, nil
}

// CreateFileVersion はファイルの新しいバージョンを作成します
func (c *Client) CreateFileVersion(ctx context.Context, fileID string, req interface{}) (*shared.File, error) {
	var file shared.File
	err := c.doRequest(ctx, "POST", "/file/"+fileID, req, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to create file version: %w", err)
	}
	return &file, nil
}

// DeleteFile はファイルを削除します
func (c *Client) DeleteFile(ctx context.Context, fileID string) error {
	err := c.doRequest(ctx, "DELETE", "/file/"+fileID, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

// DeleteFileVersion はファイルバージョンを削除します
func (c *Client) DeleteFileVersion(ctx context.Context, fileID string, versionID int) error {
	path := fmt.Sprintf("/file/%s/%d", fileID, versionID)
	err := c.doRequest(ctx, "DELETE", path, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete file version: %w", err)
	}
	return nil
}

// FinishFileDataUpload はファイルデータのアップロードを完了します
func (c *Client) FinishFileDataUpload(ctx context.Context, fileID string, versionID int, req interface{}) (*shared.File, error) {
	var file shared.File
	path := fmt.Sprintf("/file/%s/%d/file/finish", fileID, versionID)
	err := c.doRequest(ctx, "PUT", path, req, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to finish file data upload: %w", err)
	}
	return &file, nil
}
