/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/2 5:06 下午
 */

package _letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", nil},
		{"2", []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		if actual := letterCombinations(test.input); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %s, expect = %#v, actual = %#v", test.input, test.expect, actual)
		}
	}
}
