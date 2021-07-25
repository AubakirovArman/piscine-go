package piscine

func ToUpper(s string) string {
	s1 := []rune(s)
	r1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r := "abcdefghijklmnopqrstuvwxyz"
	for i := range s1 {
		for i1 := range r {
			if s1[i] == rune(r[i1]) {
				s1[i] = rune(r1[i1])
			}
		}
	}
	return string(s1)
}
