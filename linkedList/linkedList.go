package main

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Append(value interface{}) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList) Prepend(value interface{}) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *LinkedList) Remove(node *Node) {
	if node == l.head {
		l.head = l.head.next
	}
	if node == l.tail {
		l.tail = l.tail.next
	}
	if node.next != nil {
		node.next = node.next.next
	}
}

func (l *LinkedList) Contains(value interface{}) bool {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (l *LinkedList) Len() int {
	count := 0
	currentNode := l.head
	for currentNode != nil {
		count++
		currentNode = currentNode.next
	}
	return count
}

func (l *LinkedList) ToSlice() []interface{} {
	slice := make([]interface{}, l.Len())
	currentNode := l.head
	for currentNode != nil {
		slice = append(slice, currentNode.value)
		currentNode = currentNode.next
	}
	return slice
}

func main() {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Prepend(0)
	fmt.Println(list.Contains(2)) // true
	fmt.Println(list.Len())       // 4
	fmt.Println(list.ToSlice())   // [0, 1, 2, 3]

}
