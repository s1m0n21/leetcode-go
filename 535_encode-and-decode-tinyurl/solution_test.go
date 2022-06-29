/**
 * @Author         : s1m0n21
 * @Description    : Test solution
 * @Project        : leetcode-go
 * @File           : solution_test.go
 * @Date           : 2022/6/29 15:51
 */

package _encode_and_decode_tinyurl

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	codec := Constructor()

	short := codec.encode("https://leetcode.com/problems/design-tinyurl")
	fmt.Println(short, len(short))
	fmt.Println(codec.decode(short))
}
