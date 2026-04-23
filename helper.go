package main

import (
	"time"
)

func parseTime(pubDate string) time.Time {
	for _, format := range []string{time.RFC1123, time.RFC822, time.RFC850, time.RFC1123Z, time.RFC3339} {
		if parsedTime, err := time.Parse(format, pubDate); err == nil {
			return parsedTime
		}
	}
	return time.Time{}
}
