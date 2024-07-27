package main

import (
	"data_structures/structures"
	"fmt"
)

func linkedList() {
	linkedList := structures.LinkedList[int]{}

	linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.Insert(4)
	linkedList.Insert(5)
	linkedList.Insert(6)

	// current := linkedList.First

	// for current != nil {
	// 	fmt.Printf("-> %v ", current.Value)
	// 	current = current.Next
	// }
	// fmt.Println("")
}

func stack() {
	stack := structures.Stack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// fmt.Print(stack.Pop())
}

func set() {
	set := structures.Set[int]{}
	set2 := structures.Set[int]{}

	set.Add(1)
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(4)

	set2.Add(6)
	set2.Add(8)
	set2.Add(10)

	set.Remove(2)

	set3 := set.Union(set2)

	for _, item := range set3.Values() {
		fmt.Println(item)
	}
}

func main() {
	fmt.Println("-------------------LinkedList------------------")

	linkedList()

	fmt.Println("--------------------STACK----------------------")

	stack()

	fmt.Println("--------------------SET----------------------")

	set()
}
