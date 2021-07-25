package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	numberFrom := convId(nbr, baseFrom)
	numberTo10 := convTo10(numberFrom, baseFrom)
	return conve10TobaseTo(baseTo, numberTo10)
}

func convId(nbr, baseFrom string) []int {
	massiv := []int{}
	for _, s := range nbr {
		for idx, val := range baseFrom {
			if s == val {
				massiv = append(massiv, idx)
			}
		}
	}
	return massiv
}

func convTo10(numberFrom []int, baseFrom string) int {
	num := 0
	for idx, val := range numberFrom {
		b := 1
		for i := 1; i < (len(numberFrom) - idx); i++ {
			b = b * len(baseFrom)
		}
		num = num + (val * b)
	}
	return num
}

func conve10TobaseTo(baseTo string, num int) string {
	res := ""
	for num > 0 {
		idx := num % len(baseTo)
		res = string(baseTo[idx]) + res
		num = num / len(baseTo)
	}
	return res
}
