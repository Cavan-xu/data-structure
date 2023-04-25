package bplustree

const (
	LevelMaxNode = 512
)

type node interface {
	parent() *InteriorNode
	setParent(node *InteriorNode)
	isFull() bool
}
