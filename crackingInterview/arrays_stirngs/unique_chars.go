// This checks if ASCII string has all unique chars.
package arrays_stirngs

const (
	allowedCodes = 127 // Total allowed ASCII codes
)

func initBitArrayToFalse() [allowedCodes]bool {
	var seenBefore [allowedCodes]bool
	for i := range seenBefore {
		seenBefore[i] = false
	}
	return seenBefore
}

func hasUniqueCharactersBitArray(str string) bool {
	if len(str) > 128 {
		return false
	}

	seenBefore := initBitArrayToFalse()

	for _, r := range str {
		if seenBefore[r] != false {
			return false
		} else {
			seenBefore[r] = true
		}
	}
	return true
}
