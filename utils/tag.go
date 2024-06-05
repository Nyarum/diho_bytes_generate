package utils

import (
	"strings"
)

// ParseStructTag parses a struct tag and returns a map of keys and values.
func ParseStructTag(tag string) map[string][]string {
	tag = strings.TrimLeft(tag, "`")
	tag = strings.TrimRight(tag, "`")

	result := make(map[string]string)

	// Iterate through the entire tag string
	for tag != "" {
		// Split the tag into key and the rest
		i := strings.Index(tag, ":")
		if i == -1 {
			break
		}
		key := tag[:i]
		tag = tag[i+1:]

		// Handle quoted value
		if tag[0] == '"' {
			i = 1
			for i < len(tag) {
				if tag[i] == '"' {
					if i+1 < len(tag) && tag[i+1] == '"' {
						i += 2
						continue
					}
					break
				}
				i++
			}
			result[key] = tag[1:i]
			if i+1 < len(tag) && tag[i+1] == ' ' {
				tag = tag[i+2:]
			} else {
				tag = tag[i+1:]
			}
		} else {
			i = strings.Index(tag, " ")
			if i == -1 {
				result[key] = tag
				tag = ""
			} else {
				result[key] = tag[:i]
				tag = tag[i+1:]
			}
		}
	}

	// Convert the map to a slice
	resultSlice := make(map[string][]string, len(result))
	for k, v := range result {
		splitByComma := strings.Split(v, ",")

		resultSlice[k] = splitByComma
	}

	return resultSlice
}
