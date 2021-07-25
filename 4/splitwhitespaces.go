package piscine

func SplitWhiteSpaces(s string) []string {
	st := 0
	en := 0
	li := []string{}
	for idx, val := range s {
		if val == '\n' || val == '	' || val == ' ' || idx == len(s)-1 {
			st = en
			en = idx
			if st != 0 {
				st++
			}
			if idx == len(s)-1 {
				en++
			}
			if st < en {
				li = append(li, string(s[st:en]))
			}
		}
	}
	return li
}
