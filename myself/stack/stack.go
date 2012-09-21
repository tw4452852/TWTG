package stack

import (
	"errors"
	"strconv"
	"fmt"
)

const (
	LIMIT = 4
)

type obj interface{}

type Stack struct {
	cells	[LIMIT]obj
	index	int
}

func (s *Stack) Push(n obj) error {
	if s.IsFull() {
		return errors.New("Stack is full")
	}
	s.cells[s.index] = n
	s.index++
	return nil
}

func (s *Stack) Pop() (obj, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	s.index--
	return s.cells[s.index], nil
}

func (s *Stack) Len() int {
	return s.index
}

func (s *Stack) IsEmpty() bool {
	return s.index == 0
}

func (s *Stack) IsFull() bool {
	return s.index == LIMIT
}

func (s *Stack) Top() obj {
	return s.cells[s.index - 1]
}

func (s *Stack) String() (result string) {
	for i := 0; i < s.index; i++ {
		result += " [" + strconv.Itoa(i) + ":"
		switch s.cells[i].(type) {
		case int:
			result += strconv.Itoa(s.cells[i].(int))
		case string:
			result += s.cells[i].(string)
		default:
			result += fmt.Sprintf("%v", s.cells[i])
		}
		result += "]"
	}
	return result
}

func NewStack() *Stack {
	return new(Stack)
}
