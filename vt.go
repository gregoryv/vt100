package vt100

import (
	"regexp"
)

// Split returns a slice of strings split on vt100 escape codes
func Split(vt string) []string {
	re := regexp.MustCompile(`\033\[[0-9]+m`)
	parts := re.FindAllIndex([]byte(vt), -1)
	if len(parts) == 0 {
		return []string{vt}
	}
	result := make([]string, 0)
	firstStart := parts[0][0]
	if firstStart > 0 {
		// string up to first color
		result = append(result, vt[:firstStart])
	}
	var last []int
	for _, p := range parts {
		if last != nil {
			if last[0] < p[0] {
				// a string between two
				result = append(result, vt[last[0]:p[0]])
			}
		}
		// The code
		result = append(result, vt[p[0]:p[1]])
		last = p
	}
	if last[1] < len(vt)-1 {
		// Add remaining string
		result = append(result, vt[last[1]:])
	}
	return result
}
