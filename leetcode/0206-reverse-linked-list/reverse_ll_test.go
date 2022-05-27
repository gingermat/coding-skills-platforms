package reverse_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	return fmt.Sprintf("ListNode(Next=%v, Val=%v)", ln.Next, ln.Val)
}

var tests = []struct {
	head *ListNode
	want *ListNode
}{
	{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
	},
	{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	},
	{
		nil,
		nil,
	},
}

func TestReverseList(t *testing.T) {
	for _, tt := range tests {
		got := reverseList(tt.head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("reverseList(%v) = %v, want %v", tt.head, got, tt.want)
		}
	}
}
