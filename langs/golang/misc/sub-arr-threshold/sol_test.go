package subarrthreshold

import (
	"fmt"
	"testing"
)

func validSubarraySize(nums []int, threshold int) int {
	l, r := 0, len(nums)

	for i := r; i > l; i-- {
		if nums[i-1] < threshold/i {
			// 3
			fmt.Println("less than threshold/len(nums)", 6/4)
		}
	}

	return 0
}

func TestSubarray(t *testing.T) {
	nums := []int{1, 3, 4, 3, 1}
	// nums := []int{3, 4, 1, 4, 1}
	threshold := 6
	validSubarraySize(nums, threshold)
}
