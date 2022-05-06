/**
 * @Author         : s1m0n21
 * @Description    : Test util
 * @Project        : leetcode-go
 * @File           : test.go
 * @Date           : 2022/3/17 01:39
 */

package utils

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
	"reflect"
	"testing"
	"time"
)

type testCaseConfig[T any] struct {
	checkFunc        func(a, b T) bool
	outputInAnyOrder bool
}

type TestCase[I, E any] struct {
	t           *testing.T
	input       I
	expect      E
	fn          func(I) E
	config      testCaseConfig[E]
}

func NewTestCase[I, E any](t *testing.T, fn func(I) E) TestCase[I, E] {
	return TestCase[I, E]{fn: fn, config: testCaseConfig[E]{}, t: t}
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

func (tc *TestCase[I, E]) check(v E) bool {
	if tc.config.checkFunc != nil {
		return tc.config.checkFunc(v, tc.expect)
	}
	if tc.config.outputInAnyOrder {
		elements := make(map[uint64]int)
		exp := reflect.ValueOf(tc.expect)
		out := reflect.ValueOf(v)

		if exp.Len() != out.Len() {
			return false
		}

		for i := 0; i < exp.Len(); i++ {
			elements[tc.hash(exp.Index(i).Interface())]++
			elements[tc.hash(out.Index(i).Interface())]--
		}

		for _, n := range elements {
			if n != 0 {
				return false
			}
		}

		return true
	}
	return reflect.DeepEqual(v, tc.expect)
}

func (tc *TestCase[I, E]) RunTest() {
	var copies []any
	if tc.isSlice(tc.input) {
		copies = tc.copySlice(tc.input)
	}

	st := time.Now()
	actual := tc.fn(tc.input)
	elapsed := time.Now().Sub(st).String()

	if !tc.check(actual) {
		if copies != nil {
			tc.failed(copies, actual)
		} else {
			tc.failed(tc.input, actual)
		}
	} else {
		if copies != nil {
			tc.success(copies, elapsed)
		} else {
			tc.success(tc.input, elapsed)
		}
	}
}

func (tc *TestCase[I, E]) SetAndRun(input I, expect E) {
	tc.SetInput(input)
	tc.SetExpect(expect)
	tc.RunTest()
}

func (tc *TestCase[I, E]) success(input any, elapsed string) {
	tc.t.Logf("\033[32mPASS\033[0m | INPUT: %v | ELAPSED: %s", input, elapsed)
}

func (tc *TestCase[I, E]) failed(input, actual any) {
	tc.t.Errorf("\033[31mFAIL\033[0m | INPUT: %v | EXPECT: \033[32m%v\033[0m | ACTUAL: \033[31m%v\033[0m",
		input, tc.expect, actual)
}

func (tc *TestCase[I, E]) hash(v any) uint64 {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		panic(err)
	}

	hash := fnv.New64()
	if _, err := hash.Write(buf.Bytes()); err != nil {
		panic(err)
	}

	return hash.Sum64()
}

func (tc *TestCase[I, E]) copySlice(v any) []any {
	var copies []any

	value := reflect.ValueOf(v)
	for i := 0; i < value.Len(); i++ {
		if value.Index(i).Kind() == reflect.Slice || value.Index(i).Kind() == reflect.Array {
			copies = append(copies, tc.copySlice(value.Index(i).Interface()))
		} else {
			copies = append(copies, value.Index(i).Interface())
		}
	}

	return copies
}

func (tc *TestCase[I, E]) isSlice(v any) bool {
	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

func (tc *TestCase[I, E]) ConfigSetCheckFunc(fn func(a, b E) bool) {
	tc.config.checkFunc = fn
}

func (tc *TestCase[I, E]) ConfigSetOutputInAnyOrder() {
	kind := reflect.TypeOf(tc.expect).Kind()
	if !tc.isSlice(tc.expect) {
		tc.t.Fatalf("output type must be a `slice` or an `array`, got `%s`", kind.String())
	}

	tc.config.outputInAnyOrder = true
}
