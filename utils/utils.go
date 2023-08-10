package utils

import (
	"regexp"
	"time"
)

const (
	ErrorLogFormat = "got err: %v, context: %s - %s"
)

func IsValidSanitizeSQL(queryParam string) bool {
	regexQueryParam := regexp.MustCompile(`^[\w ]+$`)
	return regexQueryParam.MatchString(queryParam)
}

func StringUnitToDuration(input string) time.Duration {
	durationMapping := map[string]time.Duration{
		"second": time.Second,
		"minute": time.Minute,
		"hour":   time.Hour,
		"day":    24 * time.Hour,
		"week":   7 * 24 * time.Hour,
		"month":  30 * 24 * time.Hour,
		"year":   365 * 24 * time.Hour,
	}

	switch input {
	case "second":
		return durationMapping["second"]
	case "minute":
		return durationMapping["minute"]
	case "hour":
		return durationMapping["hour"]
	case "week":
		return durationMapping["week"]
	case "month":
		return durationMapping["month"]
	case "year":
		return durationMapping["year"]
	default:
		return time.Second
	}
}
