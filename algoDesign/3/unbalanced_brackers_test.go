package algoDesign

import "testing"

func TestFindUnbalanced(t *testing.T) {
    // VALID CASES
    if val := FindUnbalancedBracketIndex("[]"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndex(""); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndex("((()))"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndex("((())())()"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndex("[]()"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    //INVALID CASES
    if val := FindUnbalancedBracketIndex("}[]"); val != 0 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndex("[]}{"); val != 2 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndex("(()()"); val != 0 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndex("())"); val != 2 {
        t.Error("Unbalanced bracket found at: ", val)
    }
}
