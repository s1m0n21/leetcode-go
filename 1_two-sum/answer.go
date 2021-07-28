/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/two-sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/7 1:28 上午
 */

package _two_sum

func twoSum(nums []int, target int) []int {
	var index = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		n := target - nums[i]

		if idx, has := index[n]; has {
			if idx == i {
				continue
			}

			return []int{i, idx}
		}
	}

	return nil
}
