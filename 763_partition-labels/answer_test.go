/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/21 4:21 PM
 */

package _partition_labels

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
	}

	for i, test := range tests {
		if actual := partitionLabels(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %s, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
