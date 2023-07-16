// https://leetcode.com/problems/excel-sheet-column-number/
// Given a string columnTitle that represents the column title
// as appears in an Excel sheet, return its corresponding column number.
package main

import (
	"log"
	"math"
)

type example struct {
	columnTitle string
}

func main() {
	examples := []example{
		{"A"},
		{"AB"},
		{"ZY"},
	}

	for _, ex := range examples {
		log.Printf("input:  %s\n", ex.columnTitle)
		log.Printf("output: %d\n", titleToNumber(ex.columnTitle))
	}
}

// For example, "ZY" = 701 = 26*(26^1) + 25*(26^0)
func titleToNumber(columnTitle string) int {
	out := 0
	pos := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		out += int(columnTitle[i]-'A'+1) * int(math.Pow(26.0, float64(pos)))
		pos++
	}
	return out
}
