package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		l1 := l.Head
		for l1.Next != nil {
			l1 = l1.Next
		}
		l1.Next = n
		l.Tail = n
	}
}
