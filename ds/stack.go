package ds

// Stack implements a stack data structure. Methods are IsEmpty, Push, and Pop.
type Stack[T any] []T

// IsEmpty checks to see if the stack is empty.
func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}

// Push adds a new item to the stack.
func (s *Stack[T]) Push(n T) {
	*s = append(*s, n)
}

// Pop decreases the size of the stack by removing the top item. Returns the popped item and true if successful.
// If the stack is empty, returns the 0 value of the item, and false.
func (s *Stack[T]) Pop() (popped T, success bool) {
	if (*s).IsEmpty() {
		return
	}
	indexOfLastItem := len(*s) - 1
	popped = (*s)[indexOfLastItem]
	*s = (*s)[:indexOfLastItem]
	return popped, true
}

// Peek returns the value of the top item in the stack and if the operation was successful.
func (s *Stack[T]) Peek() (T, bool) {
	length := len(*s)
	if length == 0 {
		var zeroedReturn T
		return zeroedReturn, false
	}
	return (*s)[len(*s)-1], true
}

// package ds

// // Stack implements a stack data structure. Methods are IsEmpty, Push, and Pop.
// type Stack []int

// // IsEmpty checks to see if the stack is empty.
// func (s Stack) IsEmpty() bool {
// 	return len(s) == 0
// }

// // Push adds a new item to the stack.
// func (s *Stack) Push(n int) {
// 	*s = append(*s, n)
// }

// // Pop decreases the size of the stack by removing the top item. Returns the popped item and true if successful.
// // If the stack is empty, returns the 0 value of the item, and false.
// func (s *Stack) Pop() (popped int, success bool) {
// 	if (*s).IsEmpty() {
// 		return
// 	}
// 	indexOfLastItem := len(*s) - 1
// 	popped = (*s)[indexOfLastItem]
// 	*s = (*s)[:indexOfLastItem]
// 	return popped, true
// }

// // Peek returns the value of the top item in the stack and if the operation was successful.
// func (s *Stack) Peek() (int, bool) {
// 	length := len(*s)
// 	if length == 0 {
// 		return 0, false
// 	}
// 	return (*s)[len(*s)-1], true
// }
