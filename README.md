![logo](./logo.png)

## ds

ds (short for data structures) is an easy to use, light weight library for implementing various generic data structures such as BSTs, tries, and linked lists in Golang.

`import "github.com/gckirchoff/ds/ds"`

### Current data structures and methods:
* Binary Search Tree
    * `Insert(key int, data T) bool`
    * `Search(key int) (T, bool)`
    * `Remove(key int) (*treeNode[T], bool)`
    * `PrintInOrder() string`
    * `PrintPreOrder() string`
    * `PrintPostOrder() string`
* Hash Table
    * `Insert(key string, value T) bool`
    * `Search(key string) (T, bool)`
    * `Delete(key string) (T, bool)`
* Heap
    * `Insert(key T)`
    * `Extract() (T, bool)`
* Linked List
    * `Length() int`
    * `Prepend(value T)`
    * `Append(value T)`
    * `Delete(val T) (T, bool)`
    * `Print() string`
* Queue
    * `IsEmpty() bool`
    * `Enqueue(n T)`
    * `Dequeue() (dequeued T, success bool)`
* Stack
    * `IsEmpty() bool`
    * `Push(n T)`
    * `Pop() (popped T, success bool)`
    * `Peek() (T, bool)`
* Trie
    * `NewTrie() *trie`
    * `Insert(word string) bool`
    * `Search(word string) bool`
    * `WordOptions(word string) []string`