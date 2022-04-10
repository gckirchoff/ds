package ds

// Queue implements a queue data structure.
type Queue[T any] []T

// IsEmpty returns a boolean indicating if the queue is empty.
func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}

// Enqueue adds a new item to the back of the queue.
func (q *Queue[T]) Enqueue(n T) {
	*q = append(*q, n)
}

// Dequeue removes an item from the front of the queue and if it was successful.
func (q *Queue[T]) Dequeue() (dequeued T, success bool) {
	if (*q).IsEmpty() {
		return
	}

	dequeued = (*q)[0]
	*q = (*q)[1:]
	return dequeued, true
}

// package ds

// // Queue implements a queue data structure.
// type Queue []int

// // IsEmpty returns a boolean indicating if the queue is empty.
// func (q Queue) IsEmpty() bool {
// 	return len(q) == 0
// }

// // Enqueue adds a new item to the back of the queue.
// func (q *Queue) Enqueue(n int) {
// 	*q = append(*q, n)
// }

// // Dequeue removes an item from the front of the queue and if it was successful.
// func (q *Queue) Dequeue() (dequeued int, success bool) {
// 	if (*q).IsEmpty() {
// 		return
// 	}

// 	dequeued = (*q)[0]
// 	*q = (*q)[1:]
// 	return dequeued, true
// }
