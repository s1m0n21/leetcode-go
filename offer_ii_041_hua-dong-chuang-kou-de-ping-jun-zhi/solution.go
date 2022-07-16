/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/qIsx9U/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/16 13:35
 */

package offer_ii_041_hua_dong_chuang_kou_de_ping_jun_zhi

type MovingAverage struct {
	sum  int
	size int
	nums []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{size: size, nums: make([]int, 0, size)}
}

func (avg *MovingAverage) Next(val int) float64 {
	if len(avg.nums) == avg.size {
		avg.sum -= avg.nums[0]
		avg.nums = avg.nums[1:]
	}
	avg.nums = append(avg.nums, val)
	avg.sum += val

	return float64(avg.sum) / float64(len(avg.nums))
}
