package piscine

func Index(s string, toFind string) int {
	if s == toFind {
		return 0
	} else {
		if len(toFind) > len(s) {
			return -1
		} else {
			for i := 0; i <= len(s)-len(toFind); i++ {
				if s[i:(len(toFind)+i)] == toFind {
					return i
				}
			}
			return -1
		}
	}
}
