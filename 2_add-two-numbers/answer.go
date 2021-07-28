/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/add-two-numbers/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/4 8:38 下午
 */

package _add_two_numbers

import "github.com/s1m0n21/leetcode-go/utils"

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
     var tail *utils.ListNode
     var head *utils.ListNode
     var carry = 0

     for l1 != nil || l2 != nil {
          n1 := 0
          if l1 != nil {
               n1 = l1.Val
               l1 = l1.Next
          }

          n2 := 0
          if l2 != nil {
               n2 = l2.Val
               l2 = l2.Next
          }

          sum := n1 + n2 + carry
          carry = sum / 10
          sum = sum % 10

          if head != nil {
               tail.Next = &utils.ListNode{Val: sum}
               tail = tail.Next
               continue
          }

          head = &utils.ListNode{Val: sum}
          tail = head
     }

     if carry > 0 {
          tail.Next = &utils.ListNode{Val: carry}
     }

     return head
}
