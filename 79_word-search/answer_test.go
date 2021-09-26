/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/26 1:42 下午
 */

package _word_search

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		board [][]byte
		word  string
	}

	tests := []struct {
		input  input
		expect bool
	}{
		{input{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED"}, true},
		{input{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "SEE"}, true},
		{input{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCB"}, false},
	}

	for _, test := range tests {
		if actual := exist(test.input.board, test.input.word); actual != test.expect {
			t.Errorf("input = %+v, expect = %t, actual = %t", test.input, test.expect, actual)
		}
	}
}
