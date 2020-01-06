package lru

import (
	"container/list"
)

// https://leetcode.com/problems/lru-cache/

// LRUCache lru结构
type LRUCache struct {
	cap      int
	mapCache map[int]*list.Element
	list     *list.List
}

// LRUNode 存放数据
type LRUNode struct {
	key int
	val int
}

// Constructor 构造
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:      capacity,
		mapCache: make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get 获取数据
func (lp *LRUCache) Get(key int) int {

	if elem, ok := lp.mapCache[key]; ok {
		lp.list.MoveToFront(elem)
		return lp.mapCache[key].Value.(LRUNode).val
	}

	return -1
}

// Put 存放数据
func (lp *LRUCache) Put(key int, value int) {

	if elem, ok := lp.mapCache[key]; ok {
		lp.list.MoveToFront(elem)
		elem.Value = LRUNode{
			key: key,
			val: value,
		}
	} else {

		if lp.list.Len() >= lp.cap { // 淘汰尾部元素
			delete(lp.mapCache, lp.list.Back().Value.(LRUNode).key)
			lp.list.Remove(lp.list.Back())
		}

		nod := LRUNode{
			key: key,
			val: value,
		}
		lp.list.PushFront(nod)
		lp.mapCache[nod.key] = lp.list.Front()
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
