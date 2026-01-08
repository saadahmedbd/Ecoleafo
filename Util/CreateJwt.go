package util

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/saadahmedbd/Treestore/Config"
)

// type Payload struct {
// 	Sub       int    `json:"sub"`
// 	UserId    uint   `json:"user_id"`
// 	Role      string `json:"role"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Exp       int64  `json:"exp"`
// 	Iat       int64  `json:"iat"`
// }

// func CreateJwt(userID uint, roles []string, ttl time.Duration) (string, error) {
// 	header := map[string]interface{}{
// 		"alg": "HS256",
// 		"typ": "JWT",
// 	}
// 	byteArrHeader, err := json.Marshal(header)
// 	if err != nil {
// 		return "", err
// 	}
// 	headerB64 := Base64UrlEncode(byteArrHeader)

// 	payload := map[string]interface{}{
// 		"user_id": userID,
// 		"roles":   roles,
// 		"exp":     time.Now().Add(ttl).Unix(),
// 	}
// 	payloadJSON, _ := json.Marshal(payload)
// 	PayloadB64 := Base64UrlEncode(payloadJSON)

// 	cnf := Config.GetConfig()

// 	unsigned := headerB64 + "." + PayloadB64
// 	signature := signHS256(unsigned, []byte(cnf.JwtSecretKey))
// 	return unsigned + "." + signature, nil

// }

// byteArrData, err := json.Marshal(payload)
// if err != nil {
// 	return "", err
// }
// PayloadB64 := base64UrlEncode(byteArrData)
// message := headerB64 + "." + PayloadB64

// byteArrSecret := []byte(secret)
// byteArrMessage := []byte(message)

// h := hmac.New(sha256.New, byteArrSecret)
// h.Write(byteArrMessage)
// signature := h.Sum(nil)
// signatureB64 := base64UrlEncode(signature)
// jwt := headerB64 + "." + PayloadB64 + "." + signatureB64
// return jwt, nil

// built in

const (
	AccessTokenTTL  = 15 * time.Minute   // Short-lived for security
	RefreshTokenTTL = 7 * 24 * time.Hour // Long-lived refresh token
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"` // seconds until access token expires
}

func CreateJwt(userID uint, firstname, lastname string, roles []string, roleID uint, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userID,
		"first_name": firstname,
		"last_name":  lastname,
		"role":       roles,
		"role_id":    roleID,
		"exp":        time.Now().Add(ttl).Unix(),
		"iat":        time.Now().Unix(),
		"type":       "access", // Distinguish token type
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Config.GetConfig().JwtSecretKey))

}
