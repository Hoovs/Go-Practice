package arrays_stirngs

import "testing"

func TestUniqueCharsBitArray(t *testing.T) {
	isUnique := hasUniqueCharactersBitArray("t")
	if !isUnique {
		t.Error("Should be True")
	}

	isUnique = hasUniqueCharactersBitArray("")
	if !isUnique {
		t.Error("Should be True")
	}

	isUnique = hasUniqueCharactersBitArray("abc")
	if !isUnique {
		t.Error("Should be True")
	}

	isUnique = hasUniqueCharactersBitArray("abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")
	if !isUnique {
		t.Error("Should be True")
	}

	isUnique = hasUniqueCharactersBitArray("tT")
	if !isUnique {
		t.Error("Should be True")
	}

	isUnique = hasUniqueCharactersBitArray("tt")
	if isUnique {
		t.Error("Should be False")
	}

	isUnique = hasUniqueCharactersBitArray("not unique string")
	if isUnique {
		t.Error("Should be False")
	}

	isUnique = hasUniqueCharactersBitArray("abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()a")
	if isUnique {
		t.Error("Should be False")
	}
}
