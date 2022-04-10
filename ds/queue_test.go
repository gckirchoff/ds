package ds

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue[int]{}

	isEmpty := queue.IsEmpty()
	if !isEmpty {
		t.Errorf("Initialized empty queue should be empty. Want [], got %v", queue)
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	isEmpty = queue.IsEmpty()
	if isEmpty {
		t.Errorf("Queue should not be empty. Want [1 2], got %v", queue)
	}

	dequeued, success := queue.Dequeue()
	if !success {
		t.Errorf("Dequeue not successful. Current queue should be [2], got %v", queue)
	}
	if dequeued != 1 {
		t.Errorf("Incorrect item dequeued. Want 1, got %d", dequeued)
	}

	queue.Dequeue()

	isEmpty = queue.IsEmpty()
	if !isEmpty {
		t.Errorf("Queue should be empty. Want [], got %v", queue)
	}

	dequeued, success = queue.Dequeue()
	if dequeued != 0 {
		t.Errorf("Improper value returned from what should be an empty queue. Want 0, got %d", dequeued)
	}
	if success {
		t.Errorf("Dequeuing from empty queue should not be successful. Want success: false, got success: true")
	}

}

// package ds

// import (
// 	"testing"
// )

// func TestQueue(t *testing.T) {
// 	queue := Queue{}

// 	isEmpty := queue.IsEmpty()
// 	if !isEmpty {
// 		t.Errorf("Initialized empty queue should be empty. Want [], got %v", queue)
// 	}

// 	queue.Enqueue(1)
// 	queue.Enqueue(2)
// 	isEmpty = queue.IsEmpty()
// 	if isEmpty {
// 		t.Errorf("Queue should not be empty. Want [1 2], got %v", queue)
// 	}

// 	dequeued, success := queue.Dequeue()
// 	if !success {
// 		t.Errorf("Dequeue not successful. Current queue should be [2], got %v", queue)
// 	}
// 	if dequeued != 1 {
// 		t.Errorf("Incorrect item dequeued. Want 1, got %d", dequeued)
// 	}

// 	queue.Dequeue()

// 	isEmpty = queue.IsEmpty()
// 	if !isEmpty {
// 		t.Errorf("Queue should be empty. Want [], got %v", queue)
// 	}

// 	dequeued, success = queue.Dequeue()
// 	if dequeued != 0 {
// 		t.Errorf("Improper value returned from what should be an empty queue. Want 0, got %d", dequeued)
// 	}
// 	if success {
// 		t.Errorf("Dequeuing from empty queue should not be successful. Want success: false, got success: true")
// 	}

// }
