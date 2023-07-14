package main

import (
	"log"
	"sort"
	"strings"
)

type example struct {
	values []string
}

func main() {
	examples := []example{
		{[]string{"flower", "flow", "flight"}},
		{[]string{"dog", "racecar", "car"}},
		{[]string{"hello", "hell", "he", ""}},
		{[]string{"hello", "hell", "helioscope", "helicopter"}},
		{[]string{"hello", "hell", "helioscope", "gyrocopter"}},
		{[]string{"cir", "car"}},
	}

	for _, example := range examples {
		log.Printf("input = %v", example.values)
		log.Printf("output = %s", longestCommonPrefix(example.values))
	}
}

func longestCommonPrefix(strs []string) string {
	// We're guaranteed  1 <= strs.length <= 200
	// and 0 <= strs[i].length <= 200 (i.e., could be empty string).

	// Probably an unnecessary allocation,
	// but it makes me feel better
	// for when we do a sort on this later.
	sorted := make([]string, len(strs))
	copy(sorted, strs)

	// We'll assume that the longest common prefix == shortest string,
	// and we'll shorten that assumed prefix until we find the
	// actual longest common prefix.
	shortest := findShortestString(sorted)
	if shortest == "" {
		return shortest
	}

	shortestLen := len(shortest)
	offset := 0
	done := false

	for shortestLen-offset > 0 && !done {
		for _, str := range sorted {
			// If we find a mismatch on the substring, that means we haven't found a common
			// prefix yet, so we need to try the next shortest substring.
			if !strings.EqualFold(str[:shortestLen-offset], shortest[:shortestLen-offset]) {
				offset++
				done = false
				break
			}
			done = true
		}
	}

	return shortest[:shortestLen-offset]
}

func findShortestString(strs []string) string {
	// This could be a pretty expensive operation,
	// depending on the internal Go implementation
	// of sort. But it does make it convenient to
	// find a starting point.
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}
	return shortest
}
