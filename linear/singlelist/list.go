package singlelist

import (
	"math"
	"sync"
)

/*
	单链表
*/

type List struct {
	sync.RWMutex
	tail  INode
	count int
}

func NewList() *List {
	return &List{
		tail: NewNode(math.MaxInt32, nil),
	}
}

func (lst *List) Add(n INode) {
	lst.Lock()
	defer lst.Unlock()

	node := lst.tail
	for node.Next() != nil {
		node = node.Next()
	}
	node.SetNext(n)
	lst.count++
}

func (lst *List) Get(key int) INode {
	lst.RLock()
	defer lst.RUnlock()

	node := lst.tail
	for node != nil {
		if node.Key() == key {
			return node
		}
		node = node.Next()
	}
	return nil
}

func (lst *List) Remove(key int) {
	lst.Lock()
	defer lst.Unlock()

	node := lst.tail
	for node.Next() != nil {
		if node.Next().Key() == key {
			lst.count--
			node.SetNext(node.Next().Next())
			return
		}
		node = node.Next()
	}
}

func (lst *List) Travel() []INode {
	lst.RLock()
	defer lst.RUnlock()

	nodeList := make([]INode, 0, lst.count)
	node := lst.tail
	for node.Next() != nil {
		nodeList = append(nodeList, node.Next())
		node = node.Next()
	}
	return nodeList
}
