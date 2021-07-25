package piscine

func Capitalize(s string) string {
	r := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r1 := "abcdefghijklmnopqrstuvwxyz"
	dig := "0123456789"
	statUpper := 0
	nof := 0
	text := []rune(s)
	for i, s1 := range s {
		for i1, rb := range r {
			if s1 == rb {
				text[i] = rune(r1[i1])
			}
		}
	}
	for i := 0; i < len(text); i++ {
		if statUpper == 0 {
			nof = 1
			for up := 0; up < len(r1); up++ {
				if text[i] == rune(r1[up]) {
					text[i] = rune(r[up])
					statUpper = 1
					nof = 0
					break
				}
			}
			if nof == 1 {
				nof = 2
				for up := 0; up < len(dig); up++ {
					if text[i] == rune(dig[up]) {
						statUpper = 1
						nof = 0
						break
					}
				}
			}
			if nof == 2 {
				statUpper = 0
			}
		} else if statUpper == 1 {
			nof = 1
			for up := 0; up < len(r1); up++ {
				if text[i] == rune(r1[up]) {
					statUpper = 1
					nof = 0
					break
				}
			}
			if nof == 1 {
				nof = 2
				for up := 0; up < len(dig); up++ {
					if text[i] == rune(dig[up]) {
						statUpper = 1
						nof = 0
						break
					}
				}
			}
			if nof == 2 {
				statUpper = 0
			}
		}
	}
	return string(text)
}
