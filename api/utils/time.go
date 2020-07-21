package utils

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const defaultLayout = time.RFC3339

// ParseTimestamp Converts timestamp to time by RFC3339 layout
func ParseTimestamp(timestamp string) (time.Time, error) {
	return ParseTimestampBy(defaultLayout, timestamp)
}

// ParseTimestamp Converts timestamp to time by custom layout
func ParseTimestampBy(layout, timestamp string) (time.Time, error) {
	if len(layout) == len(timestamp) {
		return time.Parse(layout, timestamp)
	}
	return time.Parse(defaultLayout, timestamp)
}

// ParseTimestamp Converts timestamp to time
func FormatTimestamp(timestamp time.Time) string {
	emptyTime := time.Time{}

	if timestamp != emptyTime {
		return timestamp.Format(time.RFC3339)
	}

	return ""
}

// FormatTime Converts kubernetes time to formatted timestamp
func FormatTime(time *metav1.Time) string {
	if time != nil {
		return FormatTimestamp(time.Time)
	}

	return ""
}
