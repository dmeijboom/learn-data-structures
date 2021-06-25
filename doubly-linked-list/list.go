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

	node := NewNode(data)

	current.next = node
	node.prev = current
}

func (list *LinkedList) Prepend(data int) {
	if list.head == nil {
		list.head = NewNode(data)
		return
	}

	newHead := NewNode(data)
	newHead.next = list.head

	list.head.prev = newHead
	list.head = newHead
}

func (list *LinkedList) Delete(data int) {
	if list.head == nil {
		return
	}

	current := list.head

	for current.next != nil {
		if current.data == data {
			if current.prev != nil {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
		}

		current = current.next
	}

	if list.head.data == data {
		list.head = list.head.next
		list.head.prev = nil
	}
}
