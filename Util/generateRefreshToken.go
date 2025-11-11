package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

// GenerateRefreshToken creates a cryptographically secure random token
func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// CreateTokenPair generates both access and refresh tokens
func CreateTokenPair(db *gorm.DB, userID uint, firstname, lastname string, roles []string) (*TokenPair, error) {
	// Create access token
	accessToken, err := CreateJwt(userID, firstname, lastname, roles, AccessTokenTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	// Generate refresh token
	refreshTokenString, err := GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Store refresh token in database
	refreshToken := models.RefreshToken{
		UserID:    userID,
		Token:     refreshTokenString,
		ExpiresAt: time.Now().Add(RefreshTokenTTL),
		IsRevoked: false,
	}

	if err := db.Create(&refreshToken).Error; err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
		ExpiresIn:    int64(AccessTokenTTL.Seconds()),
	}, nil
}

// ValidateRefreshToken checks if refresh token is valid and not revoked
func ValidateRefreshToken(db *gorm.DB, tokenString string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken

	err := db.Where("token = ? AND is_revoked = ? AND expires_at > ?",
		tokenString, false, time.Now()).First(&refreshToken).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("invalid or expired refresh token")
		}
		return nil, err
	}

	return &refreshToken, nil
}

// RevokeRefreshToken marks a refresh token as revoked
func RevokeRefreshToken(db *gorm.DB, tokenString string) error {
	return db.Model(&models.RefreshToken{}).
		Where("token = ?", tokenString).
		Update("is_revoked", true).Error
}

// RevokeAllUserTokens revokes all refresh tokens for a user (useful for logout all devices)
func RevokeAllUserTokens(db *gorm.DB, userID uint) error {
	return db.Model(&models.RefreshToken{}).
		Where("user_id = ? AND is_revoked = ?", userID, false).
		Update("is_revoked", true).Error
}

// CleanupExpiredTokens removes expired tokens (run periodically)
func CleanupExpiredTokens(db *gorm.DB) error {
	return db.Where("expires_at < ?", time.Now()).Delete(&models.RefreshToken{}).Error
}
