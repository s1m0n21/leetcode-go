/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sort-an-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/18 10:00 上午
 */

package _sort_an_array

//More better
//func sortArray(nums []int) []int {
//	sort.Ints(nums)
//	return nums
//}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && pivot <= nums[right] {
			right--
		}
		nums[left] = nums[right]

		for left < right && pivot >= nums[left] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}
