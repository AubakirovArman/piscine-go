package PrintNbr

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) (int, error) {
	res := "0123456789"
	dlina := 0

	if n == 0 {
		z01.PrintRune(rune(res[0]))
		return 1, nil
	} else {
		if n < 0 {
			if n == -9223372036854775808 {
				res = "-9223372036854775808"
				dlina = 20
				for b := 0; b < 20; b++ {
					z01.PrintRune(rune(res[b]))
				}
				return dlina, nil
			} else {
				z01.PrintRune(rune("-"[0]))
				n = n * -1
			}
			dlina++
		}
		for p := n; p > 10; p = p / 10 {
			dlina++
		}
		dlina++
		if n/1000000000000000 > 0 {
			n1 := n / 1000000000000000
			n = n % 1000000000000000
			t := 0
			for i := 1; i < n1; i = i * 10 {
				t = i
			}
			for i1 := t; i1 >= 1; i1 = i1 / 10 {
				f := n1 / i1
				z01.PrintRune(rune(res[f]))
				n1 = n1 % i1
			}
			t1 := 0
			for i := 1; i < n; i = i * 10 {
				t1 = i
			}
			for i1 := t1; i1 >= 1; i1 = i1 / 10 {
				f := n / i1
				z01.PrintRune(rune(res[f]))
				n = n % i1
			}
		} else {
			t := 0
			for i := 1; i < n; i = i * 10 {
				t = i
			}
			for i1 := t; i1 >= 1; i1 = i1 / 10 {
				f := n / i1
				n = n % i1
				z01.PrintRune(rune(res[f]))
			}
		}
	}
	return dlina, nil
}
