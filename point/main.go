package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	li := "0123456789"
	f := points.x / 10
	t := points.x % 10
	res := "x = " + string(li[f]) + string(li[t])
	f = points.y / 10
	t = points.y % 10
	res = res + ", y = " + string(li[f]) + string(li[t]) + "\n"
	for _, i := range res {
		z01.PrintRune(i)
	}
}
