/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/31 3:37 下午
 */

package _minimum_index_sum_of_two_lists

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	var res []string
	var min = math.MaxInt32
	var index = make(map[string]int)

	for i, r := range list1 {
		index[r] = i
	}

	for i := 0; i < len(list2); i++ {
		if idx, has := index[list2[i]]; has {
			if i+idx <= min {
				if i+idx < min {
					res = res[len(res):]
				}
				min = i + idx
				res = append(res, list2[i])
			}
		}
	}

	return res
}
