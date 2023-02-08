package linear

import "sync"

/*
	单链表
 */

func NewSingleListNode(val int) *SingleListNode {
	return &SingleListNode{
		Val: val,
	}
}

type SingleListNode struct {
	Val int
	Next *SingleListNode
}

func (node *SingleListNode) SetNext(n *SingleListNode) {
	node.Next = n
}

type SingleList struct {
	sync.RWMutex
	tail *SingleListNode
}

func (lst *SingleList) Add(n *SingleListNode) {
	lst.Lock()
	defer lst.Unlock()

	node := lst.tail
	for node.Next != nil {
		node = node.Next
	}

	node.SetNext(n)
}