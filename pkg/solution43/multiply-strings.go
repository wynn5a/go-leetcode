package solution43

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	// If either input is empty, return an empty string
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}

	// Convert the input strings to byte slices for easier manipulation
	num1Bytes := []byte(num1)
	num2Bytes := []byte(num2)

	// Allocate a slice to hold the result of each multiplication
	result := make([]int, len(num1Bytes)+len(num2Bytes))

	// Iterate through each digit of num1
	for i := len(num1Bytes) - 1; i >= 0; i-- {
		// Convert the current digit to an integer
		digit1 := int(num1Bytes[i] - '0')

		// Initialize the carry to 0
		carry := 0

		// Iterate through each digit of num2
		for j := len(num2Bytes) - 1; j >= 0; j-- {
			// Convert the current digit to an integer
			digit2 := int(num2Bytes[j] - '0')

			// Multiply the two digits and add the carry
			product := digit1*digit2 + carry + result[i+j+1]

			// Update the carry and the result
			carry = product / 10
			result[i+j+1] = product % 10
		}

		// Add any remaining carry to the next digit of num1
		if carry > 0 {
			result[i] += carry
		}
	}

	// Convert the result slice to a string
	var sb strings.Builder
	for _, digit := range result {
		sb.WriteString(strconv.Itoa(digit))
	}

	// Remove any leading zeros from the result
	resultStr := strings.TrimLeft(sb.String(), "0")

	// If the result is empty, return "0"
	if resultStr == "" {
		return "0"
	}

	return resultStr
}
