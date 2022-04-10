package ds

// arraySize is the size of the hash table array
const arraySize = 100

type bucketNode[T any] struct {
	key   string
	value T
	next  *bucketNode[T]
}

type bucket[T any] struct {
	head *bucketNode[T]
}

type hashTable[T any] [arraySize]*bucket[T]

func (b *bucket[T]) insert(key string, value T) bool {
	if _, success := b.search(key); success {
		return false
	}
	newNode := &bucketNode[T]{key: key, value: value}
	newNode.next = b.head
	b.head = newNode
	return true
}

func (b *bucket[T]) search(key string) (T, bool) {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return currentNode.value, true
		}
		currentNode = currentNode.next
	}
	var zeroValue T
	return zeroValue, false
}

func (b *bucket[T]) delete(key string) (T, bool) {
	var zeroValue T
	if b.head == nil {
		return zeroValue, false
	}

	if b.head.key == key {
		returnVal := b.head.value
		b.head = b.head.next
		return returnVal, true
	}

	prev := b.head
	for prev.next != nil {
		if prev.next.key == key {
			returnVal := prev.next.value
			prev.next = prev.next.next
			return returnVal, true
		}
	}
	return zeroValue, false
}

func NewHashTable[T comparable]() *hashTable[T] {
	result := &hashTable[T]{}
	for i := range result {
		result[i] = &bucket[T]{}
	}
	return result
}

func (h *hashTable[T]) Insert(key string, value T) bool {
	index := hash(key)
	return h[index].insert(key, value)
}

func (h *hashTable[T]) Search(key string) (T, bool) {
	index := hash(key)
	return h[index].search(key)
}

func (h *hashTable[T]) Delete(key string) (T, bool) {
	index := hash(key)
	return h[index].delete(key)
}

func hash[K string](key K) int {
	var sum int
	for _, char := range key {
		sum += int(char)
	}
	return sum % arraySize
}
