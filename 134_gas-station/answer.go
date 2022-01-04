/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/gas-station/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/4 12:50 PM
 */

package _gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		gasSum, costSum, cnt := 0, 0, 0
		for ; cnt < n; cnt++ {
			j := (i + cnt) % n
			gasSum += gas[j]
			costSum += cost[j]
			if costSum > gasSum {
				break
			}
		}
		if cnt == n {
			return i
		} else {
			i = i + cnt + 1
		}
	}

	return -1
}
