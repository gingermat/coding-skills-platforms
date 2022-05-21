package merge_two_tree

import (
	"fmt"
	"reflect"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) String() string {
	return fmt.Sprintf("TreeNode(Value=%v, Left=%v, Right=%v)", tn.Val, tn.Left, tn.Right)
}

var tests = []struct {
	root1 *TreeNode
	root2 *TreeNode
	want  *TreeNode
}{
	{
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		&TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	},
	{
		&TreeNode{
			Val: 1,
		},
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		&TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	},
}

func TestMergeTrees(t *testing.T) {
	for _, tt := range tests {
		got := mergeTrees(tt.root1, tt.root2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("mergeTrees(%v, %v) = %v, want %v",
				tt.root1, tt.root2, got, tt.want)
		}
	}
}
