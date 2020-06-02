package binarysearchtree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

//插入节点
func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val: val}
	}
	if val > root.val {
		root.right = Insert(root.right, val)

	} else if val < root.val {
		root.left = Insert(root.left, val)
	}
	return root
}

//递归查找节点
func Find(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val > root.val {
		return Find(root.right, val)

	} else if val < root.val {
		return Find(root.left, val)
	} else {
		return root
	}
}

//非递归查找节点
func NoRecursiveFind(root *TreeNode, val int) *TreeNode {
	current := root
	for current != nil {
		if val > current.val {
			current = current.right
		} else if val < current.val {
			current = current.left
		} else {
			return current
		}
	}
	return nil
}

//中序遍历
func VisitInOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := VisitInOrder(root.left)
	right := VisitInOrder(root.right)
	return append(append(left, root.val), right...)
}

//后序遍历
func VisitPostOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := VisitInOrder(root.left)
	right := VisitInOrder(root.right)
	return append(append(left, right...), root.val)
}
