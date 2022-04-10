package ds

import (
	"fmt"
)

// LinkedList implements a doubly linked list.
type LinkedList[T comparable] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// node implements a node to be appended or prepended to a linked list. Contains data and pointers to the previous
// and next nodes.
type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

// Length returns the length of the linked list.
func (l *LinkedList[T]) Length() int {
	return l.length
}

// Prepend adds a new node to the beginning of the list.
func (l *LinkedList[T]) Prepend(value T) {
	n := &node[T]{data: value}
	prevHead := l.head
	n.next = prevHead
	l.head = n
	if prevHead == nil {
		l.tail = n
	} else {
		n.prev = prevHead
	}
	l.length++
}

func (l *LinkedList[T]) Append(value T) {
	n := &node[T]{data: value}
	prevTail := l.tail
	n.prev = prevTail
	l.tail = n
	if prevTail == nil {
		l.head = n
	} else {
		prevTail.next = n
	}
	l.length++
}

// Print returns a string of all node data in the linked list separated by ", "
func (l LinkedList[T]) Print() string {
	var llData string
	currentNode := l.head
	for currentNode != nil {
		llData += fmt.Sprintf("%+v, ", currentNode.data)
		currentNode = currentNode.next
	}

	llData = llData[:len(llData)-2]
	return llData
}

// Delete removes a node from the linked list by value. Returns the value of the deleted node.
func (l *LinkedList[T]) Delete(val T) (T, bool) {
	var zeroedReturnValue T
	if l.length == 0 {
		return zeroedReturnValue, false
	}
	if l.head.data == val {
		returnData := l.head.data
		l.head = l.head.next
		l.length--
		return returnData, true
	}
	prev := l.head
	currentNode := prev.next
	for currentNode != nil {
		nodeData := currentNode.data
		if nodeData == val {
			prev.next = currentNode.next
			l.length--
			return nodeData, true
		}
		prev = currentNode
		currentNode = currentNode.next
	}
	return zeroedReturnValue, false
}

// package ds

// import (
// 	"fmt"
// )

// // LinkedList implements a doubly linked list.
// type LinkedList[T any] struct {
// 	head   *node[T]
// 	tail   *node[T]
// 	length int
// }

// // node implements a node to be appended or prepended to a linked list. Contains data and pointers to the previous
// // and next nodes.
// type node[T any] struct {
// 	data T
// 	next *node[T]
// 	prev *node[T]
// }

// // Length returns the length of the linked list.
// func (l *LinkedList[T]) Length() int {
// 	return l.length
// }

// // Prepend adds a new node to the beginning of the list.
// func (l *LinkedList[T]) Prepend(value T) {
// 	n := &node[T]{data: value}
// 	prevHead := l.head
// 	n.next = prevHead
// 	l.head = n
// 	if prevHead == nil {
// 		l.tail = n
// 	} else {
// 		n.prev = prevHead
// 	}
// 	l.length++
// }

// func (l *LinkedList) Append(value int) {
// 	n := &node{data: value}
// 	prevTail := l.tail
// 	n.prev = prevTail
// 	l.tail = n
// 	if prevTail == nil {
// 		l.head = n
// 	} else {
// 		prevTail.next = n
// 	}
// 	l.length++
// }

// // Print returns a string of all node data in the linked list separated by ", "
// func (l LinkedList) Print() string {
// 	var llData string
// 	currentNode := l.head
// 	for currentNode != nil {
// 		llData += fmt.Sprintf("%+v, ", currentNode.data)
// 		currentNode = currentNode.next
// 	}

// 	llData = llData[:len(llData)-2]
// 	return llData
// }

// // Delete removes a node from the linked list by value. Returns the value of the deleted node.
// func (l *LinkedList) Delete(val int) (int, bool) {
// 	if l.length == 0 {
// 		return 0, false
// 	}
// 	if l.head.data == val {
// 		returnData := l.head.data
// 		l.head = l.head.next
// 		l.length--
// 		return returnData, true
// 	}
// 	prev := l.head
// 	currentNode := prev.next
// 	for currentNode != nil {
// 		nodeData := currentNode.data
// 		if nodeData == val {
// 			prev.next = currentNode.next
// 			l.length--
// 			return nodeData, true
// 		}
// 		prev = currentNode
// 		currentNode = currentNode.next
// 	}
// 	return 0, false
// }
