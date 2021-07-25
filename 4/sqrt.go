package piscine

func Sqrt(nb int) int {
	res := 0
	if nb > 0 {
		for i := nb; i > 0; i-- {
			if i*i == nb {
				return i
			}
		}
	} else {
		return 0
	}
	return res
}
