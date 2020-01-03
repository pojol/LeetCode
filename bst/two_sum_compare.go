package bst

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

func step(root *TreeNode, iarr *[]int) {

	if root != nil {
		step(root.Left, iarr)
		step(root.Right, iarr)

		*iarr = append(*iarr, root.Val)
	}
}

// FindTarget 在二叉树中搜索是否有两个节点的值相加==k的情况
func FindTarget(root *TreeNode, k int) bool {

	iarr := []int{}
	step(root, &iarr)
	//sort.Ints(iarr)

	for i := 0; i < len(iarr); i++ {
		for j := i + 1; j < len(iarr); j++ {
			if iarr[i]+iarr[j] == k {
				return true
			}
		}
	}

	return false
}
