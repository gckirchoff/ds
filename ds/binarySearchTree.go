package ds

import "fmt"

// Search/Insertion is O(h) or O(log n) at best or if completely unbalanced, O(n)

// treeNode implements a node to be inserted into a binary search tree. Contains data and pointers to
// its left and right children.
type treeNode[T any] struct {
	key   int
	value T
	left  *treeNode[T]
	right *treeNode[T]
}

// BinarySearchTree implements a binary search tree.
type BinarySearchTree[T any] struct {
	root *treeNode[T]
}

// Insert adds a new node to a binary search tree. Returns a boolean if successful or not.
func (b *BinarySearchTree[T]) Insert(key int, data T) bool {
	n := &treeNode[T]{key: key, value: data}
	if b.root == nil {
		b.root = n
		return true
	}

	newRoot, success := insertNode(b.root, n)
	b.root = newRoot
	return success
}

// Search finds a key and returns its value. Returns if the operation was successful.
func (b *BinarySearchTree[T]) Search(key int) (T, bool) {
	success := true

	var find func(r *treeNode[T], key int) T
	find = func(r *treeNode[T], key int) T {
		if r == nil {
			success = false
			var zeroedReturn T
			return zeroedReturn
		}
		if key == r.key {
			return r.value
		}

		if key < r.key {
			return find(r.left, key)
		} else {
			return find(r.right, key)
		}
	}

	return find(b.root, key), success
}

func insertNode[T any](r, n *treeNode[T]) (*treeNode[T], bool) {
	success := true

	var traverse func(r, n *treeNode[T]) *treeNode[T]
	traverse = func(r, n *treeNode[T]) *treeNode[T] {
		if r == nil {
			return n
		}
		if n.key < r.key {
			r.left = traverse(r.left, n)
		} else if n.key > r.key {
			r.right = traverse(r.right, n)
		} else {
			success = false
			return r
		}

		return r
	}

	return traverse(r, n), success
}

func (b *BinarySearchTree[T]) PrintInOrder() string {
	var values string

	var inOrder func(n *treeNode[T])
	inOrder = func(n *treeNode[T]) {
		if n == nil {
			return
		}
		inOrder(n.left)
		values += fmt.Sprintf("Key of %d: %+v", n.key, n.value) + ", "
		inOrder(n.right)
	}
	inOrder(b.root)
	values = values[:len(values)-2]
	return values
}

func (b *BinarySearchTree[T]) PrintPreOrder() string {
	var values string

	var inOrder func(n *treeNode[T])
	inOrder = func(n *treeNode[T]) {
		if n == nil {
			return
		}
		values += fmt.Sprintf("Key of %d: %+v", n.key, n.value) + ", "
		inOrder(n.left)
		inOrder(n.right)
	}
	inOrder(b.root)
	values = values[:len(values)-2]
	return values
}

func (b *BinarySearchTree[T]) PrintPostOrder() string {
	var values string

	var inOrder func(n *treeNode[T])
	inOrder = func(n *treeNode[T]) {
		if n == nil {
			return
		}
		inOrder(n.left)
		inOrder(n.right)
		values += fmt.Sprintf("Key of %d: %+v", n.key, n.value) + ", "
	}
	inOrder(b.root)
	values = values[:len(values)-2]
	return values
}

func (b *BinarySearchTree[T]) Remove(key int) (*treeNode[T], bool) {
	var removedNode treeNode[T]
	var timesEditedRemovedNode int
	var success bool

	var removeRec func(n *treeNode[T], key int) *treeNode[T]
	removeRec = func(n *treeNode[T], key int) *treeNode[T] {
		if n == nil {
			return n
		}
		if key < n.key {
			n.left = removeRec(n.left, key)
		}
		if key > n.key {
			n.right = removeRec(n.right, key)
		}

		if key == n.key {
			if timesEditedRemovedNode == 0 {
				removedNode = *n
				success = true
				timesEditedRemovedNode++
			}
			if n.left == nil && n.right == nil {
				n = nil
				return n
			}
			if n.left == nil && n.right != nil {
				temp := n.right
				n = nil
				n = temp
				return n
			}
			if n.left != nil && n.right == nil {
				temp := n.left
				n = nil
				n = temp
				return n
			}

			tempNode := minValue(n.right)
			n.key, n.value = tempNode.key, tempNode.value
			n.right = removeRec(n.right, tempNode.key)
		}
		return n
	}
	b.root = removeRec(b.root, key)
	removedNode.left = nil
	removedNode.right = nil

	if !success {
		return nil, success
	}
	return &removedNode, success
}

