/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/13 2:50 下午
 */

package _count_servers_that_communicate

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect int
	}{
		{[][]int{{1, 0}, {0, 1}}, 0},
		{[][]int{{1}, {0}}, 0},
		{[][]int{{1, 0}, {1, 1}}, 3},
		{[][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}, 4},
	}

	for i, test := range tests {
		if actual := countServers(test.input); actual != test.expect {
			t.Errorf("%d: input = %+v, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
