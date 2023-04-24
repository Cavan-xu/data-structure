package lru

import (
	"container/list"
	"sync"
)

type LRU struct {
	maxByte int
	curByte int
	list    *list.List
	cache   map[string]*list.Element
	mu      sync.RWMutex
}

func New(maxByte int) *LRU {
	return &LRU{
		maxByte: maxByte,
		curByte: 0,
		list:    list.New(),
		cache:   make(map[string]*list.Element),
	}
}

func (l *LRU) Get(key string) ([]byte, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	ele, ok := l.cache[key]
	if !ok {
		return nil, false
	}
	// 元素被访问，移动到链表最前面
	l.list.MoveToFront(ele)
	entry := ele.Value.(IEntry)
	return entry.Value(), true
}

func (l *LRU) Add(entry IEntry) {
	l.mu.Lock()
	defer l.mu.Unlock()

	defer func() {
		// 已使用字节数大于最大字节数时，需要移除链表尾部节点，直到已使用字节数小于最大字节数
		for l.maxByte > 0 && l.maxByte < l.curByte {
			l.removeOldest()
		}
	}()

	ele, ok := l.cache[entry.Key()]
	if !ok {
		l.cache[entry.Key()] = l.list.PushFront(entry)
		l.curByte += len(entry.Key()) + entry.Len()
		return
	}

	l.list.MoveToFront(ele)
	oldEntry := ele.Value.(IEntry)
	l.curByte += entry.Len() - oldEntry.Len()
	oldEntry.SetValue(entry.Value())
}

func (l *LRU) Len() int {
	return l.list.Len()
}

func (l *LRU) removeOldest() {
	ele := l.list.Back()
	if ele != nil {
		l.list.Remove(ele)
		entry := ele.Value.(IEntry)
		delete(l.cache, entry.Key())
		l.curByte -= len(entry.Key()) + entry.Len()
	}
}
