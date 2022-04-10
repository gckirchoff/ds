package ds

import "testing"

func TestStack(t *testing.T) {
	testStack := Stack[int]{}

	isInitializedStackEmpty := testStack.IsEmpty()
	if !isInitializedStackEmpty {
		t.Errorf("Initialized stack is not empty: %+v", testStack)
	}

	testStack.Push(1)
	testStack.Push(2)
	testStack.Push(3)
	if len(testStack) != 3 {
		t.Errorf("Stack: %v is of length %d, expected 3", testStack, len(testStack))
	}

	popped, success := testStack.Pop()
	if popped != 3 {
		t.Errorf("Incorrect number was popped from the stack. Want 3, got %d", popped)
	}
	if !success {
		t.Error("Stack was not successfully popped.")
	}

	isEmpty := testStack.IsEmpty()
	if isEmpty {
		t.Errorf("Stack is incorrectly flagged as empty. Want [1 2], got %v", testStack)
	}

	testStack.Pop()
	testStack.Pop()

	isEmpty = testStack.IsEmpty()
	if !isEmpty {
		t.Errorf("Stack incorrectly flagged as not empty. Want [], got %v", testStack)
	}

	popped, success = testStack.Pop()
	if success {
		t.Errorf("Empty stack should not be successfully popped. Got success: true, want success: false")
	}
	if popped != 0 {
		t.Errorf("0 Value not returned from empty stack. Want 0, got %d", popped)
	}

	peeked, success := testStack.Peek()
	if success {
		t.Error("Should not be able to peak into empty stack")
	}
	if peeked != 0 {
		t.Errorf("Should return 0 value for peeked value if stack is empty. Want 0, got %d", peeked)
	}

	testStack.Push(5)
	testStack.Push(10)
	peeked, success = testStack.Peek()
	if !success {
		t.Error("Should be able to peek value in stack")
	}
	if peeked != 10 {
		t.Errorf("Should return correct value when peeking. Want 10, got %d", peeked)
	}
}

// package ds

// import "testing"

// func TestStack(t *testing.T) {
// 	testStack := Stack{}

// 	isInitializedStackEmpty := testStack.IsEmpty()
// 	if !isInitializedStackEmpty {
// 		t.Errorf("Initialized stack is not empty: %+v", testStack)
// 	}

// 	testStack.Push(1)
// 	testStack.Push(2)
// 	testStack.Push(3)
// 	if len(testStack) != 3 {
// 		t.Errorf("Stack: %v is of length %d, expected 3", testStack, len(testStack))
// 	}

// 	popped, success := testStack.Pop()
// 	if popped != 3 {
// 		t.Errorf("Incorrect number was popped from the stack. Want 3, got %d", popped)
// 	}
// 	if !success {
// 		t.Error("Stack was not successfully popped.")
// 	}

// 	isEmpty := testStack.IsEmpty()
// 	if isEmpty {
// 		t.Errorf("Stack is incorrectly flagged as empty. Want [1 2], got %v", testStack)
// 	}

// 	testStack.Pop()
// 	testStack.Pop()

// 	isEmpty = testStack.IsEmpty()
// 	if !isEmpty {
// 		t.Errorf("Stack incorrectly flagged as not empty. Want [], got %v", testStack)
// 	}

// 	popped, success = testStack.Pop()
// 	if success {
// 		t.Errorf("Empty stack should not be successfully popped. Got success: true, want success: false")
// 	}
// 	if popped != 0 {
// 		t.Errorf("0 Value not returned from empty stack. Want 0, got %d", popped)
// 	}

// 	peeked, success := testStack.Peek()
// 	if success {
// 		t.Error("Should not be able to peak into empty stack")
// 	}
// 	if peeked != 0 {
// 		t.Errorf("Should return 0 value for peeked value if stack is empty. Want 0, got %d", peeked)
// 	}

// 	testStack.Push(5)
// 	testStack.Push(10)
// 	peeked, success = testStack.Peek()
// 	if !success {
// 		t.Error("Should be able to peek value in stack")
// 	}
// 	if peeked != 10 {
// 		t.Errorf("Should return correct value when peeking. Want 10, got %d", peeked)
// 	}
// }
