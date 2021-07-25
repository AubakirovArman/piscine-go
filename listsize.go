package piscine

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	} else {
		val := l.Head
		count := 1
		for val.Next != nil {
			val = val.Next
			count++
		}
		return count
	}
}
