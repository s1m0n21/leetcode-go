/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/duplicate-zeros/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/17 09:18
 */

package _duplicate_zeros

func duplicateZeros(arr []int) {
	n := len(arr)
	top := 0
	i := -1
	for top < n {
		i++
		if arr[i] != 0 {
			top++
		} else {
			top += 2
		}
	}

	j := n - 1
	if top == n+1 {
		arr[j] = 0
		j--
		i--
	}

	for j >= 0 {
		arr[j] = arr[i]
		j--
		if arr[i] == 0 {
			arr[j] = arr[i]
			j--
		}
		i--
	}
}
