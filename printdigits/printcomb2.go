package main

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	res := "0123456789"
	for i := 0; i < 100; i++ {
		f := i / 10
		s := i % 10
		for i1 := i + 1; i1 < 100; i1++ {
			f1 := i1 / 10
			s1 := i1 % 10
			if i > 0 || i1 > 1 {
				z01.PrintRune(rune((",")[0]))
				z01.PrintRune(rune((" ")[0]))
			}
			z01.PrintRune(rune(res[f]))
			z01.PrintRune(rune(res[s]))
			z01.PrintRune(rune((" ")[0]))
			z01.PrintRune(rune(res[f1]))
			z01.PrintRune(rune(res[s1]))
		}
	}
	z01.PrintRune(rune(("\n")[0]))
}
