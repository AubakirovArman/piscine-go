package piscine

func AtoiBase(s string, base string) int {
	b := []rune(base)
	len1 := len(b)
	if len1 < 2 {
		return 0
	}
	for i, s := range b {
		for i1 := i + 1; i1 < len1; i1++ {
			if s == b[i1] {
				return 0
			}
		}
		if s == '-' || s == '+' {
			return 0
		}
	}
	res := 1
	ot := 0
	for i := (len(s) - 1); i >= 0; i-- {
		for idx, val := range b {
			if val == rune(s[i]) {
				ot = ot + idx*res
				res = res * len1
			}
		}
	}
	return ot
}
