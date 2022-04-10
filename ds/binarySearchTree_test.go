package ds

import (
	"regexp"
	"strings"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	myBST := BinarySearchTree[int]{}
	if myBST.root != nil {
		t.Error("Should be empty when instantiated.")
	}
	success := myBST.Insert(5, 1)
	if !success {
		t.Error("Should be able to insert first node.")
	}
	if myBST.root == nil {
		t.Error("Should not have root == nill after insertion of first node.")
	}
	rootKey := myBST.root.key
	rootData := myBST.root.value
	if rootKey != 5 {
		t.Errorf("Should have proper key at root. Want 5, got %d", rootKey)
	}
	if rootData != 1 {
		t.Errorf("Should have proper data at root. Want 1, got %d", rootData)
	}

	success = myBST.Insert(5, 4)
	if success {
		t.Error("Should not be able to insert key that already exists.")
	}
	rootData = myBST.root.value
	if rootData != 1 {
		t.Errorf("Should not change data at root if key already existed. Want 1, got %d", rootData)
	}

	val, success := myBST.Search(5)
	if !success {
		t.Error("Should be able to find node that exists")
	}

	if val != 1 {
		t.Errorf("Should retrieve proper value for found node. Want 1, got %d", val)
	}

	val, success = myBST.Search(2)
	if success {
		t.Error("Should not be able to retrieve node that doesn't exist.")
	}
	if val != 0 {
		t.Errorf("Should return 0 value for node that doesn't exist. Want 0, got %d", val)
	}

	success = myBST.Insert(10, 4)
	if !success {
		t.Error("Should be able to insert new node")
	}
	rightKey := myBST.root.right.key
	rightVal := myBST.root.right.value
	if rightKey != 10 {
		t.Errorf("Should have inserted 10 to the right of root. Want 10, got %d", rightKey)
	}
	if rightVal != 4 {
		t.Errorf("Inserted node should have correct vaule. Want 4, got %d", rightVal)
	}

	myBST.Insert(2, 1)
	leftKey := myBST.root.left.key
	leftVal := myBST.root.left.value
	if leftKey != 2 {
		t.Errorf("Should have inserted 10 to the left of root. Want 10, got %d", leftKey)
	}
	if leftVal != 1 {
		t.Errorf("Inserted node should have correct vaule. Want 4, got %d", leftVal)
	}

	success = myBST.Insert(12, 1)
	if !success {
		t.Error("Should be able to insert new node")
	}

	myBST.Insert(15, 1)
	success = myBST.Insert(12, 4)
	if success {
		t.Error("Should not be able to insert duplicate key")
	}

	val, success = myBST.Search(12)
	if !success {
		t.Error("Should be able to find key that exists")
	}
	if val != 1 {
		t.Errorf("Should return proper value for node. Want 1, got %d", val)
	}

	val, success = myBST.Search(11)
	if success {
		t.Error("Should not be able to find key that does not exist")
	}
	if val != 0 {
		t.Errorf("Should return 0 value for key that does not exist. Want 0, got %d", val)
	}

	inOrder := parseKeys(myBST.PrintInOrder())
	if inOrder != "2, 5, 10, 12, 15" {
		t.Errorf("Should printInOrder properly. Want \"2, 5, 10, 12, 15\", got \"%s\"", inOrder)
	}

	preOrder := parseKeys(myBST.PrintPreOrder())
	if preOrder != "5, 2, 10, 12, 15" {
		t.Errorf("Should printInOrder properly. Want \"2, 5, 10, 12, 15\", got \"%s\"", inOrder)
	}

	postOrder := parseKeys(myBST.PrintPostOrder())
	if postOrder != "2, 15, 12, 10, 5" {
		t.Errorf("Should printInOrder properly. Want \"2, 5, 10, 12, 15\", got \"%s\"", inOrder)
	}

	myBST.Insert(7, 14)
	myBST.Insert(6, 12)
	myBST.Insert(8, 16)
	removedNode, success := myBST.Remove(7)
	if !success {
		t.Error("Should be able to remove node")
	}
	if removedNode.value != 14 {
		t.Errorf("Removed node should contain correct value. Want 14, got %d", removedNode.value)
	}
	values := parseKeys(myBST.PrintInOrder())
	wantValues := "2, 5, 6, 8, 10, 12, 15"
	if values != wantValues {
		t.Errorf("Should remove proper node. Want %s, got %s", wantValues, values)
	}

	removedNode, success = myBST.Remove(999)
	if success {
		t.Error("Should not be able to remove node that does not exist")
	}
	if removedNode != nil {
		t.Errorf("Should return nil for removed node that does not exist. Want <nil>, got %+v", removedNode)
	}

}

func parseKeys(values string) string {
	removedReplaceAll := strings.ReplaceAll(values, "Key of ", "")
	r := regexp.MustCompile(": [0-9]{1,2}")
	return r.ReplaceAllString(removedReplaceAll, "")
}

// package ds

// import (
// 	"regexp"
// 	"strings"
// 	"testing"
// )

