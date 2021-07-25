package piscine

func AlphaCount(s string) int {
	s1 := []rune(s)
	ans := 0
	for _, l := range s1 {
		if (64 < l && l < 91) || (96 < l && l < 123) {
			ans++
		}
	}
	return ans
}
