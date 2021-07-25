package piscine

func AppendRange(min, max int) []int {
	li := []int{}
	b := 0
	for i := min; i < max; i++ {
		li = append(li, i)
		b++
	}
	if b == 0 {
		return []int(nil)
	} else {
		return li
	}
}
