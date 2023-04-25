package bplustree

/*
	b+树
*/

type Tree struct {
	root  *InteriorNode // 根
	first *LeafNode     // 第一个叶子
}

func NewTree() *Tree {
	return &Tree{
		root: newInteriorNode(nil, nil),
	}
}

func (t *Tree) Insert() {

}

func (t *Tree) Search() {

}

func (t *Tree) First() {

}
