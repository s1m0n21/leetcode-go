/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/destination-city/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/16 17:39
 */

package _destination_city

func destCity(paths [][]string) string {
	next := make(map[string]string)
	for _, path := range paths {
		next[path[0]] = path[0]
		if _, in := next[path[1]]; !in {
			next[path[1]] = ""
		}
	}

	for loc, nex := range next {
		if nex == "" {
			return loc
		}
	}

	return ""
}
