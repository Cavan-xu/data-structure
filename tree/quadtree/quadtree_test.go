package quadtree

import (
	"testing"
)

type Monster struct {
	Id int
	X  float64
	Y  float64
}

func (m *Monster) GetKey() int {
	return m.Id
}

func (m *Monster) GetX() float64 {
	return m.X
}

func (m *Monster) GetY() float64 {
	return m.Y
}

func TestQuadTree(t *testing.T) {
	root := NewTreeNode(0, 0, 100, 0)
	entityList := []*Monster{
		{1, 1, 1},
		{2, 10, 10},
		{3, 5, 90},
		{4, 70, 5},
		{5, 99, 99},
		{6, 80, 80},
	}
	for _, entity := range entityList {
		root.Add(entity)
	}
}
