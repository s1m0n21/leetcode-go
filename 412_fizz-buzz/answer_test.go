/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/13 9:40 上午
 */

package _fizz_buzz

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect []string
	}{
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
		{0, []string{}},
	}

	for i, test := range tests {
		if actual := fizzBuzz(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %d, expect = %#v, actual = %#v", i, test.input, test.expect, actual)
		}
	}
}
