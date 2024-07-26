package structures

import "fmt"

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
}

func (l *LinkedList) Insert(value int) {
	node := &node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		head := l.head

		for head.next != nil {
			head = head.next
		}

		head.next = node
	}
}

func (l *LinkedList) Show() {
	head := l.head

	for head != nil {
		fmt.Printf("-> %v ", head.value)
		head = head.next
	}
}
