package structures

/*
Связанный список - структура данных, состоящий из набора элементов
называющимися узлами, каждый из которых содержит данные и ссылку
на следующий узел в списке

Таким образом элементы связанного списка расположены произвольно
в памяти и соединены ссылками, что позволяет добавлять и удалять
элементы в любом месте списка без перераспределения памяти

*/

type node[T any] struct {
	Value T
	Next  *node[T]
	Prev  *node[T]
}

type LinkedList[T any] struct {
	First *node[T]
	last  *node[T]
}

func (l *LinkedList[T]) Insert(value T) {
	node := &node[T]{Value: value}
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
