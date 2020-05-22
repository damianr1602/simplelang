package util

import (
	logger "github.com/damianr1602/simplelang/logging"
)

// Stack generics
type Stack []interface{}

// Stack object
// type Stack []Object

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(val interface{}) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
	logger.Log.Println("Stack elements", s)

}

// Pop remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return Stack{}
	}
	index := len(*s) - 1
	element := (*s)[index]
	logger.Log.Println("Stack elements before pop", s)
	*s = (*s)[:index]

	return element
}
