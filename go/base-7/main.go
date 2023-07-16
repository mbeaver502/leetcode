package main

import (
	"log"
	"strconv"
	"strings"
)

type example struct {
	n int
}

func main() {
	examples := []example{
		{100},
		{-10},
		{-7},
		{1024},
		{-1024},
	}

	for _, ex := range examples {
		log.Printf("input:  %d\n", ex.n)
		log.Printf("output: %s\n", convertToBase7(ex.n))
	}
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	var sb strings.Builder

	n := num
	if n < 0 {
		n *= (-1)
	}

	for n > 0 {
		d := n % 7
		sb.WriteString(strconv.FormatInt(int64(d), 7))
		n /= 7
	}

	if num < 0 {
		sb.WriteString("-")
	}

	s := sb.String()
	sb.Reset()

	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}

	return sb.String()
}
