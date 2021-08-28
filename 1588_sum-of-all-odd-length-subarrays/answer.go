/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/29 12:18 上午
 */

package _sum_of_all_odd_length_subarrays

func sumOddLengthSubarrays(arr []int) int {
	var res int
	if len(arr) == 1 {
		res += arr[0]
		return res
	}

	for i := 1; i <= len(arr); i++ {
		if i % 2 == 0 { continue }
		for j := 0; j < len(arr); j++ {
			if j+i <= len(arr) {
				for _, n := range arr[j:j+i] {
					res += n
				}
			}
		}
	}

	return res
}
