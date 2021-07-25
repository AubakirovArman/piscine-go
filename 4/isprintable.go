package piscine

func IsPrintable(s string) bool {
	for _, l := range s {
		if 31 < l && l < 128 {
			continue
		} else {
			return false
		}
	}
	return true
}
