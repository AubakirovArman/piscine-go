package piscine

func Map(f func(int) bool, a []int) []bool {
	res := make([]bool, len(a))
	for idx, val := range a {
		res[idx] = f(val)
	}
	return res
}
