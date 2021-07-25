package piscine

func Swap(a *int, b *int) {
	b1 := *a
	*a = *b
	*b = b1
}
