package solution824

import "strings"

func toGoatLatin(sentence string) string {
	var res []string
	for i, word := range strings.Split(sentence, " ") {
		res = append(res, toGoatLatinWord(i+1, word))
	}
	return strings.Join(res, " ")
}

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

// toGoatLatinWord handle single word
func toGoatLatinWord(i int, word string) string {
	if word == "" {
		return ""
	}

	if _, ok := vowels[rune(word[0])]; ok {
		return word + "ma" + strings.Repeat("a", i)
	}

	if _, ok := vowels[rune(word[0])]; !ok {
		return word[1:] + string(word[0]) + "ma" + strings.Repeat("a", i)
	}

	return ""
}
