package selleraccountsettingservice

import (
	"errors"

	selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"
)

// UpdateNotificationPreferences updates notification settings
func (s *SellerAccountSettingService) UpdateNotificationPreferences(sellerID uint, req *selleraccountsetting.NotificationPreferences) error {
	updates := map[string]interface{}{
		"order_email":     req.OrderEmail,
		"order_sms":       req.OrderSMS,
		"order_push":      req.OrderPush,
		"message_email":   req.MessageEmail,
		"message_sms":     req.MessageSMS,
		"message_push":    req.MessagePush,
		"marketing_email": req.MarketingEmail,
		"marketing_sms":   req.MarketingSMS,
		"marketing_push":  req.MarketingPush,
	}

	if err := s.repo.UpdateNotificationPreferences(sellerID, updates); err != nil {
		return errors.New("failed to update notification preferences")
	}

	return nil
}
