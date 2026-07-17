package main

/*
-1 HAS
0 Gray
1 GOOD

*/

func GetIncorrectState(c rune, a string) int {
	for _, aChar := range a {
		if aChar == c {
			return -1
		}
	}

	return 0
}

func ReplaceAtIndex(s string, r rune, i int) string {
	output := []rune(s)
	output[i] = r
	return string(output)
}

func GetStatesOfLetters(a string, g string) [5]int {
	// remember g and output share the same positioning
	// in terms of rune/char to square
	var output [5]int

	workingCopy := []rune(a)
	for pos, char := range g {
		if char == rune(a[pos]) {
			workingCopy[pos] = '_'
			output[pos] = 1

		}
	}

	for pos, char := range g {
		if output[pos] != 1 {
			output[pos] = GetIncorrectState(char, string(workingCopy))
		}
	}

	return output
}
