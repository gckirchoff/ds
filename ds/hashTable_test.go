package ds

import "testing"

const appleAmount = 23
const mangoAmount = 55

func TestHashTable(t *testing.T) {
	myHashTable := NewHashTable[int]()

	success := myHashTable.Insert("apple", appleAmount)
	if !success {
		t.Error("Should be able to insert key")
	}

	success = myHashTable.Insert("apple", 25)
	if success {
		t.Error("Should not be able to insert duplicate key")
	}

	success = myHashTable.Insert("mango", mangoAmount)
	if !success {
		t.Error("Should be able to insert of same hash value")
	}

	val, success := myHashTable.Search("apple")
	if !success {
		t.Error("Should be able to successfully find inserted key")
	}

	if val != appleAmount {
		t.Errorf("Should return correct value for returned key. Want: %d, got: %d", appleAmount, val)
	}

	val, success = myHashTable.Search("mango")
	if !success {
		t.Error("Should be able to successfully find inserted key")
	}

	if val != mangoAmount {
		t.Errorf("Should return correct value for returned key. Want: %d, got: %d", mangoAmount, val)
	}

	val, success = myHashTable.Delete("apple")
	if !success {
		t.Error("Should be able to delete key in hash table")
	}

	if val != appleAmount {
		t.Errorf("Should be able to return correct value of deleted node. Want: %d, got: %d", appleAmount, val)
	}

	val, success = myHashTable.Delete("banana")
	if success {
		t.Error("Should not be able to delete key that does not exist")
	}

	if val != 0 {
		t.Errorf("Should return 0 value for unsuccessful deletion. Want: 0, got: %d", val)
	}

	val, success = myHashTable.Delete("mango")
	if !success {
		t.Error("Should be able to delete key in hash table")
	}

	if val != mangoAmount {
		t.Errorf("Should be able to return correct value of deleted node. Want: %d, got: %d", mangoAmount, val)
	}
}
