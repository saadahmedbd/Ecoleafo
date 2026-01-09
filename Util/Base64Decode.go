package util

import (
	"encoding/base64"
	"strings"
)

// Base64 URL decoding
func Base64UrlDecode(data string) ([]byte, error) {
	// add padding back
	padding := 4 - (len(data) % 4)
	if padding < 4 {
		data += strings.Repeat("=", padding)
	}
	return base64.URLEncoding.DecodeString(data)
}
func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
