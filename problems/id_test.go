package problems

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	addTwoNumbers(l1, l2)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Print(lengthOfLongestSubstring("dddd"))
}

func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Print(findMedianSortedArrays([]int{1, 3, 7}, []int{2, 4, 6}))
	//	123467
}
