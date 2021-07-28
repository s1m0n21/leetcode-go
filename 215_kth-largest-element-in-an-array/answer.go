/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/28 11:26 下午
 */

package _kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	qsort(&nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func qsort(nums *[]int, left int, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	qsort(nums, left, p-1)
	qsort(nums, p+1, right)
}

func partition(nums *[]int, left int, right int) int {
	p := (*nums)[right]
	i := left - 1
	for j := left; j < right; j++ {
		if (*nums)[j] < p {
			i++
			temp := (*nums)[i]
			(*nums)[i] = (*nums)[j]
			(*nums)[j] = temp
		}
	}
	temp := (*nums)[i+1]
	(*nums)[i+1] = (*nums)[right]
	(*nums)[right] = temp
	return i + 1
}
