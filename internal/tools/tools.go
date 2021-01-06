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

func MakeRuneSlice(r rune, length int) (o []rune) {
	for i := 0; i < length; i += 1 {
		o = append(o, r)
	}
	return
}