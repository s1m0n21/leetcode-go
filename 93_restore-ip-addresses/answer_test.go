/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/5 12:54 PM
 */

package _restore_ip_addresses

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect []string
	}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"0000", []string{"0.0.0.0"}},
		{"1111", []string{"1.1.1.1"}},
		{"010010", []string{"0.10.0.10", "0.100.1.0"}},
		{"101023", []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
	}

	for i, test := range tests {
		if actual := restoreIpAddresses(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %s, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
