// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// Given two strings needle and haystack,
// return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.
package main

import (
	"log"
	"regexp"
)

type example struct {
	needle   string
	haystack string
}

func main() {
	examples := []example{
		{needle: "sad", haystack: "sadbutsad"},  // 0
		{needle: "leeto", haystack: "leetcode"}, // -1
	}

	for _, ex := range examples {
		log.Printf("input: haystack = %s, needle = %s", ex.haystack, ex.needle)
		log.Printf("output: %d", strStr(ex.haystack, ex.needle))
	}
}

func strStr(haystack string, needle string) int {
	rex := regexp.MustCompile(needle)
	loc := rex.FindIndex([]byte(haystack))
	if loc == nil {
		return -1
	}
	return loc[0]
}
