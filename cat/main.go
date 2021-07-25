package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		for idx, args := range os.Args {
			if idx > 0 {
				content, err := ioutil.ReadFile(args)
				if err != nil {
					er := "ERROR: " + err.Error() + "\n"
					for _, va := range er {
						z01.PrintRune(va)
					}
					os.Exit(1)
				} else {
					for _, va := range content {
						z01.PrintRune(rune(va))
					}
				}
			}
		}
	}
}
