package main

import "testing"

func TestMirrorDiamond(t *testing.T) {
	returnedRow := buildRow(0, 2) // Base case just to verify working
	expected := "  A"
	if returnedRow != "  A" {
		t.Errorf("Wrong row generation expected: %s, got: %s", expected, returnedRow)
	}

	returnedRows := MirrorDiamond('D') // Base case
	expectedRows := []string{"   A",
	                         "  B B",
	                         " C   C",
	                         "D     D",
	                         " C   C",
		                     "  B B",
		                     "   A"}
	for i, row := range returnedRows {
		if row != expectedRows[i] {
			t.Errorf("Expected %s, got %s", expectedRows[i], row)
		}
	}
}
