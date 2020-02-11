package slidingwindow

import "container/heap"

// Pair 值结构
type Pair struct {
	n1idx int
	n2idx int
	sum   int
}

// PairHeap pair堆
type PairHeap []Pair

func (ph PairHeap) Len() int           { return len(ph) }
func (ph PairHeap) Less(i, j int) bool { return ph[i].sum < ph[j].sum }
func (ph PairHeap) Swap(i, j int)      { ph[i], ph[j] = ph[j], ph[i] }

// Push heap push实现
func (ph *PairHeap) Push(x interface{}) {
	*ph = append(*ph, x.(Pair))
}

// Pop heap pop实现
func (ph *PairHeap) Pop() interface{} {
	root := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return root
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	nums1len := len(nums1)
	nums2len := len(nums2)
	if nums1len == 0 || nums2len == 0 || k == 0 {
		return nil
	}

	if k > nums1len*nums2len {
		k = nums1len * nums2len
	}

	ph := new(PairHeap)
	heap.Init(ph)

	for i1, n1 := range nums1 {
		for i2, n2 := range nums2 {
			heap.Push(ph, Pair{n1idx: i1, n2idx: i2, sum: n1 + n2})
		}
	}

	res := [][]int{}
	for i := 0; i < k; i++ {
		p := heap.Pop(ph).(Pair)
		res = append(res, []int{nums1[p.n1idx], nums2[p.n2idx]})
	}

	return res
}
