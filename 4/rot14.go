package piscine

func Rot14(s string) string {
	alpa := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sample := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	s11 := ""
	for _, s1 := range s {
		if s1 != ' ' {
			for i, val := range alpa {
				if s1 == val {
					s11 = s11 + string(alpa[i+14])
					break
				}
			}
			for i1, val1 := range sample {
				if s1 == val1 {
					s11 = s11 + string(sample[i1+14])
					break
				}
			}
		} else {
			s11 = s11 + " "
		}
	}
	return s11
}
