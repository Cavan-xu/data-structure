package bplustree

type kv struct {
	key   int
	value string
}

type kvs [LevelMaxNode]*kv

type LeafNode struct {
	kvs   *kvs
	count int
	p     *InteriorNode
}

func NewLeafNode(p *InteriorNode) *LeafNode {
	return &LeafNode{
		p: p,
	}
}

func (l *LeafNode) find() {

}

func (l *LeafNode) insert() {

}

func (l *LeafNode) split() {

}

func (l *LeafNode) parent() *InteriorNode {
	return l.p
}

func (l *LeafNode) setParent(n *InteriorNode) {
	l.p = n
}

func (l *LeafNode) isFull() bool {
	return l.count == len(l.kvs)
}
