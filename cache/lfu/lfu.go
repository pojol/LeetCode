package lfu

import "container/list"

import "fmt"

// LFUCache  Least Frequently Used
type LFUCache struct {
	cap      int
	size     int
	mapCache map[int]*cacheEntry
	freqList *list.List // 频率链表
}

// 缓存中的实际对象
type cacheEntry struct {
	key        int
	value      int
	freqParent *list.Element // 指向列的头节点
	orderIndex *list.Element
}

type freqNode struct {
	freq      int
	orderList *list.List // 存储频率链表中的每一项，确保不重复(set
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

// 更新访问频率，将节点进行右移
func (lfu *LFUCache) increment(e *cacheEntry) {
	curNode := e.freqParent
	curIdx := e.orderIndex
	var nextFreq int
	var nextNode *list.Element

	if curNode == nil { // new node
		nextFreq = 1
		nextNode = lfu.freqList.Front() // head
	} else {
		nextFreq = curNode.Value.(*freqNode).freq + 1
		nextNode = curNode.Next()
	}

	if nextNode == nil || nextNode.Value.(*freqNode).freq != nextFreq {
		// create a new list entry
		newcol := new(freqNode)
		newcol.freq = nextFreq
		newcol.orderList = list.New()
		if curNode == nil {
			nextNode = lfu.freqList.PushFront(newcol)
		} else {
			nextNode = lfu.freqList.InsertAfter(newcol, curNode)
		}
	}

	e.freqParent = nextNode
	e.orderIndex = nextNode.Value.(*freqNode).orderList.PushBack(e)
	if curNode != nil {
		// remove from current position
		lfu.remove(curNode, curIdx)
	}
}

// 逐出一个节点
func (lfu *LFUCache) evicts() {

	if item := lfu.freqList.Front(); item != nil {
		delete(lfu.mapCache, item.Value.(*freqNode).orderList.Front().Value.(*cacheEntry).key)
		lfu.remove(item, item.Value.(*freqNode).orderList.Front())
		lfu.size--
	}
}

func (lfu *LFUCache) remove(freqItem *list.Element, orderItem *list.Element) {
	freq := freqItem.Value.(*freqNode)
	freq.orderList.Remove(orderItem)

	if freq.orderList.Len() == 0 {
		lfu.freqList.Remove(freqItem)
	}
}

func (lfu *LFUCache) print() {

	for i := lfu.freqList.Front(); i != nil; i = i.Next() {
		fmt.Println("freq ", i.Value.(*freqNode).freq)
		entries := i.Value.(*freqNode).orderList
		for j := entries.Front(); j != nil; j = j.Next() {
			fmt.Println("order", j.Value.(*cacheEntry).key, j.Value.(*cacheEntry).value)
		}
	}

	fmt.Println("---")

}
