package piscine

func TrimAtoi(s string) int {
	numc := 0
	d := 0
	b := 1
	dig := "0123456789"
	for _, va := range s {
		for i := 0; i < len(dig); i++ {
			if va == rune(dig[i]) {
				d = (d + i) * 10
				numc = 1
			}
		}
		if va == '-' && numc == 0 {
			b = -1
		}
	}
	return (d / 10) * b
}
