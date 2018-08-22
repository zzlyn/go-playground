package list

// LinkedNode is the very basic node unit for linked lists
type LinkedNode struct {
	value interface{}
	next  *LinkedNode
}

// DoubleLinkedNode is a linked list node with previous pointer
type DoubleLinkedNode struct {
	value    interface{}
	previous *DoubleLinkedNode
	next     *DoubleLinkedNode
}
