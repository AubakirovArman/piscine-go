package piscine

func IsSorted(f1 func(a, b int) int, a []int) bool {
	if len(a) >= 2 {
		if a[0] > a[1] {
			for i := 0; i < len(a)-2; i++ {
				if (f1(a[i], a[i+1])) < 0 {
					return false
				}
			}
		} else {
			for i := 0; i < len(a)-2; i++ {
				if (f1(a[i], a[i+1])) > 0 {
					return false
				}
			}
		}
	} else {
		return true
	}
	return true
}

func f(a, b int) int {
	if b < a {
		return 1
	} else if b == a {
		return 0
	} else {
		return -1
	}
}
