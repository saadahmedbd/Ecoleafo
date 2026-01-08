package database

import (
	"log"

	"gorm.io/gorm"
)

func InitializeMessaging(db *gorm.DB) {
	// Add unique constraint for conversations
	db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_conversation 
		ON conversations (
			LEAST(participant1_id, participant2_id),
			GREATEST(participant1_id, participant2_id),
			COALESCE(context_type, ''),
			COALESCE(context_id, 0)
		) WHERE deleted_at IS NULL
	`)
	log.Println("✅ Messaging system initialized successfully")
}
