package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		val := l.Head
		for val.Next != nil {
			val = val.Next
		}
		return val.Data
	}
}
