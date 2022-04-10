package ds

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	myLinkedList := LinkedList[int]{}
	if myLinkedList.Length() != 0 {
		t.Errorf("Initialized linked list should have proper length. Want 0, got %d", myLinkedList.length)
	}
	myLinkedList.Prepend(5)

	if myLinkedList.tail.data != 5 {
		t.Errorf("Should have tail equal to head if linked list of size 1. Want 5, got %d", myLinkedList.tail.data)
	}

	myLinkedList.Prepend(7)
	myLinkedList.Prepend(2)
	if myLinkedList.Length() != 3 {
		t.Errorf("Linked list should have proper length. Want 3, got %d", myLinkedList.length)
	}

	deleted, success := myLinkedList.Delete(7)
	if !success {
		t.Errorf("Should be able to remove node that exists in linked list. Want success: true, got success: false")
	}
	if deleted != 7 {
		t.Errorf("Should remove proper node. Want 7, got %d", deleted)
	}
	if myLinkedList.Length() != 2 {
		t.Errorf("Should decrease in size after successful deletion. Want 2, got %d", myLinkedList.length)
	}

	deleted, success = myLinkedList.Delete(500)
	if success {
		t.Errorf("Should not be successful when attempting to delete node that does not exist. Want success: false, got success: true")
	}

	if deleted != 0 {
		t.Errorf("Deletion of node that does not exist should return 0 value. Want 0, got %d", deleted)
	}
	if myLinkedList.Length() != 2 {
		t.Errorf("Should not change size if node not successfully deleted. Want 2, got %d", myLinkedList.length)
	}

	deleted, success = myLinkedList.Delete(5)
	if !success {
		t.Errorf("Should be able to delete the head. Want success: true, got success: false")
	}
	if deleted != 5 {
		t.Errorf("Should return proper value when deleting head. Want 5, got %d", deleted)
	}
	if myLinkedList.Length() != 1 {
		t.Errorf("Should have length decremented when deleting head. Want 1, got %d", myLinkedList.length)
	}

	deleted, success = myLinkedList.Delete(2)
	if !success {
		t.Errorf("Should be able to delete last item in list.")
	}
	if deleted != 2 {
		t.Errorf("Deletion of last value should return proper value. Want 2, got %d", deleted)
	}
	if myLinkedList.Length() != 0 {
		t.Errorf("Should have 0 length after deleting final node. Want 0, got %d", myLinkedList.length)
	}

	deleted, success = myLinkedList.Delete(2)
	if success {
		t.Error("Should not be able to delete from empty list")
	}
	if deleted != 0 {
		t.Errorf("Deleting from empty list should return 0 value. Want 0, got %d", deleted)
	}

	myLinkedList.Prepend(10)
	if myLinkedList.Length() != 1 {
		t.Errorf("Should have length of 1 after prepending to empty list. Want 1, got %d", myLinkedList.length)
	}

	myLinkedList.Prepend(15)
	myLinkedList.Append(3)
	myLinkedList.Append(8)
	llData := myLinkedList.Print()
	expectedLLData := "15, 10, 3, 8"
	if llData != expectedLLData {
		t.Errorf("Should append node to back of linked list. Want %s, got %s", expectedLLData, llData)
	}
	if myLinkedList.Length() != 4 {
		t.Errorf("Should have proper length after apending data. Want 4, got %d", myLinkedList.Length())
	}
}

// package ds

// import (
// 	"testing"
// )

// func TestLinkedList(t *testing.T) {
// 	myLinkedList := LinkedList{}
// 	if myLinkedList.Length() != 0 {
// 		t.Errorf("Initialized linked list should have proper length. Want 0, got %d", myLinkedList.length)
// 	}
// 	myLinkedList.Prepend(5)

// 	if myLinkedList.tail.data != 5 {
// 		t.Errorf("Should have tail equal to head if linked list of size 1. Want 5, got %d", myLinkedList.tail.data)
// 	}

// 	myLinkedList.Prepend(7)
// 	myLinkedList.Prepend(2)
// 	if myLinkedList.Length() != 3 {
// 		t.Errorf("Linked list should have proper length. Want 3, got %d", myLinkedList.length)
// 	}

// 	deleted, success := myLinkedList.Delete(7)
// 	if !success {
// 		t.Errorf("Should be able to remove node that exists in linked list. Want success: true, got success: false")
// 	}
// 	if deleted != 7 {
// 		t.Errorf("Should remove proper node. Want 7, got %d", deleted)
// 	}
// 	if myLinkedList.Length() != 2 {
// 		t.Errorf("Should decrease in size after successful deletion. Want 2, got %d", myLinkedList.length)
// 	}

// 	deleted, success = myLinkedList.Delete(500)
// 	if success {
// 		t.Errorf("Should not be successful when attempting to delete node that does not exist. Want success: false, got success: true")
// 	}

// 	if deleted != 0 {
// 		t.Errorf("Deletion of node that does not exist should return 0 value. Want 0, got %d", deleted)
// 	}
// 	if myLinkedList.Length() != 2 {
// 		t.Errorf("Should not change size if node not successfully deleted. Want 2, got %d", myLinkedList.length)
// 	}

// 	deleted, success = myLinkedList.Delete(5)
// 	if !success {
// 		t.Errorf("Should be able to delete the head. Want success: true, got success: false")
// 	}
// 	if deleted != 5 {
// 		t.Errorf("Should return proper value when deleting head. Want 5, got %d", deleted)
// 	}
// 	if myLinkedList.Length() != 1 {
// 		t.Errorf("Should have length decremented when deleting head. Want 1, got %d", myLinkedList.length)
// 	}

// 	deleted, success = myLinkedList.Delete(2)
// 	if !success {
// 		t.Errorf("Should be able to delete last item in list.")
// 	}
// 	if deleted != 2 {
// 		t.Errorf("Deletion of last value should return proper value. Want 2, got %d", deleted)
// 	}
// 	if myLinkedList.Length() != 0 {
// 		t.Errorf("Should have 0 length after deleting final node. Want 0, got %d", myLinkedList.length)
// 	}

// 	deleted, success = myLinkedList.Delete(2)
// 	if success {
// 		t.Error("Should not be able to delete from empty list")
// 	}
// 	if deleted != 0 {
// 		t.Errorf("Deleting from empty list should return 0 value. Want 0, got %d", deleted)
// 	}

// 	myLinkedList.Prepend(10)
// 	if myLinkedList.Length() != 1 {
// 		t.Errorf("Should have length of 1 after prepending to empty list. Want 1, got %d", myLinkedList.length)
// 	}

// 	myLinkedList.Prepend(15)
// 	myLinkedList.Append(3)
// 	myLinkedList.Append(8)
// 	llData := myLinkedList.Print()
// 	expectedLLData := "15, 10, 3, 8"
// 	if llData != expectedLLData {
// 		t.Errorf("Should append node to back of linked list. Want %s, got %s", expectedLLData, llData)
// 	}
// 	if myLinkedList.Length() != 4 {
// 		t.Errorf("Should have proper length after apending data. Want 4, got %d", myLinkedList.Length())
// 	}
// }
