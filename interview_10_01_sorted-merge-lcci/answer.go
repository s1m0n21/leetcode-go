/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sorted-merge-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/29 9:07 下午
 */

package interview_10_01_sorted_merge_lcci

func merge(A []int, m int, B []int, n int) {
	m--
	n--

	for i := len(A)-1; i >= 0; i-- {
		if m >= 0 && n >= 0 {
			if A[m] > B[n] {
				A[i] = A[m]
				m--
			} else {
				A[i] = B[n]
				n--
			}
		} else if m >= 0 {
			A[i] = A[m]
			m--
		} else {
			A[i] = B[n]
			n--
		}
	}
}
