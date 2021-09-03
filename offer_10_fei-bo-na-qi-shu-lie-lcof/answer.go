/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/4 2:38 上午
 */

package offer_10_fei_bo_na_qi_shu_lie_lcof

func fib(n int) int {
	n1, n2 := 0, 1
	for i := 0; i < n; i++ {
		n1, n2 = n2, (n1+n2)%(1e9+7)
	}

	return n1
}
