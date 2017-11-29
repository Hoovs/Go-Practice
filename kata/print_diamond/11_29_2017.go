package main

import "strings"

//For example: print-diamond 'E' prints
//
//    A
//   B B
//  C   C
// D     D
//E       E
// D     D
//  C   C
//   B B
//    A
//
//For example: print-diamond 'C' prints
//
//  A
// B B
//C   C
// B B
//  A
//0, 1, 2: Can't diamond less than 3
//allocate string array of # letters * 2 - 1
//3: Print C -> Top row has # letters - 1 - i  -> 2 space
//           -> Second row has # letters - 1 - i -> 1 spaces
//           -> Third row has # letters -1 - i -> 0 spaces, middle spaces = (2 * (row - 1)) + 1
//           -> copy second row into next array
//           -> copy top etc....
// O(n) space, O(n) Time assuming using bytes instead of + for strings. Not done here
// for simplicity of logic.
//
// Also there is no need to use and allocate an array to store the strings. Could just pass
// the buffer instead and append them in memory. O(n) still but with a smaller footprint.
const (
	spaceSeparator = " "
	uppercaseChar = 'A'
)

func buildRow(rowNum int, totalLength int) string {
	letter := string(rune(rowNum + uppercaseChar))
	frontStr := strings.Repeat(spaceSeparator, totalLength - rowNum) + letter
	if midStrLength := (2 * (rowNum - 1)) + 1; midStrLength > -1 {
		midStr := strings.Repeat(spaceSeparator, midStrLength) + letter
		return frontStr + midStr
	}
	return frontStr
}

func MirrorDiamond(startLetter rune) []string {
	numberChars := int(startLetter - uppercaseChar)
	if numberChars < 2 { // Cannot have a diamond of less than 3 letters.
		return nil
	}

	outputArraySize := (numberChars * 2) + 1
	rows := make([]string, outputArraySize)

	i := 0
	for ; i <= numberChars; i++ {
		rows[i] = buildRow(i, numberChars)
	}

	j := i
	for ; j < outputArraySize; j++ { // inverse around the middle row.
		i = i - 1
		rows[j] = rows[i - 1]
	}
	return rows
}

func main() {
	rows := MirrorDiamond(rune('D'))
	for _, row := range rows {
		println(row)
	}
}