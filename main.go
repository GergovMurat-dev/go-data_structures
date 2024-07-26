package main

import (
	"data_structures/structures"
	"fmt"
)

func main() {
	linkedList := structures.LinkedList{}

	linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.Insert(4)
	linkedList.Insert(5)
	linkedList.Insert(6)

	current := linkedList.First

	for current != nil {
		fmt.Printf("-> %v ", current.Value)
		current = current.Next
	}
}
