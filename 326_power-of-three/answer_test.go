/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/9/23 9:40 上午
 */

 package _power_of_three

import (
  "testing"
)

 func TestAnswer(t *testing.T) {
  tests := []struct{
    input  int
    expect bool
  }{
    {27, true},
    {0, false},
    {9, true},
    {45, false},
  }

  for _, test := range tests {
    if actual := isPowerOfThree(test.input); actual != test.expect {
      t.Errorf("input = %d, expect = %t, actual = %t", test.input, test.expect, actual)
    }
  }
 }

