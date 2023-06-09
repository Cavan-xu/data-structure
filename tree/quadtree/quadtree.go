package quadtree

const (
	maxChild    = 4 // 最大子树个数
	maxCapacity = 5 // 每个子树最大容量
	maxDepth    = 5 // 四叉树最大深度
)

type azimuth int

const (
	leftUp azimuth = iota
	rightUp
	leftDown
	rightDown
)

type Entity interface {
	GetKey() int
	GetX() float64
	GetY() float64
}

func NewCurrentArea(xStart, yStart, width float64) CurrentArea {
	return CurrentArea{xStart, yStart, width}
}

type CurrentArea struct {
	XStart    float64
	YStart    float64
	AreaWidth float64
}

func (c *CurrentArea) Includes(x, y float64) bool {
	return c.XStart < x && (c.XStart+c.AreaWidth) > x && c.YStart < y && (c.YStart+c.AreaWidth) > y
}

func (c *CurrentArea) CanCut() bool {
	return (c.XStart+c.AreaWidth)/2 > 0 && (c.YStart+c.AreaWidth)/2 > 0
}

func (c *CurrentArea) Cut() [maxChild]CurrentArea {
	var child [maxChild]CurrentArea
	width := c.AreaWidth / 2
	child[leftUp] = CurrentArea{
		XStart:    c.XStart,
		YStart:    c.YStart,
		AreaWidth: width,
	}
	child[rightUp] = CurrentArea{
		XStart:    c.XStart + width,
		YStart:    c.YStart,
		AreaWidth: width,
	}
	child[leftDown] = CurrentArea{
		XStart:    c.XStart,
		YStart:    c.YStart + width,
		AreaWidth: width,
	}
	child[rightDown] = CurrentArea{
		XStart:    c.XStart + width,
		YStart:    c.YStart + width,
		AreaWidth: width,
	}
	return child
}

func NewData() *Data {
	return &Data{EntityList: make([]Entity, 0, maxCapacity)}
}

type Data struct {
	EntityList []Entity
}

func (d *Data) Add(entity Entity) {
	d.EntityList = append(d.EntityList, entity)
}

func (d *Data) Get(key int) (Entity, bool) {
	for _, entity := range d.EntityList {
		if entity.GetKey() == key {
			return entity, true
		}
	}
	return nil, false
}

func (d *Data) Delete(key int) {
	for i, entity := range d.EntityList {
		if entity.GetKey() == key {
			d.EntityList = append(d.EntityList[:i], d.EntityList[i+1:]...)
			return
		}
	}
}

func (d *Data) Replace(prev, cur Entity) {
	for i, entity := range d.EntityList {
		if entity.GetKey() == prev.GetKey() {
			d.EntityList[i] = cur
			return
		}
	}
}

func (d *Data) Clear() {
	d.EntityList = make([]Entity, 0, maxCapacity)
}

func (d *Data) Range(f func(entity Entity) bool) {
	for _, entity := range d.EntityList {
		if !f(entity) {
			return
		}
	}
}

func NewTreeNode(xStart, yStart, width float64, deep int) *Node {
	return &Node{
		child:       [maxChild]*Node{},
		data:        NewData(),
		deep:        deep,
		isLeaf:      true,
		CurrentArea: NewCurrentArea(xStart, yStart, width),
	}
}

type Node struct {
	child  [maxChild]*Node
	data   *Data
	deep   int
	isLeaf bool
	CurrentArea
}

func (n *Node) Add(entity Entity) {
	if n.isLeafNode() && n.needGrew() {
		n.grewTree()
	}
	if n.isLeafNode() {
		n.data.Add(entity)
		return
	}
	n.child[n.findIndex(entity.GetX(), entity.GetY())].Add(entity)
}

func (n *Node) Delete(entity Entity) {
	if n.isLeafNode() {
		n.data.Delete(entity.GetKey())
		return
	}
	n.child[n.findIndex(entity.GetX(), entity.GetY())].Delete(entity)
}

func (n *Node) Update(prev, cur Entity) {
	if n.isLeafNode() {
		_, ok := n.data.Get(prev.GetKey())
		if ok {
			if n.Includes(cur.GetX(), cur.GetY()) {
				n.data.Replace(prev, cur)
				return
			}
		}
		n.data.Delete(prev.GetKey())
		n.data.Add(cur)
	}
	n.child[n.findIndex(prev.GetY(), prev.GetY())].Update(prev, cur)
}

func (n *Node) Search(entity Entity, r float64, result *[]Entity) {
	if n.isLeafNode() {
		n.search(result)
		return
	}

	x, y := entity.GetX(), entity.GetY()
	xMin, xMax := x-r, x+r
	yMin, yMax := y-r, y+r
	for i := 0; i < maxChild; i++ {
		if n.child[i].Includes(xMin, yMin) || n.child[i].Includes(xMin, yMax) ||
			n.child[i].Includes(xMax, yMin) || n.child[i].Includes(xMax, yMax) {
			n.child[i].Search(entity, r, result)
		}
	}
}

func (n *Node) isLeafNode() bool {
	return n.isLeaf
}

func (n *Node) needGrew() bool {
	return n.getCapacity() == maxCapacity && n.getDeep() < maxDepth && n.CanCut()
}

func (n *Node) getCapacity() int {
	return len(n.data.EntityList)
}

func (n *Node) getDeep() int {
	return n.deep
}

func (n *Node) grewTree() {
	n.isLeaf = false
	newAreaList := n.Cut()
	for i, area := range newAreaList {
		n.child[i] = NewTreeNode(area.XStart, area.YStart, area.AreaWidth, n.deep+1)
		n.data.Range(func(entity Entity) bool {
			if n.child[i].Includes(entity.GetX(), entity.GetY()) {
				n.child[i].data.Add(entity)
			}
			return true
		})
	}
	n.data.Clear()
}

func (n *Node) findIndex(x, y float64) azimuth {
	// 因为每个格子的左上角都是起始点
	// 那么 右下方的格子的左上角则是这个格子的中心
	// 那么就根据右下方个字的左上角去比较到底这个坐标在哪里
	// 这样只需要比两次就好
	if x < n.child[rightDown].XStart {
		if y < n.child[rightDown].YStart {
			return leftUp
		}
		return leftDown
	}
	if y < n.child[rightDown].YStart {
		return rightUp
	}
	return rightDown
}

func (n *Node) search(result *[]Entity) {
	n.data.Range(func(entity Entity) bool {
		*result = append(*result, entity)
		return true
	})
}
