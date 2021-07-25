package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	insert := false
	insert_val := ""
	target_val := os.Args[len(os.Args)-1]
	order := false
	if len(os.Args) <= 2 {
		if target_val == "-h" || target_val == "--help" || len(os.Args) == 1 {
			helpShow()
		} else {
			fmt.Println(target_val)
		}
	} else {
		for i := 1; i < len(os.Args); i++ {
			if len(os.Args[i]) > 7 {
				if (os.Args[i])[:8] == "--insert" {
					insert = true
					insert_val = (os.Args[i])[9:]
				} else if (os.Args[i])[:2] == "-i" {
					insert = true
					insert_val = (os.Args[i])[3:]
				} else {
					target_val = os.Args[i]
				}
			} else if len(os.Args[i]) == 7 {
				if (os.Args[i]) == "--order" {
					order = true
				} else if (os.Args[i])[:2] == "-i" {
					insert = true
					insert_val = (os.Args[i])[3:]
				} else {
					target_val = os.Args[i]
				}
			} else if len(os.Args[i]) == 2 {
				if (os.Args[i])[:2] == "-o" {
					order = true
				} else {
					target_val = os.Args[i]
				}
			} else if len(os.Args[i]) != 0 {
				if (os.Args[i])[:2] == "-i" {
					insert = true
					insert_val = (os.Args[i])[3:]
				} else {
					target_val = os.Args[i]
				}
			}
		}
		if insert {
			target_val = target_val + insert_val
		}
		t := []rune{}
		for _, s := range target_val {
			t = append(t, s)
		}
		if order {
			for i := 0; i < len(t)-1; i++ {
				for i1 := i + 1; i1 < len(t); i1++ {
					if t[i] > t[i1] {
						per := t[i]
						per1 := t[i1]
						t[i] = rune(per1)
						t[i1] = rune(per)
					}
				}
			}
		}
		for _, s := range t {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}

func helpShow() {
	help := "--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument."
	fmt.Println(help)
}
