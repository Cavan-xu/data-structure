package lru

type IEntry interface {
	GetKey() string
	GetValue() []byte
	SetValue([]byte)
	Len() int
}

type Entry struct {
	Key   string
	Value []byte
}

func (e *Entry) GetKey() string {
	return e.Key
}

func (e *Entry) GetValue() []byte {
	return e.Value
}

func (e *Entry) SetValue(v []byte) {
	e.Value = v
}

func (e *Entry) Len() int {
	return len(e.Value)
}
