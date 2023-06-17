package main

import "fmt"

// Defining a Node
type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d\n", current.data)
		current = current.next
	}
	fmt.Println()
}
func (ll *LinkedList) Find(val int) int {
	current := ll.head
	index := 0
	for current != nil {
		if current.data == val {
			return index
		}
		index++
		current = current.next
	}
	return -1
}
func (ll *LinkedList) Length() int {
	current := ll.head
	counter := 0
	for current != nil {
		counter++
		current = current.next
	}
	return counter
}
func (ll *LinkedList) GetAt(index int) *Node {
	if ll.Length() < index {
		return nil
	}
	counter := 0
	current := ll.head
	for current.next != nil {
		if counter == index {
			return current
		}
		current = current.next
		counter++
	}
	return nil
}
func (ll *LinkedList) InsertAt(index int, val int) {
	// # BASIC CHECKS
	previous := ll.GetAt(index - 1)
	previous.next = &Node{
		data: val,
		next: previous.next,
	}

}
func main() {
	ll := LinkedList{}
	ll.Insert(5)
	ll.Insert(10)
	ll.Insert(15)
	ll.Insert(20)
	ll.Display()
	fmt.Println("------")
	// fmt.Printf("%d\n", ll.Find(200))
	// fmt.Printf("%d\n", ll.GetAt(21).data)
	ll.InsertAt(2, 99)
	ll.Display()
}
