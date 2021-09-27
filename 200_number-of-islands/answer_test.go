/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/27 3:40 下午
 */

package _number_of_islands

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]byte
		expect int
	}{
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '1'},
				{'1', '1', '0', '1', '1'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			4,
		},
	}

	for _, test := range tests {
		if actual := numIslands(test.input); actual != test.expect {
			t.Errorf("input = %+v, exepct = %d, actual = %d", test.input, test.expect, actual)
		}
	}
}
