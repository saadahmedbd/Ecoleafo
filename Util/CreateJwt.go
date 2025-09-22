package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	Sub       int    `json:"sub"`
	UserId    uint   `json:"user_id"`
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Exp       int64  `json:"exp"`
	Iat       int64  `json:"iat"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncode(byteArrHeader)
	now := time.Now().Unix()
	payload := Payload{
		Sub:       data.Sub,
		UserId:    data.UserId,
		Role:      data.Role,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Exp:       now + 60*60*24,
		Iat:       now,
	}

	byteArrData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	PayloadB64 := base64UrlEncode(byteArrData)
	message := headerB64 + "." + PayloadB64

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)
	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)
	jwt := headerB64 + "." + PayloadB64 + "." + signatureB64
	return jwt, nil

}
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
