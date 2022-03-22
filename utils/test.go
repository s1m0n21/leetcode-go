/**
 * @Author         : s1m0n21
 * @Description    : Test util
 * @Project        : leetcode-go
 * @File           : test.go
 * @Date           : 2022/3/17 01:39
 */

package utils

import (
	"reflect"
	"testing"
	"time"
)

type ConfigOption func(config *testCaseConfig)

type testCaseConfig struct{}

type TestCase[I, E any] struct {
	t      *testing.T
	input  I
	expect E
	fn     func(I) E
	config *testCaseConfig
}

func NewTestCase[I, E any, T TestCase[I, E]](t *testing.T, fn func(I) E, options ...ConfigOption) T {
	config := &testCaseConfig{}
	for _, op := range options {
		op(config)
	}

	return T{fn: fn, config: config, t: t}
}

func (tc *TestCase[I, E]) SetInput(v I) {
	tc.input = v
}

func (tc *TestCase[I, E]) GetInput() I {
	return tc.input
}

func (tc *TestCase[I, E]) SetExpect(v E) {
	tc.expect = v
}

func (tc *TestCase[I, E]) GetExpect() E {
	return tc.expect
}

func (tc *TestCase[I, E]) RunTest() {
	st := time.Now()
	if actual := tc.fn(tc.input); !reflect.DeepEqual(actual, tc.expect) {
		tc.t.Errorf("\033[31mFAIL\033[0m | INPUT: %#v | EXPECT: \033[32m%#v\033[0m | ACTUAL: \033[31m%#v\033[0m", tc.input, tc.expect, actual)
	} else {
		tc.t.Logf("\033[32mPASS\033[0m | INPUT: %#v | ELAPSED: %s", tc.input, time.Now().Sub(st).String())
	}
}

func (tc *TestCase[I, E]) SetAndRun(input I, expect E) {
	tc.SetInput(input)
	tc.SetExpect(expect)
	tc.RunTest()
}
