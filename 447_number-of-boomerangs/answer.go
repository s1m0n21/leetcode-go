/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/number-of-boomerangs/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/11 7:44 下午
 */

package _number_of_boomerangs

func numberOfBoomerangs(points [][]int) int {
	var out = 0
	for i := 0; i < len(points); i++ {
		var freq = make(map[int]int)
		for j := 0; j < len(points); j++ {
			if j != i {
				freq[dis(points[i], points[j])]++
			}
		}

		for _, f := range freq {
			if f >= 2 {
				out += f * (f - 1)
			}
		}
	}

	return out
}

func dis(a, b []int) int {
	return (a[0] - b[0]) * (a[0] - b[0]) + (a[1] - b[1]) * (a[1] - b[1])
}
