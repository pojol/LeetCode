package lfu

import "container/list"

// LFUCache  Least Frequently Used
type LFUCache struct {
	cap      int
	size     int
	mapCache map[int]*cacheEntry
	freqList *list.List
}

type cacheEntry struct {
	key        int
	value      int
	freqParent *list.Element // 指向访问链表根节点
}

type cacheFreq struct {
	freq    int
	freqMap map[*cacheEntry]byte
}

// Constructor 构造
func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:      capacity,
		size:     0,
		mapCache: make(map[int]*cacheEntry),
		freqList: list.New(),
	}
}

// Get 获取键值对应的节点
func (lfu *LFUCache) Get(key int) int {

	if lfu.cap == 0 {
		return -1
	}

	if e, ok := lfu.mapCache[key]; ok {
		lfu.increment(e)
		return e.value
	}

	return -1
}

// Put 添加节点数据
func (lfu *LFUCache) Put(key int, value int) {

	if lfu.cap == 0 {
		return
	}

	if e, ok := lfu.mapCache[key]; ok {
		e.value = value
		lfu.increment(e)
	} else {
		entry := &cacheEntry{
			key:   key,
			value: value,
		}
		lfu.mapCache[key] = entry
		lfu.size++
		// evicts & increment
		if lfu.size > lfu.cap {
			lfu.evicts()
		}

		lfu.increment(entry)
	}

}

// 将节点进行右移
func (lfu *LFUCache) increment(e *cacheEntry) {
	curNode := e.freqParent
	var nextFreq int
	var nextNode *list.Element

	if curNode == nil { // new node
		nextFreq = 1
		nextNode = lfu.freqList.Front() // head
	} else {
		nextFreq = curNode.Value.(*cacheFreq).freq + 1
		nextNode = curNode.Next()
	}

	if nextNode == nil || nextNode.Value.(*cacheFreq).freq != nextFreq {
		// create a new list entry
		li := new(cacheFreq)
		li.freq = nextFreq
		li.freqMap = make(map[*cacheEntry]byte)
		if curNode != nil {
			nextNode = lfu.freqList.InsertAfter(li, curNode)
		} else {
			nextNode = lfu.freqList.PushFront(li)
		}
	}

	e.freqParent = nextNode
	nextNode.Value.(*cacheFreq).freqMap[e] = 1
	if curNode != nil {
		// remove from current position
		lfu.remove(curNode, e)
	}
}

func (lfu *LFUCache) evicts() {

	if item := lfu.freqList.Front(); item != nil {

		for entry := range item.Value.(*cacheFreq).freqMap {
			delete(lfu.mapCache, entry.key)
			lfu.remove(item, entry)
			lfu.size--
			break
		}
	}
}

func (lfu *LFUCache) remove(listItem *list.Element, e *cacheEntry) {
	freq := listItem.Value.(*cacheFreq)
	delete(freq.freqMap, e)

	if len(freq.freqMap) == 0 {
		lfu.freqList.Remove(listItem)
	}
}
