package binarysearchtree

import (
	"strconv"
	"strings"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	var root *TreeNode
	root = Insert(root, 5)
	root = Insert(root, 3)
	root = Insert(root, 7)
	root = Insert(root, 1)
	if arrayToString(VisitInOrder(root)) != "1,3,5,7" {
		t.Error("In order visit wrong order")
	}
	if arrayToString(VisitPostOrder(root)) != "1,3,7,5" {
		t.Error("Post order visit wrong order")
	}
	target := 7
	node := Find(root, target)
	if node == nil || node.val != target {
		t.Error("Find() should found target")
	}
	target = 3
	node = NoRecursiveFind(root, target)
	if node == nil || node.val != target {
		t.Error("NoRecursiveFind() should found target")
	}
}

func arrayToString(arr []int) string {
	var vals []string
	for _, v := range arr {
		vals = append(vals, strconv.Itoa(v))
	}
	return strings.Join(vals, ",")
}
