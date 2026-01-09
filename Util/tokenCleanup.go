package util

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func RunTokenCleanup(db *gorm.DB) {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			if err := CleanupExpiredTokens(db); err != nil {
				log.Printf("Failed to cleanup expired tokens: %v", err)
			} else {
				log.Println("✓ Expired tokens cleaned up successfully")
			}
		}
	}()
}
