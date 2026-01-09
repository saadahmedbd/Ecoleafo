package selleraccountsettingservice

import (
	"fmt"
	"time"
)

// calculateTimeAgo calculates human-readable time difference
func (s *SellerAccountSettingService) calculateTimeAgo(t time.Time) string {
	duration := time.Since(t)

	if duration.Hours() < 1 {
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	} else if duration.Hours() < 24 {
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	} else if duration.Hours() < 168 { // 7 days
		return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
	} else {
		return t.Format("Jan 02, 2006")
	}
}
