package stack

import (
	"fmt"
)

type Stack struct {
	data []byte
	top  int
	size int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]byte, 0),
		top:  0,
	}
}

// 如果有必要可以再维护一个size
func (s *Stack) IsFull() bool {
	if s.top == s.size {
		return true
	}
	return false
}

func (s *Stack) IsEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(e byte) bool {
	s.data = append(s.data, e)

	s.top++
	return true
}

func (s *Stack) Pop() (flag bool, ret byte) {
	if s.IsEmpty() {
		return false, ret
	}

	ret = s.data[s.top-1]
	s.data = s.data[:s.top-1]
	s.top--

	return true, ret
}

func isValids(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Push(s[i])
		}
		if s[i] == ')' {
			if flag, top := stack.Pop(); flag == true && top == '(' {
				continue
			} else {
				return false
			}
		}
		if s[i] == '{' {
			stack.Push(s[i])
		}
		if s[i] == '}' {
			if flag, top := stack.Pop(); flag == true && top == '{' {
				continue
			} else {
				return false
			}
		}
		if s[i] == '[' {
			stack.Push(s[i])
		}
		if s[i] == ']' {
			if flag, top := stack.Pop(); flag == true && top == '[' {
				continue
			} else {
				return false
			}
		}
	}

	fmt.Println(stack)

	if stack.IsEmpty() {
		return true
	}
	return false
}
