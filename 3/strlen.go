package piscine

func StrLen(s string) int {
	count := 0
	for i := range s {
		if s[i] > 1 {
			count++
		}
	}
	return count
}
