package main

import "fmt"

func main() {
	var l1 = ListNode{
		Val:2,
		Next:&ListNode{
			Val:4,
			Next:nil,
		},
	}
	var l2 = ListNode{
		Val:8,
		Next:nil,
	}


	res := addTwoNumbers(&l1, &l2)
	printL(res)
}

type ListNode struct {
	Val int
	Next *ListNode
 }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rtl := new(ListNode)
	rtlCur := rtl
	var carry int
	for {
		var a,b,curNum int
		if l1 !=nil || l2 != nil || carry != 0 {
			if l1 != nil {
				a = l1.Val
				l1 = l1.Next
			}

			if l2 != nil {
				b = l2.Val
				l2 = l2.Next
			}
			curNum = a + b + carry

			if curNum >= 10 {
				carry = curNum / 10
				curNum = curNum % 10

			} else {
				carry = 0
			}

			rtlCur.Next = &ListNode{curNum, nil}
			rtlCur = rtlCur.Next
		} else {
			break
		}
	}

	return rtl.Next
}


func printL(res * ListNode) {
	var node * ListNode

	node = res
	for node != nil  {
		fmt.Printf("%v ", node.Val)
		node = node.Next
	}

	fmt.Println("")
}