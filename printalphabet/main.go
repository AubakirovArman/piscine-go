package main

import "github.com/01-edu/z01"

func main() {
	sample := "abcdefghijklmnopqrstuvwxyz\n"
	for i := 0; i < len(sample); i++ {
		z01.PrintRune(rune(sample[i]))
	}
}