// func TestBinarySearchTree(t *testing.T) {
// 	myBST := BinarySearchTree{}
// 	if myBST.root != nil {
// 		t.Error("Should be empty when instantiated.")
// 	}
// 	success := myBST.Insert(5, 1)
// 	if !success {
// 		t.Error("Should be able to insert first node.")
// 	}
// 	if myBST.root == nil {
// 		t.Error("Should not have root == nill after insertion of first node.")
// 	}
// 	rootKey := myBST.root.key
// 	rootData := myBST.root.value
// 	if rootKey != 5 {
// 		t.Errorf("Should have proper key at root. Want 5, got %d", rootKey)
// 	}
// 	if rootData != 1 {
// 		t.Errorf("Should have proper data at root. Want 1, got %d", rootData)
// 	}

// 	success = myBST.Insert(5, 4)
// 	if success {
// 		t.Error("Should not be able to insert key that already exists.")
// 	}
// 	rootData = myBST.root.value
// 	if rootData != 1 {
// 		t.Errorf("Should not change data at root if key already existed. Want 1, got %d", rootData)
// 	}

// 	val, success := myBST.Search(5)
// 	if !success {
// 		t.Error("Should be able to find node that exists")
// 	}

// 	if val != 1 {
// 		t.Errorf("Should retrieve proper value for found node. Want 1, got %d", val)
// 	}

// 	val, success = myBST.Search(2)
// 	if success {
// 		t.Error("Should not be able to retrieve node that doesn't exist.")
// 	}
// 	if val != 0 {
// 		t.Errorf("Should return 0 value for node that doesn't exist. Want 0, got %d", val)
// 	}

// 	success = myBST.Insert(10, 4)
// 	if !success {
// 		t.Error("Should be able to insert new node")
// 	}
// 	rightKey := myBST.root.right.key
// 	rightVal := myBST.root.right.value
// 	if rightKey != 10 {
// 		t.Errorf("Should have inserted 10 to the right of root. Want 10, got %d", rightKey)
// 	}
// 	if rightVal != 4 {
// 		t.Errorf("Inserted node should have correct vaule. Want 4, got %d", rightVal)
// 	}

// 	myBST.Insert(2, 1)
// 	leftKey := myBST.root.left.key
// 	leftVal := myBST.root.left.value
// 	if leftKey != 2 {
// 		t.Errorf("Should have inserted 10 to the left of root. Want 10, got %d", leftKey)
// 	}
// 	if leftVal != 1 {
// 		t.Errorf("Inserted node should have correct vaule. Want 4, got %d", leftVal)
// 	}

// 	success = myBST.Insert(12, 1)
// 	if !success {
// 		t.Error("Should be able to insert new node")
// 	}

// 	myBST.Insert(15, 1)
// 	success = myBST.Insert(12, 4)
// 	if success {
// 		t.Error("Should not be able to insert duplicate key")
// 	}

// 	val, success = myBST.Search(12)
// 	if !success {
// 		t.Error("Should be able to find key that exists")
// 	}
// 	if val != 1 {
// 		t.Errorf("Should return proper value for node. Want 1, got %d", val)
// 	}

// 	val, success = myBST.Search(11)
// 	if success {
// 		t.Error("Should not be able to find key that does not exist")
// 	}
// 	if val != 0 {
// 		t.Errorf("Should return 0 value for key that does not exist. Want 0, got %d", val)
// 	}

// 	inOrder := parseKeys(myBST.PrintInOrder())
// 	if inOrder != "2, 5, 10, 12, 15" {
// 		t.Errorf("Should printInOrder properly. Want \"2, 5, 10, 12, 15\", got \"%s\"", inOrder)
// 	}

// 	preOrder := parseKeys(myBST.PrintPreOrder())
// 	if preOrder != "5, 2, 10, 12, 15" {
// 		t.Errorf("Should printInOrder properly. Want \"2, 5, 10, 12, 15\", got \"%s\"", inOrder)
// 	}

// 	postOrder := parseKeys(myBST.PrintPostOrder())
// 	if postOrder != "2, 15, 12, 10, 5" {
// 		t.Errorf("Should printInOrder properly. Want \"2, 5, 10, 12, 15\", got \"%s\"", inOrder)
// 	}

// 	myBST.Insert(7, 14)
// 	myBST.Insert(6, 12)
// 	myBST.Insert(8, 16)
// 	removedNode, success := myBST.Remove(7)
// 	if !success {
// 		t.Error("Should be able to remove node")
// 	}
// 	if removedNode.value != 14 {
// 		t.Errorf("Removed node should contain correct value. Want 14, got %d", removedNode.value)
// 	}
// 	values := parseKeys(myBST.PrintInOrder())
// 	wantValues := "2, 5, 6, 8, 10, 12, 15"
// 	if values != wantValues {
// 		t.Errorf("Should remove proper node. Want %s, got %s", wantValues, values)
// 	}

// 	removedNode, success = myBST.Remove(999)
// 	if success {
// 		t.Error("Should not be able to remove node that does not exist")
// 	}
// 	if removedNode != nil {
// 		t.Errorf("Should return nil for removed node that does not exist. Want <nil>, got %+v", removedNode)
// 	}

// }

// func parseKeys(values string) string {
// 	removedReplaceAll := strings.ReplaceAll(values, "Key of ", "")
// 	r := regexp.MustCompile(": [0-9]{1,2}")
// 	return r.ReplaceAllString(removedReplaceAll, "")
// }
