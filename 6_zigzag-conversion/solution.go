/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/zigzag-conversion/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/9 10:01
 */

package _zigzag_conversion

func convert(s string, numRows int) string {
	var c = make([][]rune, numRows)

	i, up := 0, false
	for _, b := range s {
		c[i] = append(c[i], b)
		if numRows > 1 {
			if !up {
				if i+1 >= numRows {
					up = true
					i--
				} else {
					i++
				}
			} else {
				if i-1 < 0 {
					up = false
					i++
				} else {
					i--
				}
			}
		}
	}

	var ret []rune
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			ret = append(ret, c[i][j])
		}
	}

	return string(ret)
}
