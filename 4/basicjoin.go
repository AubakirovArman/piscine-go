package piscine

func BasicJoin(elems []string) string {
	s := ""
	for _, val := range elems {
		s = s + val
	}
	return s
}
