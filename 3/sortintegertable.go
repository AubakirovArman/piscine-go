package SortIntegerTable

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for i1 := 0; i1 < len(table)-1; i1++ {
			if int(table[i1]) > int(table[i1+1]) {
				p := table[i1]
				table[i1] = table[i1+1]
				table[i1+1] = p
			}
		}
	}
}
