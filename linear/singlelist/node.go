package singlelist

type INode interface {
	Key() int
	Value() interface{}
	Next() INode
	SetNext(INode)
}

func NewNode(key int, value interface{}) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

type Node struct {
	key   int
	value interface{}
	next  *Node
}

func (node *Node) Key() int {
	return node.key
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) Next() INode {
	if node.next == nil {
		return nil
	}
	return node.next
}

func (node *Node) SetNext(n INode) {
	next, ok := n.(*Node)
	if ok {
		node.next = next
	}
}
