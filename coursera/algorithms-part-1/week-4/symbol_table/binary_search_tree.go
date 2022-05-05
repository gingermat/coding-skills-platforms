package symbol_table

type BSTNode struct {
	key   Key
	value Value
	left  *BSTNode
	right *BSTNode
	count int
}

func NewBSTNode(k Key, v Value) *BSTNode {
	return &BSTNode{
		key:   k,
		value: v,
	}
}

type BinarySearchTree struct {
	root *BSTNode
}

func (t *BinarySearchTree) Get(k Key) Value {
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

func (t *BinarySearchTree) Put(k Key, v Value) {
	t.root = put(t.root, k, v)
}

func put(x *BSTNode, k Key, v Value) *BSTNode {
	if x == nil {
		return &BSTNode{key: k, value: v}
	}

	if k < x.key {
		x.left = put(x.left, k, v)
	} else if k > x.key {
		x.right = put(x.right, k, v)
	} else {
		x.value = v
	}

	x.count = 1 + size(x.left) + size(x.right)

	return x
}

func (t *BinarySearchTree) Delete(k Key) {
	t.root = delete(t.root, k)
}

func delete(x *BSTNode, k Key) *BSTNode {
	if x == nil {
		return nil
	}

	if k < x.key {
		x.left = delete(x.left, k)
		return x
	}
	if k > x.key {
		x.right = delete(x.right, k)
		return x
	}

	if x.left == nil && x.right == nil {
		x = nil
		return x
	}

	if x.left == nil {
		x = x.right
		return x
	}
	if x.right == nil {
		x = x.left
		return x
	}

	tmp, x := x, min(x)

	x.left = tmp.left
	x.right = deleteMin(tmp.right)

	return x
}

func (t *BinarySearchTree) deleteMin() {
	t.root = deleteMin(t.root)
}

func deleteMin(x *BSTNode) *BSTNode {
	if x.left == nil {
		return x.right
	}

	x.left = deleteMin(x.left)
	x.count = 1 + size(x.left) + size(x.right)

	return x
}

func min(x *BSTNode) *BSTNode {
	tmp := x

	for tmp != nil && tmp.left != nil {
		tmp = tmp.left
	}

	return tmp
}

func (t *BinarySearchTree) Size() int {
	return size(t.root)
}

func size(x *BSTNode) int {
	if x == nil {
		return 0
	}

	return x.count
}

func (t *BinarySearchTree) Keys() []Key {
	var keys []Key

	inorder(t.root, &keys)
	return keys
}

func inorder(x *BSTNode, keys *[]Key) {
	if x == nil {
		return
	}

	inorder(x.left, keys)
	*keys = append(*keys, x.key)
	inorder(x.right, keys)
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}
