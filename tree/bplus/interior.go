package bplustree

type kc struct {
	key   int
	child node
}

type kcs [LevelMaxNode]*kc

type InteriorNode struct {
	kcs   *kcs
	count int
	p     *InteriorNode
}

func newInteriorNode(p *InteriorNode, child node) *InteriorNode {
	i := &InteriorNode{
		p:     p,
		count: 1,
	}
	if child != nil {
		i.kcs[0].child = child
	}
	return i
}

func (i *InteriorNode) insert() {

}

func (i *InteriorNode) find() {

}

func (i *InteriorNode) split() {

}

func (i *InteriorNode) parent() *InteriorNode {
	return i.p
}

func (i *InteriorNode) setParent(n *InteriorNode) {
	i.p = n
}

func (i *InteriorNode) isFull() bool {
	return i.count == len(i.kcs)
}
