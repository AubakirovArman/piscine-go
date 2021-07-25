package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head, l.Tail = n, l.Head
	} else {
		ne := &NodeL{Data: data}
		ne.Next, l.Head = l.Head, ne
	}
}
