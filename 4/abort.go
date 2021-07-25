package piscine

func Abort(a, b, c, d, e int) int {
	li := []int{a, b, c, d, e}
	for range li {
		for i1 := 0; i1 < len(li)-1; i1++ {
			if li[i1] > li[i1+1] {
				b := li[i1]
				li[i1] = li[i1+1]
				li[i1+1] = b
			}
		}
	}
	return li[2]
}
