package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Input:
	 5
	/ \
   3   6
  / \   \
 2   4   7

	Target = 9

	Output: True
*/

func TestFindTarget(t *testing.T) {

	assert.Equal(t, FindTarget(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}, 9), true)

}
