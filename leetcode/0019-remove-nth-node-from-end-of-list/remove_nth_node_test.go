package remove_nth_node

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var tests = []struct {
	head *ListNode
	n    int
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
		2,
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	},
	{
		&ListNode{
			Val: 1,
		},
		1,
		nil,
	},
	{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
		1,
		&ListNode{
			Val: 1,
		},
	},
	{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
		2,
		&ListNode{
			Val: 2,
		},
	},
}

func TestRemoveNthFromEnd(t *testing.T) {
	for _, tt := range tests {
		got := removeNthFromEnd(tt.head, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("removeNthFromEnd(%v, %v) = %v, want %v", tt.head, tt.n, got, tt.want)
		}
	}
}
