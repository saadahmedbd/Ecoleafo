package selleraccountsettingservice

import selleraccountsetting "github.com/saadahmedbd/Treestore/Rest/DTO/sellerAccountSetting"

// GetNotificationPreferences retrieves notification settings
func (s *SellerAccountSettingService) GetNotificationPreferences(sellerID uint) (*selleraccountsetting.NotificationPreferences, error) {
	prefs, err := s.repo.GetNotificationPreferences(sellerID)
	if err != nil {
		return nil, err
	}

	return &selleraccountsetting.NotificationPreferences{
		OrderEmail:     prefs.OrderEmail,
		OrderSMS:       prefs.OrderSMS,
		OrderPush:      prefs.OrderPush,
		MessageEmail:   prefs.MessageEmail,
		MessageSMS:     prefs.MessageSMS,
		MessagePush:    prefs.MessagePush,
		MarketingEmail: prefs.MarketingEmail,
		MarketingSMS:   prefs.MarketingSMS,
		MarketingPush:  prefs.MarketingPush,
	}, nil
}
