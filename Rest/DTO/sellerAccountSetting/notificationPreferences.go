package selleraccountsetting

// NotificationPreferences - Notification settings
type NotificationPreferences struct {
	OrderEmail     bool `json:"order_email"`
	OrderSMS       bool `json:"order_sms"`
	OrderPush      bool `json:"order_push"`
	MessageEmail   bool `json:"message_email"`
	MessageSMS     bool `json:"message_sms"`
	MessagePush    bool `json:"message_push"`
	MarketingEmail bool `json:"marketing_email"`
	MarketingSMS   bool `json:"marketing_sms"`
	MarketingPush  bool `json:"marketing_push"`
}
