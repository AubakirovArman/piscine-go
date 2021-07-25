package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) (int, error) {
	for i := range s {
		z01.PrintRune(rune(s[i]))
	}
	return len(s), nil
}
