package ds

type number interface {
	int | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64
}

/*
Heap is an implementation of a max heap.
Index of left child is (parentIndex * 2) + 1
Index of right child is (parentIndex * 2) + 2

insert and extract time complexity is O(h), h being the height or O(log n) because the height and number of indexes have a logarithmic relationship.
*/
type Heap[T number] []T

// Insert adds an element to the heap
func (h *Heap[T]) Insert(key T) {
	*h = append(*h, key)
	indexOfLastKey := len(*h) - 1
	h.heapifyUp(indexOfLastKey)
}

// Extract returns the largest element from the heap and reorders the heap.
func (h *Heap[T]) Extract() (T, bool) {
	if len(*h) == 0 {
		return 0, false
	}
	extracted := (*h)[0]
	lastIndex := len(*h) - 1
	(*h)[0] = (*h)[lastIndex]
	*h = (*h)[:lastIndex]
	h.heapifyDown(0)
	return extracted, true
}

// Heapifyup heapifies from bottom to top.
func (h *Heap[T]) heapifyUp(index int) {
	for (*h)[parentIndex(index)] < (*h)[index] {
		h.swap(parentIndex(index), index)
		index = parentIndex(index)
	}
}

func (h *Heap[T]) heapifyDown(index int) {
	lastIndex := len(*h) - 1
	indexOfLeftChild, indexOfRightChild := leftChildIndex(index), rightChildIndex(index)
	var indexOfChildToCompare int

	for indexOfLeftChild <= lastIndex {
		// When leftChild is larger or is the only one present
		if indexOfLeftChild == lastIndex {
			indexOfChildToCompare = indexOfLeftChild
		} else if (*h)[indexOfLeftChild] > (*h)[indexOfRightChild] {
			indexOfChildToCompare = indexOfRightChild
			// When rightChild is larger
		} else {
			indexOfChildToCompare = indexOfRightChild
		}

		// If current node val is less than child, then swap
		if (*h)[index] < (*h)[indexOfChildToCompare] {
			h.swap(index, indexOfChildToCompare)
			index = indexOfChildToCompare
			indexOfLeftChild, indexOfRightChild = leftChildIndex(index), rightChildIndex(index)
		} else {
			return
		}
	}
}

// parentIndex accepts an index and returns the index of its parent.
func parentIndex(i int) int {
	return (i - 1) / 2
}

// leftChildIndex accepts an index and returns the index of its left child.
func leftChildIndex(i int) int {
	return (i * 2) + 1
}

// rightChildIndex accepts an index and returns the index of its right child.
func rightChildIndex(i int) int {
	return (i * 2) + 1
}

func (h *Heap[T]) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// package ds

// /*
// Heap is an implementation of a max heap.
// Index of left child is (parentIndex * 2) + 1
// Index of right child is (parentIndex * 2) + 2

// insert and extract time complexity is O(h), h being the height or O(log n) because the height and number of indexes have a logarithmic relationship.
// */
// type Heap []int

// // Insert adds an element to the heap
// func (h *Heap) Insert(key int) {
// 	*h = append(*h, key)
// 	indexOfLastKey := len(*h) - 1
// 	h.heapifyUp(indexOfLastKey)
// }

// // Extract returns the largest element from the heap and reorders the heap.
// func (h *Heap) Extract() (int, bool) {
// 	if len(*h) == 0 {
// 		return 0, false
// 	}
// 	extracted := (*h)[0]
// 	lastIndex := len(*h) - 1
// 	(*h)[0] = (*h)[lastIndex]
// 	*h = (*h)[:lastIndex]
// 	h.heapifyDown(0)
// 	return extracted, true
// }

// // Heapifyup heapifies from bottom to top.
// func (h *Heap) heapifyUp(index int) {
// 	for (*h)[parentIndex(index)] < (*h)[index] {
// 		h.swap(parentIndex(index), index)
// 		index = parentIndex(index)
// 	}
// }

// func (h *Heap) heapifyDown(index int) {
// 	lastIndex := len(*h) - 1
// 	indexOfLeftChild, indexOfRightChild := leftChildIndex(index), rightChildIndex(index)
// 	var indexOfChildToCompare int

// 	for indexOfLeftChild <= lastIndex {
// 		// When leftChild is larger or is the only one present
// 		if indexOfLeftChild == lastIndex {
// 			indexOfChildToCompare = indexOfLeftChild
// 		} else if (*h)[indexOfLeftChild] > (*h)[indexOfRightChild] {
// 			indexOfChildToCompare = indexOfRightChild
// 			// When rightChild is larger
// 		} else {
// 			indexOfChildToCompare = indexOfRightChild
// 		}

// 		// If current node val is less than child, then swap
// 		if (*h)[index] < (*h)[indexOfChildToCompare] {
// 			h.swap(index, indexOfChildToCompare)
// 			index = indexOfChildToCompare
// 			indexOfLeftChild, indexOfRightChild = leftChildIndex(index), rightChildIndex(index)
// 		} else {
// 			return
// 		}
// 	}
// }

// // parentIndex accepts an index and returns the index of its parent.
// func parentIndex(i int) int {
// 	return (i - 1) / 2
// }

// // leftChildIndex accepts an index and returns the index of its left child.
// func leftChildIndex(i int) int {
// 	return (i * 2) + 1
// }

// // rightChildIndex accepts an index and returns the index of its right child.
// func rightChildIndex(i int) int {
// 	return (i * 2) + 1
// }

// func (h *Heap) swap(i, j int) {
// 	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
// }
