package problems

//id=1
/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//2个for直接计算
func twoSum1(nums []int, target int) []int {
	for i, _ := range nums {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 使用map存储，直接求需要的值，判断在map是否存在
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, 0) //map[num]index
	for i, n := range nums {
		m[n] = i
	}
	for i, n := range nums {
		need := target - n
		idx, ok := m[need]
		if ok && idx != i {
			return []int{i, idx}
		}
	}
	return nil
}

//id=2
/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	header := &ListNode{
		Val:  0,
		Next: nil,
	}
	//尾指针，尾指针的上一个指针  初始化为头指针
	end, beforeEnd := header, header
	//单个位计算满10的进位下一轮使用，本轮使用计算本位数字
	var tmp, value int
	// 题目中没有强调说明链表的长度可能是不一致的，而且有进位
	// 在for中每次取一位
	for {
		value = tmp
		if l1 != nil {
			value += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}
		if value >= 10 {
			tmp = 1
			end.Val = value - 10
			end.Next = &ListNode{}
			beforeEnd = end
			end = end.Next
		} else {
			tmp = 0
			end.Val = value
			end.Next = &ListNode{}
			beforeEnd = end
			end = end.Next
		}
		if l1 == nil && l2 == nil {
			break
		}
	}
	// 3位数增长到4位数
	if tmp == 1 {
		end.Val = 1
	} else {
		// 剔除无用的尾节点
		beforeEnd.Next = nil
	}
	return header
}

//id=3
/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	lastAppearPos := make([]int, 128) //128 ASICC码字符数量
	//窗口左，窗口右，最长字符串
	L, R, longest := 0, 0, 0
	for R < len(s) {
		if lastAppearPos[s[R]] != 0 && L < lastAppearPos[s[R]] {
			longest = max(longest, R-L)
			L = lastAppearPos[s[R]] //缩小窗口
		}
		//记录char出现的位置 后一位
		lastAppearPos[s[R]] = R + 1
		//右移
		R++
	}
	return max(longest, R-L)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//id=3
/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//if len(nums1) < len(nums2) {
	//	nums1, nums2 = nums2, nums1
	//}
	//k := (len(nums2)+len(nums1)+1)/2
	//fk1:=
	//
	//left2, right2 := 0, len(nums2)
	//var mid1, mid2 int
	//for left < right {
	//	if left+left2 == halfLen && nums1[right] < nums2[left2] {
	//		fmt.Printf("%d %d", left, left2)
	//		return 0
	//	}
	//	mid1 = (left + right) / 2
	//	if nums1[mid1] > nums2[left2] {
	//		right = mid1
	//	}
	//	mid2 = (left2 + right2) / 2
	//	if nums2[mid2] < nums1[left2] {
	//		left2 = mid2
	//	}
	//}
	return -1
}
