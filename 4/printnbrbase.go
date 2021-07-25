package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if nbr == -9223372036854775808 {
		so := "-9223372036854775808"
		for _, s := range so {
			z01.PrintRune(rune(s))
		}
		return
	}
	var num [40]rune
	b := []rune(base)
	len := len(b)
	minus := false
	if len < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		nbr = nbr * -1
		minus = true
	}
	for i, s := range b {
		for i1 := i + 1; i1 < len; i1++ {
			if s == b[i1] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
		if s == '-' || s == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	index := 0
	for nbr > 0 {
		num[index] = b[nbr%len]
		nbr = nbr / len
		index++
	}
	if minus {
		z01.PrintRune('-')
	}
	for i := index - 1; i > -1; i-- {
		z01.PrintRune(num[i])
	}
}
