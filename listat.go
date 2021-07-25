package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	val := l
	index := 0
	for index < pos {
		if val.Next == nil {
			return nil
		}
		val = val.Next
		index++
	}
	return val
}
