package main

import "github.com/01-edu/z01"

func main() {
	sample := "abcdefghijklmnopqrstuvwxyz"
	sample1 := Reverse(sample) + "\n"
	for i := 0; i < len(sample1); i++ {
		z01.PrintRune(rune(sample1[i]))
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
