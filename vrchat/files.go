package vrchat

import (
	"context"
	"fmt"
)

// File はファイル情報です
type File struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	MimeType  string            `json:"mimeType"`
	Extension string            `json:"extension"`
	Tags      []string          `json:"tags"`
	Versions  []FileVersion     `json:"versions"`
}

// FileVersion はファイルバージョン情報です
type FileVersion struct {
	Version    int                    `json:"version"`
	Status     string                 `json:"status"`
	CreatedAt  string                 `json:"created_at"`
	File       *FileDescriptor        `json:"file,omitempty"`
	Signature  *FileVersionSignature  `json:"signature,omitempty"`
	Delta      *FileVersionDelta      `json:"delta,omitempty"`
}

// FileDescriptor はファイル記述子です
type FileDescriptor struct {
	MD5      string            `json:"md5"`
	SizeInBytes int64          `json:"sizeInBytes"`
	URL      string            `json:"url"`
	UploadID string            `json:"uploadId"`
	FileName string            `json:"fileName"`
}

// FileVersionSignature はファイルバージョンの署名です
type FileVersionSignature struct {
	MD5      string `json:"md5"`
	SizeInBytes int64 `json:"sizeInBytes"`
}

// FileVersionDelta はファイルバージョンの差分です
type FileVersionDelta struct {
	MD5      string `json:"md5"`
	SizeInBytes int64 `json:"sizeInBytes"`
}

// GetFile は指定されたファイルIDのファイル情報を取得します
func (c *Client) GetFile(ctx context.Context, fileID string) (*File, error) {
	var file File
	err := c.doRequest(ctx, "GET", "/file/"+fileID, nil, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}
	return &file, nil
}

// CreateFileRequest はファイル作成リクエストです
type CreateFileRequest struct {
	Name      string   `json:"name"`
	MimeType  string   `json:"mimeType"`
	Extension string   `json:"extension"`
	Tags      []string `json:"tags,omitempty"`
}

// CreateFile は新しいファイルを作成します
func (c *Client) CreateFile(ctx context.Context, req CreateFileRequest) (*File, error) {
	var file File
	err := c.doRequest(ctx, "POST", "/file", req, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	return &file, nil
}

// DeleteFile はファイルを削除します
func (c *Client) DeleteFile(ctx context.Context, fileID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/file/"+fileID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

// DownloadFile はファイルのダウンロードURLを取得します
func (c *Client) DownloadFile(ctx context.Context, fileID string, version int) (string, error) {
	var response struct {
		URL string `json:"url"`
	}
	path := fmt.Sprintf("/file/%s/%d", fileID, version)
	err := c.doRequest(ctx, "GET", path, nil, &response)
	if err != nil {
		return "", fmt.Errorf("failed to get download URL: %w", err)
	}
	return response.URL, nil
}

// CreateFileVersionRequest はファイルバージョン作成リクエストです
type CreateFileVersionRequest struct {
	SignatureMD5      string `json:"signatureMd5"`
	SignatureSizeInBytes int64 `json:"signatureSizeInBytes"`
	FileMD5           string `json:"fileMd5,omitempty"`
	FileSizeInBytes   int64  `json:"fileSizeInBytes,omitempty"`
}

// CreateFileVersion は新しいファイルバージョンを作成します
func (c *Client) CreateFileVersion(ctx context.Context, fileID string, req CreateFileVersionRequest) (*File, error) {
	var file File
	err := c.doRequest(ctx, "POST", "/file/"+fileID, req, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to create file version: %w", err)
	}
	return &file, nil
}

// DeleteFileVersion はファイルバージョンを削除します
func (c *Client) DeleteFileVersion(ctx context.Context, fileID string, version int) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	path := fmt.Sprintf("/file/%s/%d", fileID, version)
	err := c.doRequest(ctx, "DELETE", path, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete file version: %w", err)
	}
	return nil
}

// FinishFileDataUpload はファイルデータアップロードを完了します
func (c *Client) FinishFileDataUpload(ctx context.Context, fileID string, version int, uploadType string) (*File, error) {
	req := struct {
		ETag         string `json:"etag,omitempty"`
		MaxParts     int    `json:"maxParts,omitempty"`
		NextPartNumber int  `json:"nextPartNumber,omitempty"`
		Parts        []struct {
			ETag       string `json:"etag"`
			PartNumber int    `json:"partNumber"`
		} `json:"parts,omitempty"`
	}{}

	var file File
	path := fmt.Sprintf("/file/%s/%d/%s/finish", fileID, version, uploadType)
	err := c.doRequest(ctx, "PUT", path, req, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to finish file data upload: %w", err)
	}
	return &file, nil
}

// StartFileDataUpload はファイルデータアップロードを開始します
func (c *Client) StartFileDataUpload(ctx context.Context, fileID string, version int, uploadType string, partNumber int) (*FileDescriptor, error) {
	var descriptor FileDescriptor
	path := fmt.Sprintf("/file/%s/%d/%s/start", fileID, version, uploadType)
	if partNumber > 0 {
		path += fmt.Sprintf("?partNumber=%d", partNumber)
	}
	err := c.doRequest(ctx, "PUT", path, nil, &descriptor)
	if err != nil {
		return nil, fmt.Errorf("failed to start file data upload: %w", err)
	}
	return &descriptor, nil
}

// GetFileDataUploadStatus はファイルデータアップロードのステータスを取得します
func (c *Client) GetFileDataUploadStatus(ctx context.Context, fileID string, version int, uploadType string) (string, error) {
	var response struct {
		Status string `json:"status"`
	}
	path := fmt.Sprintf("/file/%s/%d/%s/status", fileID, version, uploadType)
	err := c.doRequest(ctx, "GET", path, nil, &response)
	if err != nil {
		return "", fmt.Errorf("failed to get file data upload status: %w", err)
	}
	return response.Status, nil
}
