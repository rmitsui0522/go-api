package auth0

import (
	"context"
	"strings"
)

// トークンから取得したいクレーム
type CustomClaims struct {
	Scope string `json:"scope"`
}

// TODO: 実装の要否確認（インターフェースを満たすのみで問題ないか）
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// クレームに期待されるスコープが含まれているか検証する
func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")

	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}
