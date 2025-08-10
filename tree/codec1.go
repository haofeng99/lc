package tree

import (
	"strconv"
	"strings"
)

// 中序遍历可以还原二叉树, 但是不唯一
// 前序遍历可以还原任意二叉树
type Codec1 struct {
}

func Constructor1() Codec1 {
	return Codec1{}
}

// Serializes a tree to a single string.
func (this *Codec1) serialize1(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	var res []string
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, "null")
			return
		}
		res = append(res, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return "[" + strings.Join(res, ",") + "]"
}

// Deserializes your encoded data to tree.
// 前序遍历反序列化
func (this *Codec1) deserialize1(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = data[1 : len(data)-1]

	valArr := strings.Split(data, ",")

	cursor := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		valStr := valArr[cursor]
		cursor++
		if valStr == "null" {
			return nil
		}
		val, _ := strconv.Atoi(valStr)
		root := &TreeNode{Val: val}

		root.Left = dfs()
		root.Right = dfs()
		return root
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
