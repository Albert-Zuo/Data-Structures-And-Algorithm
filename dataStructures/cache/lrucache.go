package cache

// Element node to store cache item
type Element struct {
	prev, next *Element
	Key        interface{}
	Value      interface{}
}

// Prev get prev node for current node
func (e *Element) Prev() *Element {
	return e.prev
}

// Next get next node for current node
func (e *Element) Next() *Element {
	return e.next
}

// LRUCache  a data structure that is efficient to insert/fetch/delete cache items [both O(1) time complexity]
type LRUCache struct {
	cache    map[interface{}]*Element
	head     *Element
	tail     *Element
	capacity int
}

// New create lru cache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{make(map[interface{}]*Element), nil, nil, capacity}
}

// Front - get front element of lru cache
func (lc *LRUCache) Front() *Element {
	return lc.head
}

// Back - get back element of lru cache
func (lc *LRUCache) Back() *Element {
	return lc.tail
}

// Get find element from lru cache with key
func (lc *LRUCache) Get(key interface{}) (interface{}, bool) {
	if e, ok := lc.cache[key]; ok {
		lc.up(e)
		return e.Value, ok
	}
	return nil, false
}

// ForEach traverses the elements in the lru cache in order
// f - element operation function, end traversal when false is returned
func (lc *LRUCache) ForEach(f func(key, value interface{}) bool) {
	for node := lc.head; node != nil; node = node.next {
		if !f(node.Key, node.Value) {
			break
		}
	}
}

// Add add cache element into lru cache
func (lc *LRUCache) Add(key, value interface{}) {
	// key is exist
	if e, ok := lc.cache[key]; ok {
		e.Value = value
		lc.up(e)
		return
	}

	if lc.capacity == 0 {
		return
	} else if len(lc.cache) >= lc.capacity {
		// evict the oldest item
		delete(lc.cache, lc.tail.Key)
		lc.remove(lc.tail)
	}

	e := &Element{nil, lc.head, key, value}
	lc.cache[key] = e
	if len(lc.cache) != 1 {
		lc.head.prev = e
	} else {
		lc.tail = e
	}
	lc.head = e
}

// Remove delete element from lru cache with key
func (lc *LRUCache) Remove(key interface{}) {
	if e, ok := lc.cache[key]; ok {
		delete(lc.cache, e)
		lc.remove(e)
	}
}

// Update
func (lc *LRUCache) Update(key interface{}, f func(value *interface{})) {
	if e, ok := lc.cache[key]; ok {
		f(&e.Value)
		lc.up(e)
	}
}

// Len - length of lru cache
func (lc *LRUCache) Len() int {
	return len(lc.cache)
}

// Capacity - capacity of lru cache
func (lc *LRUCache) Capacity() int {
	return lc.capacity
}

// up Float up to the linked head
func (lc *LRUCache) up(e *Element) {
	if e.prev != nil {
		e.prev.next = e.next
		if e.next == nil {
			lc.tail = e.prev
		} else {
			e.next.prev = e.prev
		}
		e.prev = nil
		e.next = lc.head
		lc.head.prev = e
		lc.head = e
	}
}

func (lc *LRUCache) remove(e *Element) {
	if e.prev == nil {
		lc.head = e.next
	} else {
		e.prev.next = e.next
	}
	if e.next == nil {
		lc.tail = e.prev
	} else {
		e.next.prev = e.prev
	}
}
