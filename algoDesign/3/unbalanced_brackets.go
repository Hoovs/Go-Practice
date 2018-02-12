package algoDesign

import (
    "strings"
    "Go-Practice/structures/stack"
)

type Node struct {
    value string
}

// Normally it would be best to check if the length is div by 2 and return non 0 if so. However, the
// requirement is to return the index.
func FindUnbalancedBracketIndexStack(brackets string) int {
    if len(brackets) <= 0 {return -1}

    correspondingBrace := map[string]string {"[":"]", "(":")", "{":"}"}
    bracketStack := stack.NewStack()
    chars := strings.Split(brackets, "")
    count := 0

    for i := 0; i < len(chars); i++ {
        count++
        val, ok := correspondingBrace[chars[i]]
        if ok {
            bracketStack.Push(&stack.Element{val})
        } else {
            popped, err := bracketStack.Pop()
            if err != nil {
                return count - 1
            }

            if popped == nil {
                return i
            }
            if chars[i] != popped.(*stack.Element).Value {
                return i
            }
        }
    }

    return -1
}

func FindUnbalancedBracketIndexIter(brackets string) int {
    numOpen, badLocation := 0, 0

    openBrack, closedBrack := int32(40), int32(41)

    for i, c := range brackets {
        if c == openBrack {
            if numOpen == 0 {
                badLocation = i
            }
            numOpen++
        } else if c == closedBrack {
            if numOpen == 0 {
                return i
            }
            numOpen = numOpen - 1
        }
    }
    if numOpen == 0 {
        return -1
    } else {
        return badLocation
    }
}