package piscine

func IsNumeric(s string) bool {
	r := "0123456789"
	for _, i := range s {
		count := 0
		for _, v := range r {
			if i == v {
				count++
			}
		}
		if count == 0 {
			return false
		}
	}
	return true
}
