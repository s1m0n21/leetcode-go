/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/crawler-log-folder/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/15 09:18
 */

package _crawler_log_folder

func minOperations(logs []string) int {
	var steps int
	for _, log := range logs {
		if log == "../" {
			if steps > 0 {
				steps--
			}
		} else if log == "./" {
			continue
		} else {
			steps++
		}
	}

	return steps
}
