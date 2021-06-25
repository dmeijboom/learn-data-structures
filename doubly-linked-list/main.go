package main

import "fmt"

func printList(list *LinkedList) {
	fmt.Println("list()")

	current := list.head

	for {
		fmt.Printf("  node(%d)\n", current.data)

		if current.next == nil {
			fmt.Println("list.reversed()")

			for current != nil {
				fmt.Printf("  node(%d)\n", current.data)

				current = current.prev
			}

			break
		}

		current = current.next
	}
}

func main() {
	list := NewLinkedList()
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	printList(list)

	list.Prepend(1)

	printList(list)

	list.Delete(3)
	list.Delete(1)

	printList(list)
}
