package arrays_stirngs

import "testing"

func TestIsPermutation(t *testing.T) {
	perm := isPermutationSort("ab", "a")
	if perm {
		t.Error("Should not have been a permutation")
	}

	perm = isPermutationSort("a", "a")
	if !perm {
		t.Error("Should have been a permutation")
	}

	perm = isPermutationSort("golf", "flog")
	if !perm {
		t.Error("Should have been a permutation")
	}
}

func TestIsPermtationBit(t *testing.T) {
	perm := isPermutationBit("ab", "a")
	if perm {
		t.Error("Should not have been a permutation")
	}

	perm = isPermutationBit("a", "a")
	if !perm {
		t.Error("Should have been a permutation")
	}

	perm = isPermutationBit("golf", "flog")
	if !perm {
		t.Error("Should have been a permutation")
	}
}