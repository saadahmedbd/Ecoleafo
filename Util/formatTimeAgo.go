package util

import "time"

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
		return string(rune(hours)) + " hours ago"
	} else if duration.Hours() < 168 { // 7 days
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return string(rune(days)) + " days ago"
	} else {
		weeks := int(duration.Hours() / 168)
		if weeks == 1 {
			return "1 week ago"
		}
		return string(rune(weeks)) + " weeks ago"
	}
}
