package algoDesign

import "strings"

type Node struct {
    Value string
    index int
}

type Stack struct {
    nodes []*Node
    count int
}

func NewStack() *Stack {
    return &Stack{}
}

func (s *Stack) Push(n *Node) {
    s.nodes = append(s.nodes[:s.count], n)
    s.count++
}

func (s *Stack) Pop() *Node {
    if s.count == 0 {
        return nil
    }
    s.count--
    return s.nodes[s.count]
}

// Normally it would be best to check if the length is div by 2 and return non 0 if so. However, the
// requirement is to return the index.
func FindUnbalancedBracketIndex(brackets string) int {
    if len(brackets) <= 0 {return -1}

    correspondingBrace := map[string]string {"[":"]", "(":")", "{":"}"}

    stack := NewStack()

    chars := strings.Split(brackets, "")

    for i := 0; i < len(chars); i++ {
        val, ok := correspondingBrace[chars[i]]
        if ok {
            stack.Push(&Node{val, i})
        } else {
            popped := stack.Pop()
            if popped == nil {
                return i
            }
            if chars[i] != popped.Value {
                return i
            }
        }
    }

    if len(chars) %2 != 0 && stack.count != 0 { // Final guard for a non div 2 count but all other matching.
        return stack.Pop().index
    }

    return -1
}