package rb_tree

type Key string
type Value interface{}

type Color int

const (
	Red   Color = 0
	Black Color = 1
)

type RBNode struct {
	key   Key
	value Value
	left  *RBNode
	right *RBNode
	color Color
}

type RBBinarySearchTree struct {
	root *RBNode
}

func (t *RBBinarySearchTree) Get(k Key) Value {
	for node := t.root; node != nil; {
		if k < node.key {
			node = node.left
		} else if k > node.key {
			node = node.right
		} else {
			return node.value
		}
	}

	return nil
}

func (t *RBBinarySearchTree) Put(k Key, v Value) {
	t.root = put(t.root, k, v)
}

func put(x *RBNode, k Key, v Value) *RBNode {
	if x == nil {
		return &RBNode{key: k, value: v, color: Red}
	}

	if k < x.key {
		x.left = put(x.left, k, v)
	} else if k > x.key {
		x.right = put(x.right, k, v)
	} else {
		x.value = v
	}

	if isRed(x.right) && !isRed(x.left) {
		x = rotateLeft(x)
	}
	if isRed(x.left) && isRed(x.left.left) {
		x = rotateRight(x)
	}
	if isRed(x.left) && isRed(x.right) {
		flipColor(x)
	}

	return x
}

func isRed(n *RBNode) bool {
	if n == nil {
		return false
	}

	return n.color == Red
}

func rotateLeft(n *RBNode) *RBNode {
	if !isRed(n) {
		panic("node must be red")
	}

	x := n.right
	n.right = x.left

	x.left = n
	x.color = n.color
	n.color = Red

	return x
}

func rotateRight(n *RBNode) *RBNode {
	if !isRed(n) {
		panic("node can't be red")
	}

	x := n.left
	n.left = x.right

	x.right = n
	x.color = n.color
	n.color = Red

	return x
}

func flipColor(n *RBNode) {
	if isRed(n) {
		panic("node must be red")
	}
	if !isRed(n.left) {
		panic("node must be black")
	}
	if !isRed(n.right) {
		panic("node must be black")
	}

	n.color = Red
	n.left.color = Black
	n.right.color = Black
}

func NewRBBinarySearchTree() *RBBinarySearchTree {
	return &RBBinarySearchTree{}
}
