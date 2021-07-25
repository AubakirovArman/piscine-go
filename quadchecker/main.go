package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	info, err := ioutil.ReadAll(os.Stdin)
	if err == nil {
		x := 0
		y := 0
		che := 0
		for _, val := range info {
			if che == 0 {
				x++
			}
			if val == '\n' {
				if che == 0 {
					che = 1
				}
				y++
			}
		}
		x--
		otvet := []string{}
		info1 := string(info)
		if info1 == QuadMutual(x, y, 'o', '-', 'o', '|', 'o', 'o') {
			otvet = append(otvet, "[quadA] ["+inttostr(x)+"] ["+inttostr(y)+"]")
		}
		if info1 == QuadMutual(x, y, '/', '*', '\\', '*', '\\', '/') {
			otvet = append(otvet, "[quadB] ["+inttostr(x)+"] ["+inttostr(y)+"]")
		}
		if info1 == QuadMutual(x, y, 'A', 'B', 'A', 'B', 'C', 'C') {
			otvet = append(otvet, "[quadC] ["+inttostr(x)+"] ["+inttostr(y)+"]")
		}
		if info1 == QuadMutual(x, y, 'A', 'B', 'C', 'B', 'A', 'C') {
			otvet = append(otvet, "[quadD] ["+inttostr(x)+"] ["+inttostr(y)+"]")
		}
		if info1 == QuadMutual(x, y, 'A', 'B', 'C', 'B', 'C', 'A') {
			otvet = append(otvet, "[quadE] ["+inttostr(x)+"] ["+inttostr(y)+"]")
		}
		if len(otvet) == 0 {
			fmt.Println("Not a Raid function")
		} else {
			fmt.Println(splitter(otvet))
		}

	} else {
		fmt.Println("Not a Raid function")
	}
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

func splitter(a []string) string {
	if len(a) == 1 {
		return a[0]
	}
	return a[0] + " || " + splitter(a[1:])
}

func QuadMutual(x, y int, leftTop, top, rightTop, sides, leftBot, rightBot rune) string {
	res := ""
	for i := 0; i < y; i++ {
		for i1 := 0; i1 < x; i1++ {
			if (i1 == 0 || i1 == x-1) && (i == 0 || i == y-1) {
				if i1 == 0 && i == 0 {
					res += string(leftTop)
				} else if i1 == 0 && i == y-1 {
					res += string(leftBot)
				} else if i1 == x-1 && i == 0 {
					res += string(rightTop)
				} else if i1 == x-1 && i == y-1 {
					res += string(rightBot)
				}
			} else {
				if i == 0 || i == (y-1) {
					res += string(top)
				} else if i1 == 0 || i1 == x-1 {
					res += string(sides)
				} else {
					res += string(' ')
				}
			}
		}
		res = res + string('\n')
	}
	return res
}
