package utils

import "strings"

func FormatCmd(c string) []string {
	c = strings.ToLower(c)
	return strings.Fields(c)
}
