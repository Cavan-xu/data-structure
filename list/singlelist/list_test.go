package singlelist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SingleList(t *testing.T) {
	list := NewList()
	node1 := NewNode(1, nil)
	list.Add(node1)

	get := list.Get(node1.Key())
	getNode, ok := get.(*Node)
	assert.Equal(t, true, ok)
	assert.Equal(t, getNode, node1)

	node2 := NewNode(2, nil)
	node3 := NewNode(3, nil)
	node4 := NewNode(4, nil)
	node5 := NewNode(5, nil)
	node6 := NewNode(6, nil)
	list.Add(node2)
	list.Add(node3)
	list.Add(node4)
	list.Add(node5)
	list.Add(node6)
	list.Remove(node2.Key())
	nodeList := list.Travel()
	assert.Equal(t, 5, len(nodeList))
	for _, iNode := range nodeList {
		t.Logf("key: %d\n", iNode.Key())
	}

	list.Reverse()
	nodeList = list.Travel()
	assert.Equal(t, 5, len(nodeList))
	for _, iNode := range nodeList {
		t.Logf("key: %d\n", iNode.Key())
	}
}
