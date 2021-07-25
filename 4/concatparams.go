package piscine

func ConcatParams(args []string) string {
	s := ""
	for i, val := range args {
		if i == 0 {
			s = s + val
		} else {
			s = s + "\n" + val
		}
	}
	return s
}
