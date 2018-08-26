package main

import (
	"fmt"
	"go-playground/structure/list"
)

func singleListTest() {
	myList := list.NewSingleLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(myList)
	myList.Prepend(1)
	fmt.Println(myList)
	myList.Prepend(2)
	fmt.Println(myList)
	myList.Prepend(3)
	fmt.Println(myList)
	myList.Prepend(4)
	fmt.Println(myList)
	myList.RemoveFirst(5)
	fmt.Println(myList)
	myList.Append(5)
	fmt.Println(myList)
	myList.RemoveFirst(1)
	myList.RemoveFirst(1)
	fmt.Println(myList)
	myList.RemoveFirst(2)
	fmt.Println(myList)
	myList.RemoveFirst(2)
	myList.RemoveFirst(3)
	myList.RemoveFirst(3)
	myList.RemoveFirst(5)
	fmt.Println(myList)
	myList.RemoveFirst(4)
	myList.RemoveFirst(4)
	fmt.Println(myList)
	myList.RemoveFirst(4)

	myList2 := list.NewSingleLinkedList([]int{4, 5, 6, 4, 4, 4, 4, 4, 4, 7})
	fmt.Println(myList2)
	myList2.RemoveAllLazy(4)
	fmt.Println(myList2)

	myList3 := list.NewSingleLinkedList([]int{1, 4, 5, 1, 6, 4, 4, 4, 1, 1, 4, 4, 4, 7, 1})
	fmt.Println(myList3)
	myList3.RemoveAll(1)
	fmt.Println(myList3)
}

func main() {
	singleListTest()
}
