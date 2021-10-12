/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/12 7:29 下午
 */

package _distribute_candies_to_people

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		candies, numPeople int
	}

	tests := []struct {
		input  input
		expect []int
	}{
		{input{7, 4}, []int{1, 2, 3, 1}},
		{input{10, 3}, []int{5, 2, 3}},
		{input{10, 1}, []int{10}},
	}

	for i, test := range tests {
		if actual := distributeCandies(test.input.candies, test.input.numPeople); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
