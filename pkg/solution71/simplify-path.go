package solution71

import (
	"strings"
)

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	// create a stack to store the directories encountered
	var stack []string

	for _, dir := range dirs {
		switch dir {
		case "", ".":
			// skip when the directory is empty or "."
		case "..":
			if len(stack) > 0 {
				// pop the last directory from the stack if it is not empty when ".." is encountered
				stack = stack[:len(stack)-1]
			}
		default:
			// push the directory to the stack if it is not empty when a valid directory name is encountered
			stack = append(stack, dir) // push the directory
		}
	}

	return "/" + strings.Join(stack, "/")
}
