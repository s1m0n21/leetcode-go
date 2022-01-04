/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/4 8:39 PM
 */

package _k_closest_points_to_origin

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		points [][]int
		k      int
	}

	tests := []struct {
		input  input
		expect [][]int
	}{
		{input{[][]int{{1, 3}, {-2, 2}}, 1}, [][]int{{-2, 2}}},
		{input{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2}, [][]int{{-2, 4}, {3, 3}}},
	}

	for i, test := range tests {
		if actual := kClosest(test.input.points, test.input.k); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
