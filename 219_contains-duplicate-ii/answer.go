/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/contains-duplicate-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/17 10:58 下午
 */

package _contains_duplicate_ii

func containsNearbyDuplicate(nums []int, k int) bool {
	var record = make(map[int]int, k)
	for i, v := range nums {
		if record[v] > 0 && i + 1 - record[v] <= k {
			return true
		}

		if i >= k && record[nums[i - k]] > 0 {
			delete(record, nums[i - k])
		}

		record[v] = i + 1
	}

	return false
}
