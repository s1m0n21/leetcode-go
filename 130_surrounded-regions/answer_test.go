/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/27 5:00 下午
 */

package _surrounded_regions

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  [][]byte
		expect [][]byte
	}{
		{
			[][]byte{
				{'X','X','X','X'},
				{'X','O','O','X'},
				{'X','X','O','X'},
				{'X','O','X','X'},
			},
			[][]byte{
				{'X','X','X','X'},
				{'X','X','X','X'},
				{'X','X','X','X'},
				{'X','O','X','X'},
			},
		},
		{
			[][]byte{
				{'X','X','O','X'},
				{'X','O','O','X'},
				{'O','X','O','X'},
				{'X','O','X','O'},
			},
			[][]byte{
				{'X','X','O','X'},
				{'X','O','O','X'},
				{'O','X','O','X'},
				{'X','O','X','O'},
			},
		},
	}

	for _, test := range tests {
		c := copySlice(test.input)
		if solve(c); !reflect.DeepEqual(c, test.expect) {
			t.Errorf("input = %+v, expect = %+v, actual = %+v", test.input, test.expect, c)
		}
	}
}

func copySlice(src [][]byte) [][]byte {
	var c [][]byte
	for _, r := range src {
		c = append(c, append([]byte(nil), r...))
	}
	return c
}
