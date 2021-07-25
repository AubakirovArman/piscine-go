package piscine

func NRune(s string, n int) rune {
	if -1 <= n && n <= len(s) {
		if n < 0 {
			r := []rune(s)
			return r[len(s)+(n)]
		} else {
			r := []rune(s)
			le := n - 1
			if n == 0 {
				return 0
			}
			return r[le]
		}
	} else {
		return 0
	}
}
