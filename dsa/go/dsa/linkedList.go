package dsa

import "fmt"

type SinglyLinkedList struct {
	head *SLLNode
	tail *SLLNode
}

func (sll *SinglyLinkedList) append(node *SLLNode) {
	sll.tail.append(node)
	sll.tail = node
}

// singly linked list Node
type SLLNode struct {
	value any
	next  *SLLNode
}

func (sn *SLLNode) append(node *SLLNode) {
	sn.next = node
}

func (sll *SinglyLinkedList) toSlice() []any {
	content := []any{sll.head.value}
	nextNode := sll.head.next
	for {
		if nextNode == nil {
			break
		}
		content = append(content, nextNode.value)
		nextNode = nextNode.next
	}

	return content
}

func InitSLL() {
	head := &SLLNode{value: 1}
	singlyLinkedList := SinglyLinkedList{
		head: head,
		tail: head,
	}

	n2 := &SLLNode{value: 2}
	n3 := &SLLNode{value: 3}
	n4 := &SLLNode{value: 4}

	singlyLinkedList.append(n2)
	singlyLinkedList.append(n3)
	singlyLinkedList.append(n4)

	fmt.Println(singlyLinkedList.toSlice())
}
