package BasicAtoi2

func BasicAtoi2(s string) int {
	res := 0
	sy := "0123456789"
	for _, b := range s {
		pro := 1
		for i := 0; i < len(sy); i++ {
			if string(b) == string(sy[i]) {

				res = res*10 + (int(b) - 48)
				pro = 0
			}
		}
		if pro == 1 {
			return 0
		}
	}
	return res
}
