package ds

import "testing"

func TestHeap(t *testing.T) {
	testHeap := Heap[int]{}
	testHeap.Insert(5)
	testHeap.Insert(7)
	testHeap.Insert(15)
	testHeap.Insert(2)
	if testHeap[0] != 15 {
		t.Errorf("First index should contain highest value. Want [15 5 7 2], got: %v", testHeap)
	}
	if testHeap[3] != 2 {
		t.Errorf("Last index should contain lowest value. Want [15 5 7 2], got: %v", testHeap)
	}
	if len(testHeap) != 4 {
		t.Errorf("Heap should have correct length. Want 4, got: %v", len(testHeap))
	}

	extracted, success := testHeap.Extract()
	if !success {
		t.Errorf("Heap should be able to be extracted from. Got success: false, want success: true. Heap: %v", testHeap)
	}

	if extracted != 15 {
		t.Errorf("Extracted should be of maximum value in the heap. got %d, want 15", extracted)
	}

	if len(testHeap) != 3 {
		t.Errorf("Heap should be decremented in size after extracting. Got %d, want 3", len(testHeap))
	}

	testHeap.Extract()
	testHeap.Extract()
	testHeap.Extract()

	extracted, success = testHeap.Extract()
	if success {
		t.Errorf("Empty heap should not be able to be extracted from. Got success: true, want success: false")
	}
	if extracted != 0 {
		t.Errorf("Empty heap should return 0 value. Got %d, want 0", extracted)
	}
}

// package ds

// import "testing"

// func TestHeap(t *testing.T) {
// 	testHeap := Heap{}
// 	testHeap.Insert(5)
// 	testHeap.Insert(7)
// 	testHeap.Insert(15)
// 	testHeap.Insert(2)
// 	if testHeap[0] != 15 {
// 		t.Errorf("First index should contain highest value. Want [15 5 7 2], got: %v", testHeap)
// 	}
// 	if testHeap[3] != 2 {
// 		t.Errorf("Last index should contain lowest value. Want [15 5 7 2], got: %v", testHeap)
// 	}
// 	if len(testHeap) != 4 {
// 		t.Errorf("Heap should have correct length. Want 4, got: %v", len(testHeap))
// 	}

// 	extracted, success := testHeap.Extract()
// 	if !success {
// 		t.Errorf("Heap should be able to be extracted from. Got success: false, want success: true. Heap: %v", testHeap)
// 	}

// 	if extracted != 15 {
// 		t.Errorf("Extracted should be of maximum value in the heap. got %d, want 15", extracted)
// 	}

// 	if len(testHeap) != 3 {
// 		t.Errorf("Heap should be decremented in size after extracting. Got %d, want 3", len(testHeap))
// 	}

// 	testHeap.Extract()
// 	testHeap.Extract()
// 	testHeap.Extract()

// 	extracted, success = testHeap.Extract()
// 	if success {
// 		t.Errorf("Empty heap should not be able to be extracted from. Got success: true, want success: false")
// 	}
// 	if extracted != 0 {
// 		t.Errorf("Empty heap should return 0 value. Got %d, want 0", extracted)
// 	}
// }
