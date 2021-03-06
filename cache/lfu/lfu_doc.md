## LFU

* 数据结构设计
* increment 函数解析
* 示例图

#### 数据结构设计
> **map** `哈希`+ **freqList** `使用频次链表`+ **orderList** `访问先后链表`

> 1. 使用双向链表实现，将使用频次最低的节点设置在链表的头尾这样能保证O(1)的时间复杂度内操作这个节点。
> 2. 使用双层双向链表，外层负责使用频次，内层负责访问顺序可保证同样访问频次的节点排序

```go
/*
        +--> +---+ +--> +---+ +---> +---+   使用频率
            | 1 |      | 2 |       | 5 |
        <--+ +---+ <--+ +---+ <---+ +---+
                ^          ^           ^
            +   +      +   +       +   +
            v          v           v
            +---+      +---+       +---+
            | a |      | c |       | d |
            +---+      +---+       +---+
                ^
            +   +
            v
            +---+
            | b |
            +---+ 访问先后
*/
```
> 

#### increment
>   * curNode     当前节点的父节点
>   * nextFreq    更新后的访问频率值
>   * nextNode    使用频率链表节点（访问先后链表的根节点

>   1. 新节点加入

```go
curNode := e.freqParent
if curNode == nil {
    nextFreq = 1    // 访问频率设置为1
    nextNode = lfu.freqList.Front() // 将新节点的父节点设置为，访问频率链表的头节点。
}
```
>   2. 旧节点更新

```go
nextFreq = curNode.Value.(*cacheFreq).freq + 1  // 访问频率+1
nextNode = curNode.Next()
```
>   3. 右移（创建一个新的列

```go
// 当下一个节点为空 或者 下一个节点的 freq 值不等于 当前freq+1 （ 没有对应的freq列 时
if nextNode == nil || nextNode.Value.(*cacheFreq).freq != nextFreq {
    newcol := new(cacheFreq)        // 建立新列 (图
	newcol.freq = nextFreq          // 新列的 freq 值
    newcol.entries = list.New()     // 新列中的顺序链表
    
	if curNode == nil {
        // 将下一个点设置为访问频率的头节点
        nextNode = lfu.freqList.PushFront(newcol)
	} else {
        // 将下一个节点设置为原有访问频率节点的后一个节点 (图 节点变化1
		nextNode = lfu.freqList.InsertAfter(newcol, curNode)
	}
}
```
>   4. 更新

```go
// 设置新节点的父节点  
e.freqParent = nextNode
// 保存新节点在访问先后的容器中,并且将队列中的位置纪录下来方便下次更新时删除
e.orderIndex = nextNode.Value.(*freqNode).orderList.PushBack(e)
// 删除当前列上的节点 ( 图 节点变化2
if curNode != nil { 
	lfu.remove(curNode, e.orderIndex)
}
```

#### 示例图
> **建立新列**
[![image.png](https://i.postimg.cc/2S5z7f1x/image.png)](https://postimg.cc/kBk37zFV)

> **节点变化1**
[![image.png](https://i.postimg.cc/Y2gfM5q8/image.png)](https://postimg.cc/4nfh1Fch)

> **节点变化2**
[![image.png](https://i.postimg.cc/Kjz2rPz8/image.png)](https://postimg.cc/mz0JTFCv)

