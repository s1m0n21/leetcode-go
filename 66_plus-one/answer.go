/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/plus-one/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/28 7:39 下午
 */

package _plus_one

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}

	var res = make([]int, 0)
	n := 0
	for i := len(digits)-1; i >= 0; i-- {
		k := digits[i]
		if i == len(digits)-1 {
			k += 1
		} else {
			k += n
			n = 0
		}

		if k > 9 {
			n = k / 10
			k %= 10
		}

		res = append(res, k)
	}

	if n > 0 {
		res = append(res, n)
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-1-i], res[i]
	}

	return res
}
