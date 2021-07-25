package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, val := range os.Args {
		if i > 0 {
			for _, val2 := range val {
				z01.PrintRune(val2)
			}
			z01.PrintRune('\n')
		}
	}
}
