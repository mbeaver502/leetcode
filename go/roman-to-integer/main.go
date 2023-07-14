// https://leetcode.com/problems/roman-to-integer/description/
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Given a roman numeral, convert it to an integer.
package main

import "log"

type example struct {
	value string
}

func main() {
	examples := []example{
		{"III"},
		{"LVIII"},
		{"MCMXCIV"},
		{"LXXXIX"},
		{"XCV"},
		{"L"},
	}

	for _, example := range examples {
		log.Println(romanToInt(example.value))
	}
}

func romanToInt(s string) int {
	out := 0
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		switch s[i] {
		case 'M':
			out += 1000
		case 'D':
			out += 500
		case 'C':
			if i < sLen-1 {
				switch s[i+1] {
				case 'M':
					out += 900
					i++
				case 'D':
					out += 400
					i++
				default:
					out += 100
				}
			} else {
				out += 100
			}
		case 'L':
			out += 50
		case 'X':
			if i < sLen-1 {
				switch s[i+1] {
				case 'C':
					out += 90
					i++
				case 'L':
					out += 40
					i++
				default:
					out += 10
				}
			} else {
				out += 10
			}
		case 'V':
			out += 5
		case 'I':
			if i < sLen-1 {
				switch s[i+1] {
				case 'X':
					out += 9
					i++
				case 'V':
					out += 4
					i++
				default:
					out += 1
				}
			} else {
				out += 1
			}
		}
	}

	return out
}
