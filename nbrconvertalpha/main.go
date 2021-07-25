package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi11(s string) int {
	sum := 0
	for _, i := range s {
		if '0' <= i && i <= '9' {
			ost := 0
			for k := '1'; k <= i; k++ {
				ost = ost + 1
			}
			sum = sum*10 + ost
		} else {
			sum = -1
			break
		}
	}
	if (sum != -1) && !(1 <= sum && sum <= 26) {
		sum = -1
	}
	return sum
}

func main() {
	arguments := os.Args
	pos := 1
	flagupper := false
	count := 0
	for range arguments {
		count++
	}
	if count >= 2 && arguments[1] == "--upper" {
		pos = 2
		flagupper = true
	}
	for index, k := range arguments {
		if index >= pos {
			num := BasicAtoi11(k)
			if num == -1 {
				z01.PrintRune(' ')
			} else {
				if !flagupper {
					z01.PrintRune(rune('a' + num - 1))
				} else {
					z01.PrintRune(rune('A' + num - 1))
				}
			}
		}
	}
	if len(os.Args) > 1 {
		z01.PrintRune(10)
	}
}
