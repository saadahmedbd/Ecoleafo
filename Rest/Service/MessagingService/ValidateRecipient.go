package messagingservice

import "errors"

// Helper methods
func (s *messagingServiceImpl) validateRecipient(recipientID uint, recipientType string) error {
	switch recipientType {
	case "buyer":
		_, err := s.messagerepo.GetBuyerInfo(recipientID)
		return err
	case "seller":
		_, err := s.messagerepo.GetSellerInfo(recipientID)
		return err
	case "admin":
		_, err := s.messagerepo.GetAdminInfo(recipientID)
		return err
	default:
		return errors.New("invalid recipient type")
	}
}
