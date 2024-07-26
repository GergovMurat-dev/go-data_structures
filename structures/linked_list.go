package structures

/*
Связанный список - структура данных, состоящий из набора элементов
называющимися узлами, каждый из которых содержит данные и ссылку
на следующий узел в списке

Таким образом элементы связанного списка расположены произвольно
в памяти и соединены ссылками, что позволяет добавлять и удалять
элементы в любом месте списка без перераспределения памяти

*/

type node struct {
	Value any
	Next  *node
	Prev  *node
}

type LinkedList struct {
	First *node
	last  *node
}

func (l *LinkedList) Insert(value any) {
	node := &node{Value: value}
	if l.First == nil {
		l.First = node
	} else {
		head := l.First

		for head.Next != nil {
			head = head.Next
		}
		node.Prev = head
		head.Next = node
		l.last = node
	}
}
