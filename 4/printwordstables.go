package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, s := range a {
		for _, b := range s {
			z01.PrintRune(b)
		}
		z01.PrintRune('\n')
	}
}
