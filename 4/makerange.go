package piscine

func MakeRange(min, max int) []int {
	if min < max {
		li := make([]int, max-min)
		b := 0
		for i := min; i < max; i++ {
			li[b] = i
			b++
		}
		return li
	} else {
		return []int(nil)
	}
}
