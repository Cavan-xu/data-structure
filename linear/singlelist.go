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

func (lst *SingleList) Get(val int) *SingleListNode {
	lst.RLock()
	defer lst.RUnlock()

	node := lst.tail
	for node != nil {
		if node.Val == val {
			return node
		}
		node = node.Next
	}

	return nil
}

func (lst *SingleList) Remove(val int) {
	lst.Lock()
	defer lst.Unlock()

	node := lst.tail
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			return
		}
		node = node.Next
	}
}