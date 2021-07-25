package piscine

func RecursiveFactorial(nb int) int {
	if nb > -1 && nb < 21 {
		if nb == 0 {
			return 1
		}
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 0
	}
}
