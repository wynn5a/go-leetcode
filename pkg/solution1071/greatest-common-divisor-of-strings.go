package solution1071

func gcdOfStrings(str1 string, str2 string) string {
	// if str1+str2 != str2+str1, then there is no solution, because the common prefix of them is different
	if str1+str2 != str2+str1 {
		return ""
	}
	// gcd of two strings is the common prefix of them
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
