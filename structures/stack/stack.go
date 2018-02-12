package stack

import (
    "sync"

    "github.com/pkg/errors"
)

type Element struct {
    Value interface{}
}

type Stack struct {
    lock sync.Mutex
    nodes []*Element
    size int
}

func NewStack() *Stack {
    return &Stack{lock: sync.Mutex{}, nodes: make([]*Element, 0), size: 0}
}

func (s *Stack) Size() int {
    return s.size
}

func (s *Stack) Push(e *Element) {
    s.lock.Lock()
    defer s.lock.Unlock()

    s.nodes = append(s.nodes[:s.size], e)
    s.size++
}

func (s *Stack) Pop() (interface{}, error) {
    s.lock.Lock()
    defer s.lock.Unlock()

    if s.size == 0 {
        return nil, errors.New("Empty Stack")
    }

    s.size--
    return s.nodes[s.size], nil
}