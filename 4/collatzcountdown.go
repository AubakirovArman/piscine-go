package piscine

func CollatzCountdown(start int) int {
	var1 := 0
	for start != 1 {
		if start <= 0 {
			return -1
		}
		if start%2 == 0 {
			start = start / 2
			var1++
		} else {
			start = start*3 + 1
			var1++
		}
	}
	return var1
}
