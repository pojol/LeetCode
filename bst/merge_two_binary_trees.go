package bst

// https://leetcode.com/problems/merge-two-binary-trees/

func mergeTree(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil && t2 != nil {
		return t2
	} else if t1 != nil && t2 == nil {
		return t1
	}

	left := mergeTree(t1.Left, t2.Left)
	right := mergeTree(t1.Right, t2.Right)

	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  left,
		Right: right,
	}
}

// MergeTrees 合并两个二叉树
func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return mergeTree(t1, t2)
}
