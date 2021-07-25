package BasicAtoi

func BasicAtoi(s string) int {
	res := 0
	for _, b := range s {
		if string(b) != "-" && string(b) != "+" {
			res = res*10 + (int(b) - 48)
		}
	}
	return res
}
