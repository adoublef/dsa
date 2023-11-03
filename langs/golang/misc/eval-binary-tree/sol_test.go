package binary_tree

// https://appliedgo.net/generictree/
// https://go.dev/blog/intro-generics
// https://pkg.go.dev/golang.org/x/exp/constraints

import (
	"testing"
)

func algo[T int](root *treeNode[T]) bool {
	if root.left != nil {
		switch l, r := algo(root.left), algo(root.right); root.val {
		case 2:
			return l || r
		default:
			return l && r
		}
	}
	return root.val != 0
}

func TestAlgo(t *testing.T) {
	root := newTree([]int{2, 1, 3, 0, 1})
	if !algo(root) {
		t.Error("Expected true, got", algo(root))
	}

	root = &treeNode[int]{val: 0}
	if algo(root) {
		t.Error("Expected true, got", algo(root))
	}
}
