package structures


type node struct {
	val int
	next *node
}

type Stack struct {
	top *node
}

func InitStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() int {
	if s.top != nil {
		tmp := s.top
		s.top = s.top.next
		return tmp.val
	} else {
		panic("ops")
	}
}

func (s *Stack) Push(val int) {
	s.top = &node{val: val, next: s.top}
}
