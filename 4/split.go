package piscine

func Split(s, sep string) []string {
	st := ""
	li := []string{}
	for i := 0; i < len(s)-(len(sep)-1); i++ {
		if string(s[i:i+len(sep)]) == sep {
			li = append(li, st)
			st = ""
			i = i + (len(sep) - 1)
		} else {
			st = st + string(s[i])
		}
	}
	st = st + string(s[len(s)-(len(sep)-1):])
	li = append(li, st)
	return li
}
