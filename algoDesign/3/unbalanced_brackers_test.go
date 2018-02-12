package algoDesign

import "testing"

func TestFindUnbalanced(t *testing.T) {
    // VALID CASES
    if val := FindUnbalancedBracketIndexStack("[]"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexStack(""); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexStack("((()))"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexStack("((())())()"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexStack("[]()"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    //INVALID CASES
    if val := FindUnbalancedBracketIndexStack("}[]"); val != 0 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexStack("[]}{"); val != 2 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexStack("())"); val != 2 {
        t.Error("Unbalanced bracket found at: ", val)
    }
}

func TestFindUnbalancedBracketIndexIter(t *testing.T) {
    // VALID CASES
    if val := FindUnbalancedBracketIndexIter("()"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexIter(""); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexIter("((()))"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexIter("((())())()"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexIter("()()"); val != -1 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    //INVALID CASES
    if val := FindUnbalancedBracketIndexIter(")()"); val != 0 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexIter("())("); val != 2 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexIter("())"); val != 2 {
        t.Error("Unbalanced bracket found at: ", val)
    }

    if val := FindUnbalancedBracketIndexIter("((()))((()))("); val != 12 {
        t.Error("Unbalanced bracket found at: ", val)
    }
}
