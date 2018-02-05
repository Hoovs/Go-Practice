package arrays_stirngs

import (
	"sort"
	"strings"
)


func isPermutationSort(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1Split := strings.Split(str1, "")
	sort.Strings(str1Split)
	str2Split := strings.Split(str2, "")
	sort.Strings(str2Split)

	return strings.Join(str1Split, "") == strings.Join(str2Split, "")
}

func isPermutationBit(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	var seenBefore [127]uint16

	for _, r := range str1 {
		seenBefore[r]++
	}

	for _, r := range str2 {
		seenBefore[r]--
		if seenBefore[r] < 0 {
			return false
		}
	}

	return true

}