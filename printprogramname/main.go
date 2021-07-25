package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, val := range os.Args[0] {
		if i > 1 {
			z01.PrintRune(val)
		}
	}
	z01.PrintRune('\n')
}
