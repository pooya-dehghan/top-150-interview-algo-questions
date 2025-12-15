package main

import "fmt"

type customeStack struct {
	stack []string
}

func (s *customeStack) push(val string) {
	s.stack = append(s.stack, val)
}

func (s *customeStack) pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *customeStack) observe() string {
	if len(s.stack) != 0 {
		return s.stack[len(s.stack)-1]
	}
	return ""
}

func isValid(s string) bool {
	stack := customeStack{}
	for _, val := range s {
		if stack.observe() == "(" && string(val) == ")" {
			stack.pop()
			continue
		}
		if stack.observe() == "[" && string(val) == "]" {
			stack.pop()
			continue
		}
		if stack.observe() == "{" && string(val) == "}" {
			stack.pop()
			continue
		}
		stack.push(string(val))
	}
	fmt.Println(len(stack.stack))
	return len(stack.stack) == 0
}

func main() {
	isValid("()")
}
