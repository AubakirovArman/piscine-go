package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	li := " 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for _, va := range li {
		for i, val := range os.Args {
			if i > 0 {
				for _, val2 := range val {
					if va == rune(val[0]) {
						z01.PrintRune(val2)
					}
				}
				if va == rune(val[0]) {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
