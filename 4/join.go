package piscine

func Join(strs []string, sep string) string {
	result := ""
	lenstrs := 0
	for i := range strs {
		lenstrs = i + 1
	}
	for i := range strs {
		if i != lenstrs-1 {
			result = result + strs[i] + sep
		} else {
			result = result + strs[i]
		}
	}
	return result
}
