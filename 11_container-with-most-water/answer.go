/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/container-with-most-water/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/29 9:10 上午
 */

package _container_with_most_water

func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1
	width := right
	for left < right {
		tmp := min(height[left], height[right]) * width
		if tmp > max {
			max = tmp
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

		width--
	}
	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
