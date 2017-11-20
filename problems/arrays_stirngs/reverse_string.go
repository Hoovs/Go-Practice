// Reverses a string including white space.
package arrays_stirngs

func reverseString(inStr string) string {
	strLen := len(inStr) - 1

	outRune := []rune(inStr)
	for i, r := range outRune {
		tmp := rune(inStr[i])
		outRune[i] = r
		outRune[strLen-i] = tmp
	}

	return string(outRune)
}
