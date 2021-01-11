package tools

// FindLongestStringLen returns the length of the longest string in a slice of strings
func FindLongestStringLen(x []string) int {
	var longest int
	for _, y := range x {
		if z := len(y); z > longest {
			longest = z
		}
	}
	return longest
}

// MakeRuneSlice returns a slice of runes of len length, with each run being rune r
func MakeRuneSlice(r rune, length int) (o []rune) {
	for i := 0; i < length; i += 1 {
		o = append(o, r)
	}
	return
}

// RightPadString pads a string to a specific length with a specific character
func RightPadString(in string, newLen int, pad rune) string {
	for i := 0; i < newLen-len(in); i += 1 {
		in += string(pad)
	}
	return in
}

// see https://github.com/CSCoursework/cs-battleships#numbering-the-board-with-letters for info about these two
// functions

func GetAlphabetChar(i int) string {
	return string(rune(int('a') + i))
}

func GetCharNumber(i string) int {
	return int([]rune(i)[0]) - int('a')
}
