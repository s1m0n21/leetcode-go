/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/majority-element-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/22 5:04 下午
 */

package _majority_element_ii

func majorityElement(nums []int) []int {
	var ans []int
	var major1, major2 int
	var vote1, vote2 int
	var cnt1, cnt2 int

	for _, n := range nums {
		if vote1 > 0 && n == major1 {
			vote1++
		} else if vote2 > 0 && n == major2 {
			vote2++
		} else if vote1 == 0 {
			major1 = n
			vote1++
		} else if vote2 == 0 {
			major2 = n
			vote2++
		} else {
			vote1--
			vote2--
		}
	}

	for _, n := range nums {
		if vote1 > 0 && n == major1 {
			cnt1++
		}
		if vote2 > 0 && n == major2 {
			cnt2++
		}
	}

	if vote1 > 0 && cnt1 > len(nums)/3 {
		ans = append(ans, major1)
	}
	if vote2 > 0 && cnt2 > len(nums)/3 {
		ans = append(ans, major2)
	}

	return ans
}
