package middle_of_ll

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
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	},
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
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
		&ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
				},
			},
		},
	},
	{
		&ListNode{
			Val: 1,
		},
		&ListNode{
			Val: 1,
		},
	},
}

func TestMiddleNode(t *testing.T) {
	for _, tt := range tests {
		got := middleNode(tt.head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("middleNode(%v) = %v, want %v", tt.head, got, tt.want)
		}
	}
}
