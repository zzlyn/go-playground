package list

// LinkedNode is the very basic node unit for linked lists
type LinkedNode struct {
	val  int
	next *LinkedNode
}

// DoubleLinkedNode is a linked list node with previous pointer
type DoubleLinkedNode struct {
	val      int
	previous *DoubleLinkedNode
	next     *DoubleLinkedNode
}
