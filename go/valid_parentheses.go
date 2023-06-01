package main

import "container/list"

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	var last_index int = len(s) - 1
	var next_close byte
	stack := list.New()

	if s[0] == '(' {
		next_close = ')'

	} else if s[0] == '{' {
		next_close = '}'
	} else {
		next_close = ']'
	}
	stack.PushFront(next_close)
	for index := 1; index < len(s)-1; index++ {

		if s[index] == ')' || s[index] == '}' || s[index] == ']' {

			if stack.Len() <= 0 || s[index] != stack.Front().Value {
				return false
			}
			stack.Remove(stack.Front())
			continue
		}
		if s[index] == '(' {
			next_close = ')'
		} else if s[index] == '{' {
			next_close = '}'
		} else {
			next_close = ']'
		}
		stack.PushFront(next_close)
	}

	if stack.Len() > 1 {
		return false
	}
	return s[last_index] == stack.Front().Value
}
