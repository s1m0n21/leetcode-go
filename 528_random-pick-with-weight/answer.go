/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/random-pick-with-weight/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/30 9:58 上午
 */

package _random_pick_with_weight

import "math/rand"

type Solution struct {
	weights []int
	sum     int
}

func Constructor(w []int) Solution {
	s := Solution{weights: w}
	for i := 0; i < len(w); i++ {
		s.sum += w[i]
	}
	return s
}

func (this *Solution) PickIndex() int {
	if len(this.weights) == 1 {
		return 0
	}

	r := int(rand.Float64() * float64(this.sum))
	t := 0
	for i, w := range this.weights {
		t += w
		if t > r {
			return i
		}
	}

	return len(this.weights) - 1
}
