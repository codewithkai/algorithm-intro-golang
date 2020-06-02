package hashmap

type ListNode struct {
	key  string
	val  string
	next *ListNode
}

type HashMap struct {
	table []*ListNode
}

func NewHashMap(capacity int) *HashMap {
	return &HashMap{make([]*ListNode, capacity)}
}

const MULTIPLIER = 31

func (hm *HashMap) hash(key string) int {
	h := 0
	for _, v := range key {
		h = MULTIPLIER*h + int(v)
	}
	return h % len(hm.table)
}

func (hm *HashMap) Get(key string) *ListNode {
	h := hm.hash(key)
	for node := hm.table[h]; node != nil; node = node.next {
		if node.key == key {
			return node
		}
	}
	return nil
}

func (hm *HashMap) Set(key string, val string) {
	h := hm.hash(key)
	for node := hm.table[h]; node != nil; node = node.next {
		if node.key == key {
			node.val = val
			return
		}
	}
	newNode := &ListNode{key: key, val: val, next: hm.table[h]}
	hm.table[h] = newNode
}
