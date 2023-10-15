package stringutil

import (
	"strings"
)

const ellipsis = "…"

func Trim(s string, max int) string {
	if s == "" {
		return s
	}

	var sb strings.Builder

	sb.Grow(len(s))
	sb.Grow(len(ellipsis))

	trimAt := min(len(s), max+len(ellipsis))

	sb.WriteString(s[:trimAt])
	sb.WriteString(ellipsis)

	return sb.String()
}
