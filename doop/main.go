package main

import (
	"os"
)

func pow10(n int, a int) int {
	pp := a * 1
	for i := 1; i < n; i++ {
		pp *= 10
	}
	return pp
}

func strtoint(a string) int {
	array := []rune{}
	checker := false
	if a[0] == 45 {
		checker = true
	}
	if checker {
		for i := 1; i < len(a); i++ {
			array = append(array, rune(a[i]))
		}
	} else {
		for i := 0; i < len(a); i++ {
			array = append(array, rune(a[i]))
		}
	}
	arrayByte1 := []byte{}
	for j := 0; j < len(array); j++ {
		arrayByte1 = append(arrayByte1, byte(rune(array[j])-48))
	}
	result1 := 0
	for k := 0; k < len(arrayByte1); k++ {
		aByte := arrayByte1[k]
		aInt := int(aByte)
		result1 += pow10(len(arrayByte1)-k, aInt)
	}
	if checker {
		return (result1 * (-1))
	}
	return result1
}

func inttostr(a int) string {
	array := []rune{}
	checker := false
	if a < 0 {
		checker = true
	}
	if checker {
		a *= -1
	}
	for a != 0 {
		array = append(array, rune(a%10))
		a /= 10
	}
	resarr := []string{}
	for i := len(array) - 1; i >= 0; i-- {
		resarr = append(resarr, string((array[i])+48))
	}
	if checker {
		result := "-"
		for i := 0; i < len(resarr); i++ {
			result += resarr[i]
		}
		return result
	} else {
		result := ""
		for i := 0; i < len(resarr); i++ {
			result += resarr[i]
		}
		return result
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 3 {
		a := args[0]
		operator := args[1]
		b := args[2]
		if a[0] == 45 {
		} else {
			if a <= "-9223372036854775809" {
				return
			}
		}
		if b[0] == 45 {
		} else {
			if b <= "-9223372036854775809" {
				return
			}
		}
		if a[0] == 45 {
			for i := 1; i < len(a); i++ {
				if a[i] >= 48 && a[i] <= 57 {
				} else {
					return
				}
			}
		} else {
			for i := 0; i < len(a); i++ {
				if a[i] >= 48 && a[i] <= 57 {
				} else {
					return
				}
			}
		}
		if b[0] == 45 {
			for i := 1; i < len(b); i++ {
				if b[i] >= 48 && b[i] <= 57 {
				} else {
					return
				}
			}
		} else {
			for i := 0; i < len(b); i++ {
				if b[i] >= 48 && b[i] <= 57 {
				} else {
					return
				}
			}
		}
		result1 := strtoint(a)
		result2 := strtoint(b)
		switch operator {
		case "+":
			result := result1 + result2
			str := inttostr(result)
			if result > result1 {
				os.Stderr.WriteString(str + "\n")
			} else {
			}
		case "-":
			result := result1 - result2
			str := inttostr(result)
			if result < result1 {
				os.Stderr.WriteString(str + "\n")
			} else {
			}
		case "*":
			result := result1 * result2
			str := inttostr(result)
			if result/result1 == result2 {
				os.Stderr.WriteString(str + "\n")
			} else {
			}
		case "/":
			if result2 == 0 {
				os.Stderr.WriteString("No division by 0\n")
			} else {
				result := result1 / result2
				if result == 0 {
					os.Stderr.WriteString("0\n")
				} else {
					str := inttostr(result)
					os.Stderr.WriteString(str + "\n")
				}
			}
		case "%":
			if result2 == 0 {
				os.Stderr.WriteString("No modulo by 0\n")
			} else {
				result := result1 % result2
				str := inttostr(result)
				os.Stderr.WriteString(str + "\n")
			}
		}

	}
}
