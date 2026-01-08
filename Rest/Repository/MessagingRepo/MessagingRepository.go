package messagingrepo

import (
	"time"

	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

type MessagingRepository interface {
	// Conversation operations
	CreateConversation(conv *models.Conversation) error
	GetConversationByID(id uint) (*models.Conversation, error)
	GetConversationByParticipants(p1ID uint, p1Type string, p2ID uint, p2Type string, contextType string, contextID *uint) (*models.Conversation, error)
	GetUserConversations(userID uint, userType string) ([]models.Conversation, error)
	UpdateConversation(conv *models.Conversation) error
	IncrementUnreadCount(convID uint, participantNumber int) error
	ResetUnreadCount(convID uint, participantNumber int) error

	// Message operations
	CreateMessage(msg *models.Message) error
	GetMessagesByConversationID(convID uint, since time.Time, limit int) ([]models.Message, error)
	GetMessageByID(id uint) (*models.Message, error)
	MarkMessagesAsRead(convID uint, userID uint, userType string) error
	GetUnreadMessagesCount(userID uint, userType string) (int64, error)

	// User info retrieval
	GetBuyerInfo(buyerID uint) (*models.Buyer, error)
	GetSellerInfo(sellerID uint) (*models.User, error)
	GetAdminInfo(adminID uint) (*models.Admin, error)
	GetRegUserInfo(userID uint) (*models.RegUser, error)

	// Context info
	GetProductInfo(productID uint) (*models.Product, error)
	GetOrderInfo(orderID uint) (*models.Order, error)
}
type messagingRepositoryImpl struct {
	db *gorm.DB
}

func NewMessagingRepository(db *gorm.DB) MessagingRepository {
	return &messagingRepositoryImpl{
		db: db,
	}
}
