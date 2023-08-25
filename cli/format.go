package cli

import "strings"

func formatCmd(c string) []string {
	c = strings.ToLower(c)
	return strings.Fields(c)
}
