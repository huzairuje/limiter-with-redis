package utils

import "regexp"

const (
	ErrorLogFormat = "got err: %v, context: %s - %s"
)

func IsValidSanitizeSQL(queryParam string) bool {
	regexQueryParam := regexp.MustCompile(`^[\w ]+$`)
	return regexQueryParam.MatchString(queryParam)
}
