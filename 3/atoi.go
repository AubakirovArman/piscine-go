package Atoi

func Atoi(s string) int {
	res := 0
	sy := "0123456789+-"
	plus := 0
	minus := 0
	for _, b := range s {
		pro := 1
		for i := 0; i < len(sy); i++ {
			if string(b) == string(sy[i]) {
				if string(b) == "+" {
					plus++
					if plus > 1 {
						return 0
					}
				} else if string(b) == "-" {
					minus++
					if minus > 1 {
						return 0
					}
				} else {
					res = res*10 + (int(b) - 48)
				}
				pro = 0
			}
		}
		if pro == 1 {
			return 0
		}

	}
	if minus == 1 {
		if string(s[0]) == "-" {
			res = res * (-1)
		} else {
			return 0
		}
	}
	if plus == 1 {
		if string(s[0]) != "+" {
			return 0
		}
	}
	if minus == 1 && plus == 1 {
		return 0
	}
	return res
}
