/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/10/8 9:41 上午
 */

package _repeated_dna_sequences

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect []string
	}{
		{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
		{"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
		{"", nil},
		{"AAAA", nil},
		{"AAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	}

	for i, test := range tests {
		if actual := findRepeatedDnaSequences(test.input); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("%d: input = %s, expect = %#v, actual = %#v", i, test.input, test.expect, actual)
		}
	}
}
