package solution17

func letterCombinations(digits string) (rst []string) {
	if len(digits) == 0 {
		return rst
	}

	phoneMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var backtrack func(string, int, string)
	backtrack = func(digits string, index int, combination string) {
		if index == len(digits) {
			rst = append(rst, combination)
		} else {
			digit := string(digits[index])
			letters := phoneMap[digit]
			lettersCount := len(letters)
			for i := 0; i < lettersCount; i++ {
				backtrack(digits, index+1, combination+string(letters[i]))
			}
		}
	}

	backtrack(digits, 0, "")

	return rst
}
