package helpers

import (
	"regexp"
	"strings"
)

func Slugify(str string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	str = reg.ReplaceAllString(str, "-")

	str = strings.Trim(str, "-")
	str = strings.ToLower(str)

	return str
}
