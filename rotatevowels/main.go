package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	rep := []rune{}
	ans := ""
	dlina := len(os.Args)
	len := 0
	first := true
	if dlina == 1 {
		z01.PrintRune('\n')
		return
	}
	for _, arg := range args {
		for _, j := range arg {
			if pro(j) {
				rep = append(rep, j)
				len++
			}
		}
		if first {
			ans = arg
			first = false
			continue
		}
		ans = ans + " " + arg
	}
	cur := 0
	for _, c := range ans {
		if pro(c) {
			z01.PrintRune(rep[len-cur-1])
			cur++
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune(' ')
	z01.PrintRune('\n')
}

func pro(x rune) bool {
	st := "aAeEoOuUiI"
	for _, s := range st {
		if rune(s) == x {
			return true
		}
	}
	return false
}
