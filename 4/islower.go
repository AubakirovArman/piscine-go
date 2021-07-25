package piscine

func IsLower(s string) bool {
	r := "abcdefghijklmnopqrstuvwxyz"
	count := 0
	for _, i := range s {
		for _, v := range r {
			if i == v {
				count++
			}
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
