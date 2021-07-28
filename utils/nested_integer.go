/**
 * @Author         : s1m0n21
 * @Description    :
 * @Project        : leetcode-go
 * @File           : nested_integer.go
 * @Date           : 2021/7/23 3:43 下午
 */

package utils

type NestedInteger struct {
	value interface{}
}

func MakeNestedInteger(v int) *NestedInteger {
	return &NestedInteger{value: v}
}

func MakeNestedIntegerList(v ...int) *NestedInteger {
	var list = make([]*NestedInteger, 0)
	for _, n := range v {
		list = append(list, &NestedInteger{value: n})
	}
	return &NestedInteger{value: list}
}

func (this NestedInteger) IsInteger() bool {
	_, ok := this.value.(int)
	return ok
}

func (this NestedInteger) GetInteger() int {
	if this.IsInteger() {
		return this.value.(int)
	}
	panic("this is not a integer")
}

func (this *NestedInteger) SetInteger(value int) {
	this.value = value
}

func (this *NestedInteger) Add(elem *NestedInteger) *NestedInteger {
	if list, ok := this.value.([]*NestedInteger); ok {
		this.value = append(list, elem)
		return this
	}
	panic("this is not a list")
}

func (this NestedInteger) GetList() []*NestedInteger {
	if list, ok := this.value.([]*NestedInteger); ok {
		return list
	}
	panic("this is not a list")
}

