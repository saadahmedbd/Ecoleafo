package selleraccountsetting

import "time"

// LoginActivity - Single login activity record
type LoginActivity struct {
	ID        uint      `json:"id"`
	Device    string    `json:"device"`
	Browser   string    `json:"browser"`
	Location  string    `json:"location"`
	IPAddress string    `json:"ip_address"`
	TimeAgo   string    `json:"time_ago"`
	IsCurrent bool      `json:"is_current"`
	CreatedAt time.Time `json:"created_at"`
}