func minValue[T any](n *treeNode[T]) *treeNode[T] {
	currentNode := n
	for currentNode.left != nil && currentNode != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

// package ds

// import "fmt"

// // Search/Insertion is O(h) or O(log n) at best or if completely unbalanced, O(n)

// // treeNode implements a node to be inserted into a binary search tree. Contains data and pointers to
// // its left and right children.
// type treeNode struct {
// 	key   int
// 	value int
// 	left  *treeNode
// 	right *treeNode
// }

// // BinarySearchTree implements a binary search tree.
// type BinarySearchTree struct {
// 	root *treeNode
// }

// // Insert adds a new node to a binary search tree. Returns a boolean if successful or not.
// func (b *BinarySearchTree) Insert(key, data int) bool {
// 	n := &treeNode{key: key, value: data}
// 	if b.root == nil {
// 		b.root = n
// 		return true
// 	}

// 	newRoot, success := insertNode(b.root, n)
// 	b.root = newRoot
// 	return success
// }

// // Search finds a key and returns its value. Returns if the operation was successful.
// func (b *BinarySearchTree) Search(key int) (int, bool) {
// 	success := true

// 	var find func(r *treeNode, key int) int
// 	find = func(r *treeNode, key int) int {
// 		if r == nil {
// 			success = false
// 			return 0
// 		}
// 		if key == r.key {
// 			return r.value
// 		}

// 		if key < r.key {
// 			return find(r.left, key)
// 		} else {
// 			return find(r.right, key)
// 		}
// 	}

// 	return find(b.root, key), success
// }

// func insertNode(r, n *treeNode) (*treeNode, bool) {
// 	success := true

// 	var traverse func(r, n *treeNode) *treeNode
// 	traverse = func(r, n *treeNode) *treeNode {
// 		if r == nil {
// 			return n
// 		}
// 		if n.key < r.key {
// 			r.left = traverse(r.left, n)
// 		} else if n.key > r.key {
// 			r.right = traverse(r.right, n)
// 		} else {
// 			success = false
// 			return r
// 		}

// 		return r
// 	}

// 	return traverse(r, n), success
// }

// func (b *BinarySearchTree) PrintInOrder() string {
// 	var values string

// 	var inOrder func(n *treeNode)
// 	inOrder = func(n *treeNode) {
// 		if n == nil {
// 			return
// 		}
// 		inOrder(n.left)
// 		values += fmt.Sprintf("Key of %d: %+v", n.key, n.value) + ", "
// 		inOrder(n.right)
// 	}
// 	inOrder(b.root)
// 	values = values[:len(values)-2]
// 	return values
// }

// func (b *BinarySearchTree) PrintPreOrder() string {
// 	var values string

// 	var inOrder func(n *treeNode)
// 	inOrder = func(n *treeNode) {
// 		if n == nil {
// 			return
// 		}
// 		values += fmt.Sprintf("Key of %d: %+v", n.key, n.value) + ", "
// 		inOrder(n.left)
// 		inOrder(n.right)
// 	}
// 	inOrder(b.root)
// 	values = values[:len(values)-2]
// 	return values
// }

// func (b *BinarySearchTree) PrintPostOrder() string {
// 	var values string

// 	var inOrder func(n *treeNode)
// 	inOrder = func(n *treeNode) {
// 		if n == nil {
// 			return
// 		}
// 		inOrder(n.left)
// 		inOrder(n.right)
// 		values += fmt.Sprintf("Key of %d: %+v", n.key, n.value) + ", "
// 	}
// 	inOrder(b.root)
// 	values = values[:len(values)-2]
// 	return values
// }

// func (b *BinarySearchTree) Remove(key int) (*treeNode, bool) {
// 	var removedNode treeNode
// 	var timesEditedRemovedNode int
// 	var success bool

// 	var removeRec func(n *treeNode, key int) *treeNode
// 	removeRec = func(n *treeNode, key int) *treeNode {
// 		if n == nil {
// 			return n
// 		}
// 		if key < n.key {
// 			n.left = removeRec(n.left, key)
// 		}
// 		if key > n.key {
// 			n.right = removeRec(n.right, key)
// 		}

// 		if key == n.key {
// 			if timesEditedRemovedNode == 0 {
// 				removedNode = *n
// 				success = true
// 				timesEditedRemovedNode++
// 			}
// 			if n.left == nil && n.right == nil {
// 				n = nil
// 				return n
// 			}
// 			if n.left == nil && n.right != nil {
// 				temp := n.right
// 				n = nil
// 				n = temp
// 				return n
// 			}
// 			if n.left != nil && n.right == nil {
// 				temp := n.left
// 				n = nil
// 				n = temp
// 				return n
// 			}

// 			tempNode := minValue(n.right)
// 			n.key, n.value = tempNode.key, tempNode.value
// 			n.right = removeRec(n.right, tempNode.key)
// 		}
// 		return n
// 	}
// 	b.root = removeRec(b.root, key)
// 	removedNode.left = nil
// 	removedNode.right = nil

// 	if !success {
// 		return nil, success
// 	}
// 	return &removedNode, success
// }

// func minValue(n *treeNode) *treeNode {
// 	currentNode := n
// 	for currentNode.left != nil && currentNode != nil {
// 		currentNode = currentNode.left
// 	}
// 	return currentNode
// }
