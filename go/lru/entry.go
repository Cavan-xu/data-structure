package lru

type IEntry interface {
	Key() string
	Value() []byte
	SetValue([]byte)
	Len() int
}

type Entry struct {
	key   string
	value []byte
}

func (e *Entry) Key() string {
	return e.key
}

func (e *Entry) Value() []byte {
	return e.value
}

func (e *Entry) SetValue(v []byte) {
	e.value = v
}

func (e *Entry) Len() int {
	return len(e.value)
}
