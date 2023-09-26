package solution151

import (
	"strings"
)

func reverseWords(s string) string {
	//splits the string
	fields := strings.Fields(s)
	//reverse the array
	for i, j := 0, len(fields)-1; i < j; i, j = i+1, j-1 {
		fields[i], fields[j] = fields[j], fields[i]
	}
	//join to new string
	return strings.Join(fields, " ")
}
