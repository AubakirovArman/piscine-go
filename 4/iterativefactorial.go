package piscine

func IterativeFactorial(nb int) int {
	ct := 1
	if nb < 21 && nb >= 0 {
		for i := 1; i <= nb; i++ {
			ct = ct * i
		}
		return ct
	} else {
		return 0
	}
}
