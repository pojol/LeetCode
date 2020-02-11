package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFKPWSS(t *testing.T) {

	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	rq := kSmallestPairs(nums1, nums2, k)
	assert.Equal(t, len(rq), 3)
	assert.Equal(t, rq, [][]int{[]int{1, 2}, []int{1, 4}, []int{1, 6}})

}
