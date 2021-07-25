package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n > 0 {
		n1 := []rune{}
		for n > 0 {
			n1 = append(n1, rune(n%10+48))
			n = n / 10
		}
		do := rune('0')
		for i := 0; i < len(n1)-1; i++ {
			for i1 := i + 1; i1 < len(n1); i1++ {
				if n1[i] > n1[i1] {
					do = n1[i]
					n1[i] = n1[i1]
					n1[i1] = do
				}
			}
		}
		for _, s := range n1 {
			z01.PrintRune(s)
		}
	}
}
