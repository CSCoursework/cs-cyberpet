package tools

func FindLongestString(x []string) int {
	var longest int
	for _, y := range x {
		if z := len(y); z > longest {
			longest = z
		}
	}
	return longest
}
