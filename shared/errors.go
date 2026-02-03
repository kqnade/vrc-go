package shared

import "fmt"

// APIError はVRChat API固有のエラーです
type APIError struct {
	StatusCode int
	Message    string
	ErrorCode  string
}

func (e *APIError) Error() string {
	if e.ErrorCode != "" {
		return fmt.Sprintf("vrchat api error (status %d): %s [%s]",
			e.StatusCode, e.Message, e.ErrorCode)
	}
	return fmt.Sprintf("vrchat api error (status %d): %s",
		e.StatusCode, e.Message)
}

// IsAuthenticationError は認証エラーかどうかを判定します
func IsAuthenticationError(err error) bool {
	apiErr, ok := err.(*APIError)
	return ok && apiErr.StatusCode == 401
}

// IsRateLimitError はレート制限エラーかどうかを判定します
func IsRateLimitError(err error) bool {
	apiErr, ok := err.(*APIError)
	return ok && apiErr.StatusCode == 429
}

// IsNotFoundError はリソースが見つからないエラーかどうかを判定します
func IsNotFoundError(err error) bool {
	apiErr, ok := err.(*APIError)
	return ok && apiErr.StatusCode == 404
}
