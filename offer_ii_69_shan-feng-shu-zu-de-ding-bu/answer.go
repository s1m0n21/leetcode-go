/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/B1IidL/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/14 9:27 上午
 */

package offer_ii_69_shan_feng_shu_zu_de_ding_bu

func peakIndexInMountainArray(arr []int) int {
	var ans int
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[ans] {
			ans = i
		}
	}
	return ans
}
