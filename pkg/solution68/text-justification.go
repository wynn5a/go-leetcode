package solution68

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var line []string
	var lineLength int
	for _, word := range words {
		// If the line is too long, justify it and start a new line
		if lineLength+len(word)+len(line) > maxWidth {
			result = append(result, justify(line, lineLength, maxWidth))
			line = nil
			lineLength = 0
		}
		line = append(line, word)
		lineLength += len(word)
	}
	result = append(result, lastLine(line, maxWidth))
	return result
}

// The last line is left-justified and does not contain extra spaces between words.
func lastLine(line []string, width int) string {
	var result string
	for _, word := range line {
		result += word + " "
	}
	result += spaces(width - len(result))
	return result
}

// Return a string of i spaces
func spaces(i int) string {
	var result string
	for i > 0 {
		result += " "
		i--
	}
	return result
}

// Justify a line of text
func justify(line []string, length int, width int) string {
	// If the line is only one word, left-justify it
	if len(line) == 1 {
		return line[0] + spaces(width-length)
	}
	spaceTotal := width - length
	spacesBetween := spaceTotal / (len(line) - 1)
	spacesLeft := spaceTotal % (len(line) - 1)
	var result string
	// Add the first word
	for i, word := range line {
		result += word
		if i < len(line)-1 {
			result += spaces(spacesBetween)
			if i < spacesLeft {
				result += " "
			}
		}
	}
	return result
}
