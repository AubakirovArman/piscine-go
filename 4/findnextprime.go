package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		if nb <= 1 {
			return 2
		}
		for d := 2; d*d <= nb; d++ {
			if nb%d == 0 {
				return FindNextPrime(nb + 1)
			}
		}
		return nb
	} else {
		return 2
	}
}
