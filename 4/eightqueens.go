package piscine

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	n := 8
	arr := [8]int{}

	Queens(0, n, arr)
}

func IsSolution(k int, position int, arr [8]int) bool {
	for i := 0; i < k; i++ {
		test := arr[i]

		if test == position || test == position-(k-i) || test == position+(k-i) {
			return false
		}
	}

	return true
}

func printQuen(arr [8]int, n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(arr[i] + '1'))
	}
	z01.PrintRune('\n')
}

func Queens(k int, n int, arr [8]int) {
	if k == n {
		printQuen(arr, n)
	} else {
		for i := 0; i < n; i++ {
			if IsSolution(k, i, arr) {
				arr[k] = i
				if i >= 1 {
					break
				}
				Queens(k+1, n, arr)

			}
		}
	}
}
