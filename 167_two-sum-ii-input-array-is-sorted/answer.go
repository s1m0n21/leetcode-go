/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/28 11:57 下午
 */

package _two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	i, j := 0, len(numbers)-1
	for i != j {
		if numbers[i] + numbers[j] < target {
			i++
			continue
		} else if numbers[i] + numbers[j] > target {
			j--
			continue
		}

		return []int{i+1, j+1}
	}

	return nil
}
