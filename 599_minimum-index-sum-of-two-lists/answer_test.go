/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/8/31 3:40 下午
 */

package _minimum_index_sum_of_two_lists

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		list1, list2 []string
	}

	tests := []struct {
		input  input
		expect []string
	}{
		{
			input{
				[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			},
			[]string{"Shogun"},
		},
		{
			input{
				[]string{"Shogun", "KFC", "Burger King", "Tapioca Express"},
				[]string{"KFC", "Shogun", "Burger King"},
			},
			[]string{"KFC", "Shogun"},
		},
		{
			input{
				[]string{"Shogun", "Piatti", "Tapioca Express", "Burger King", "KFC"},
				[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			},
			[]string{"Piatti"},
		},
	}

	for _, test := range tests {
		if actual := findRestaurant(test.input.list1, test.input.list2); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("input = %#v, exepct = %#v, actual = %#v", test.input, test.expect, actual)
		}
	}
}
