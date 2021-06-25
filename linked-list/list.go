package main

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Append(data int) {
	if list.head == nil {
		list.head = NewNode(data)
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = NewNode(data)
}

func (list *LinkedList) Prepend(data int) {
	if list.head == nil {
		list.head = NewNode(data)
		return
	}

	newHead := NewNode(data)
	newHead.next = list.head

	list.head = newHead
}

func (list *LinkedList) Delete(data int) {
	if list.head == nil {
		return
	}

	current := list.head

	for current.next != nil {
		// skip the next node if the data matches
		if current.next != nil && current.next.data == data {
			current.next = current.next.next
		}

		current = current.next
	}

	// remove head nodes which matches the data
	for list.head.data == data {
		list.head = list.head.next
	}
}