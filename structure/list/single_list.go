package list

import (
	"fmt"
	"strconv"
)

// SingleLinkedList is a wrapper around list formed by singly linked nodes
type SingleLinkedList struct {
	head   *LinkedNode
	length int
}

// NewSingleLinkedList is the typical constructor for integer NewSingleLinkedLists
func NewSingleLinkedList(data []int) *SingleLinkedList {
	prev := &LinkedNode{val: 0, next: nil}
	temp := prev
	for _, val := range data {
		temp.next = &LinkedNode{val: val, next: nil}
		temp = temp.next
	}

	return &SingleLinkedList{head: prev.next, length: len(data)}
}

func (list *SingleLinkedList) String() string {
	var temp *LinkedNode
	temp = list.head

	var result string
	fmt.Println("Length: " + strconv.Itoa(list.length))
	for temp != nil {
		result += strconv.Itoa(temp.val) + "->"
		temp = temp.next
	}
	result += "nil\n"
	return result
}

// RemoveFirst removes the first occurence of target in list
func (list *SingleLinkedList) RemoveFirst(target int) bool {
	if list.head == nil {
		return false
	}

	temp := list.head
	if temp.val == target {
		list.head = temp.next
		list.length--
		return true
	}
	for temp.next != nil {
		if temp.next.val != target {
			temp = temp.next
			continue
		}

		temp.next = temp.next.next
		list.length--
		return true
	}
	return false
}

// Prepend adds a value to the front of the list
func (list *SingleLinkedList) Prepend(target int) {
	list.length++
	prev := LinkedNode{val: target, next: list.head}
	list.head = &prev
}

// Append adds a value to the back of the list
func (list *SingleLinkedList) Append(target int) {
	list.length++
	post := &LinkedNode{val: target, next: nil}

	if list.head == nil {
		list.head = post
		return
	}

	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = post
}
