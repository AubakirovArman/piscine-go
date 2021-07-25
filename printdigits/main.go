package main

import "github.com/01-edu/z01"

func main() {
	sample1 := "0123456789\n"

	for i := 0; i < len(sample1); i++ {
		z01.PrintRune(rune(sample1[i]))
	}
}
