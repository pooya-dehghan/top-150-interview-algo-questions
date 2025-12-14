package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	myMagicMap := make(map[*ListNode]int)

	hasHead := true

	for hasHead {
		myMagicMap[head]++

		if myMagicMap[head] > 1 {
			return true
		}

		if head.Next == nil {
			return false
		}

		head = head.Next
	}

	return false
}

func main() {
	head0 := &ListNode{Val: -21}
	head1 := &ListNode{Val: 10}
	head2 := &ListNode{Val: 17}
	head3 := &ListNode{Val: 8}
	head4 := &ListNode{Val: 4}
	head5 := &ListNode{Val: 26}
	head6 := &ListNode{Val: 5}
	head7 := &ListNode{Val: 35}
	head8 := &ListNode{Val: 33}
	head9 := &ListNode{Val: -7}
	head10 := &ListNode{Val: -16}
	head11 := &ListNode{Val: 27}
	head12 := &ListNode{Val: -12}
	head13 := &ListNode{Val: 6}
	head14 := &ListNode{Val: 29}
	head15 := &ListNode{Val: -12}
	head16 := &ListNode{Val: 5}
	head17 := &ListNode{Val: 9}
	head18 := &ListNode{Val: 20}
	head19 := &ListNode{Val: 14}
	head20 := &ListNode{Val: 14}
	head21 := &ListNode{Val: 2}
	head22 := &ListNode{Val: 13}
	head23 := &ListNode{Val: -24}
	head24 := &ListNode{Val: 21}
	head25 := &ListNode{Val: 23}
	head26 := &ListNode{Val: -21}
	head27 := &ListNode{Val: 5}

	head0.Next = head1
	head1.Next = head2
	head2.Next = head3
	head3.Next = head4
	head4.Next = head5
	head5.Next = head6
	head6.Next = head7
	head7.Next = head8
	head8.Next = head9
	head9.Next = head10
	head10.Next = head11
	head11.Next = head12
	head12.Next = head13
	head13.Next = head14
	head14.Next = head15
	head15.Next = head16
	head16.Next = head17
	head17.Next = head18
	head18.Next = head19
	head19.Next = head20
	head20.Next = head21
	head21.Next = head22
	head22.Next = head23
	head23.Next = head24
	head24.Next = head25
	head25.Next = head26
	head26.Next = head27

	output := hasCycle(head0)

	fmt.Println("output:::", output)
}
