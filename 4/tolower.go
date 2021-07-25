package piscine

func ToLower(s string) string {
	s1 := []rune(s)
	r := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r1 := "abcdefghijklmnopqrstuvwxyz"
	for i := range s1 {
		for i1 := range r {
			if s1[i] == rune(r[i1]) {
				s1[i] = rune(r1[i1])
			}
		}
	}
	return string(s1)
}
