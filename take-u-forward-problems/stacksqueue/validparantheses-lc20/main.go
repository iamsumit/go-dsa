package main

import "fmt"

// https://takeuforward.org/data-structure/check-for-balanced-parentheses/
func isValid(s string) bool {
	stack := []rune{}

	hashMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, str := range s {
		switch str {
		case '(', '{', '[':
			stack = append(stack, str)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != hashMap[str] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// type MyStack []rune

// func (m *MyStack) Push(data rune) {
// 	*m = append(*m, data)
// }

// // LIFO
// func (m *MyStack) Pop() rune {
// 	vals := *m
// 	n := len(vals)
// 	x := vals[n-1]
// 	*m = vals[0 : n-1]

// 	return x
// }

// func (m *MyStack) Empty() bool {
// 	return len(*m) == 0
// }

// func isValid(s string) bool {
// 	openParan := func(r rune) rune {
// 		switch r {
// 		case ')':
// 			return '('
// 		case '}':
// 			return '{'
// 		case ']':
// 			return '['
// 		}

// 		return '0'
// 	}

// 	myStack := &MyStack{}
// 	for _, str := range s {
// 		p := openParan(str)
// 		if p == '0' {
// 			myStack.Push(str)
// 			continue
// 		}

// 		if myStack.Empty() {
// 			return false
// 		}

// 		if p != myStack.Pop() {
// 			return false
// 		}
// 	}

// 	return myStack.Empty()
// }

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(())"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("()([]{})"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("]"))
}
