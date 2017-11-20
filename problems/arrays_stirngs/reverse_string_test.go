package arrays_stirngs

import "testing"

func TestReverseString(t *testing.T) {
	outputStr := reverseString("R")
	if outputStr != "R" {
		t.Error("Should have been R but was:", outputStr)
	}

	outputStr = reverseString("RR")
	if outputStr != "RR" {
		t.Error("Should have been RR but was:", outputStr)
	}

	outputStr = reverseString("BA")
	if outputStr != "AB" {
		t.Error("Should have been AB but was:", outputStr)
	}

	outputStr = reverseString("abcCBA")
	if outputStr != "ABCcba" {
		t.Error("Should have been ABCcba but was:", outputStr)
	}
}
