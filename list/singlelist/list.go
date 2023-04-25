package singlelist

import (
	"math"
)

/*
	单链表
*/

type List struct {
	tail  INode
	count int
}

func NewList() *List {
	return &List{
		tail: NewNode(math.MaxInt32, nil),
	}
}

func (lst *List) Add(n INode) {
	node := lst.tail
	for node.Next() != nil {
		node = node.Next()
	}
	node.SetNext(n)
	lst.count++
}

func (lst *List) Get(key int) INode {
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
	nodeList := make([]INode, 0, lst.count)
	node := lst.tail
	for node.Next() != nil {
		nodeList = append(nodeList, node.Next())
		node = node.Next()
	}
	return nodeList
}

func (lst *List) Reverse() {
	var pre INode = nil
	node := lst.tail.Next()
	for node != nil {
		next := node.Next()
		node.SetNext(pre)
		pre = node
		node = next
	}
	lst.tail.SetNext(pre)
}
