package util

import (
	"fmt"
	"time"
)

// Helper function to format time ago
func FormatTimeAgo(t time.Time) string {
	duration := time.Since(t)

	if duration.Hours() < 1 {
		return "Just now"
	} else if duration.Hours() < 24 {
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else if duration.Hours() < 168 { // 7 days
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	} else {
		weeks := int(duration.Hours() / 168)
		if weeks == 1 {
			return "1 week ago"
		}
		return fmt.Sprintf("%d weeks ago", weeks)
	}
}
