package merge_lists

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
	list1 *ListNode
	list2 *ListNode
	want  *ListNode
}{
	{
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		},
	},
	{
		nil,
		nil,
		nil,
	},
	{
		nil,
		&ListNode{
			Val: 0,
		},
		&ListNode{
			Val: 0,
		},
	},
}

func TestMergeTwoLists(t *testing.T) {
	for _, tt := range tests {
		got := mergeTwoLists(tt.list1, tt.list2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("mergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, got, tt.want)
		}
	}
}
