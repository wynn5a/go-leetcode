package solution22

func generateParenthesis(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, current string, open int, close int, max int) {
	// If the length of the current string is equal to the maximum length of the string
	// that can be generated, append the current string to the result slice and return
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	// If the number of opening parentheses is less than the maximum number of parentheses,
	// add an opening parenthesis to the current string and call the backtrack function recursively
	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}

	// If the number of closing parentheses is less than the number of opening parentheses,
	// add a closing parenthesis to the current string and call the backtrack function recursively
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}
