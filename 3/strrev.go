package piscine

func StrRev(s string) (result string) {
	for _, r := range s {
		result = string(r) + result
	}
	return result
}
