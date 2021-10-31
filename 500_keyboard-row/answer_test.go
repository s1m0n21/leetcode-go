/*
 * @Author       : s1m0n21
 * @Description  : Test answer
 * @Date         : 2021/10/31 11:56 PM
 */

package _keyboar_drow

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct{
		input []string
		expect []string
	}{
		{[]string{"Hello","Alaska","Dad","Peace"}, []string{"Alaska","Dad"}},
		{[]string{"omk"}, nil},
		{[]string{"adsdf","sfd"}, []string{"adsdf","sfd"}},
	}

	for i, test := range tests {
		if actual := findWords(test.input); !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("%d: input = %+v, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
