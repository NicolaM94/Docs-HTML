package utilities

// Replaces white spaces to zeros
func replaceSpaces(str string) string {
	out := ""
	for x := range str {
		if str[x] == ' ' {
			out = out + "0"
			continue
		}
		out = out + string(str[x])
	}
	return out
}

// Replaces - sign in strings with _
func replaceMinus(str string) string {
	out := ""
	for x := range str {
		if str[x] == '-' {
			out = out + "_"
			continue
		}
		out = out + string(str[x])
	}
	return out
}

// Wrapper for the above functions.
func NormalizeString(str string) string {
	out := replaceSpaces(str)
	out = replaceMinus(out)
	return out
}
