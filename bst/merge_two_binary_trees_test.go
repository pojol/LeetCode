package bst

import "testing"

func TestMergeTrees(t *testing.T) {

	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: nil,
	}

	tree2 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	MergeTrees(tree1, tree2)

}
